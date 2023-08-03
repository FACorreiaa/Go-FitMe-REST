package internals

import (
	"context"
	"errors"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/calculator"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/user"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server/logs"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type QueryExecMode uint

const (
	CacheStatement = iota
)

func (m QueryExecMode) value() string {
	switch m {
	case CacheStatement:
		return "cache_statement"
	default:
		return ""
	}
}

type Server struct {
	logger *zap.Logger
	router *chi.Mux
	config ServerConfig
	rdb    *redis.Client
	db     *sqlx.DB
}

func (s *Server) Close() {
	// Close the Redis client
	if s.rdb != nil {
		err := s.rdb.Close()
		if err != nil {
			s.logger.Error("Failed to close Redis client: %s", zap.Error(err))
		}
	}

	// Close the PostgreSQL connection
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			s.logger.Error("Failed to close PostgreSQL connection: %s", zap.Error(err))
		}
	}
}

func NewServer() (*Server, error) {
	cnf, err := LoadEnvVariables()
	if err != nil {
		return nil, err
	}

	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379" // Default value if environment variable is not set
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.Redis.Addr,
		Password: cnf.Redis.Password, // no password set
		DB:       cnf.Redis.DB,       // use default DB
	})

	// Create a context with timeout for the Ping operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the Redis server
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Failed to ping Redis:", err)
	}

	fmt.Println("Redis connection is open. PONG response:", pong)

	database, err := db.Connect(db.ConfigDB{
		Host:     cnf.Database.Host,
		Port:     cnf.Database.Port,
		User:     cnf.Database.User,
		Password: cnf.Database.Password,
		Name:     cnf.Database.Name,
		SslMode:  cnf.Database.SslMode,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	log := NewLogger()
	router := chi.NewRouter()

	s := Server{
		logger: log,
		config: cnf,
		router: router,
		rdb:    rdb,
		db:     database,
	}

	activityRepo, err := activity.NewRepositoryActivity(s.db)
	if err != nil {
		_ = errors.New("error injecting activity service")
	}

	userRepo, err := user.NewUserRepository(s.db)
	if err != nil {
		_ = errors.New("error injecting user service")
	}

	calculatorRepo, err := calculator.NewCalculatorRepository(s.db)
	if err != nil {
		_ = errors.New("error injecting calculator service")
	}

	deps := &server.AppDependencies{
		ActivityService:   activity.NewActivityService(activityRepo),
		UserService:       user.NewUserService(userRepo),
		CalculatorService: calculator.NewCalculatorService(calculatorRepo),
	}

	sessionDeps := &server.SessionDependencies{
		DB:  s.db,
		RDB: s.rdb,
	}

	server.Register(router, deps, sessionDeps)

	return &s, nil
}

func (s *Server) Run(ctx context.Context) error {
	logs.InitDefaultLogger()
	logs.DefaultLogger.Info("Config was successfully imported")
	logs.DefaultLogger.ConfigureLogger(
		logs.JSONFormatter,
	)
	logs.DefaultLogger.Info("Server was initialized")
	serverConfig := http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.ServerPort),
		Handler: cors.Default().Handler(s.router),
	}

	stopServer := make(chan os.Signal, 1)
	signal.Notify(stopServer, syscall.SIGINT, syscall.SIGTERM)

	defer signal.Stop(stopServer)

	// channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		s.logger.Info("REST API listening on port %d", zap.Int("port", s.config.ServerPort))
		serverErrors <- serverConfig.ListenAndServe()
	}(&wg)

	// blocking run and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("error: starting REST API server: %w", err)
	case <-stopServer:
		s.logger.Warn("server received STOP signal")
		// asking listener to Shut down
		err := serverConfig.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("graceful shutdown did not complete: %w", err)
		}
		wg.Wait()
		s.logger.Info("server was shut down gracefully")
	}
	return nil
}
