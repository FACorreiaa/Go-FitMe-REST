package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/calculator"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/measurement"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/user"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/workouts"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server/logs"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

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

func NewWorkoutService(repo *workouts.Repository) *workouts.StructWorkout {
	return &workouts.StructWorkout{
		Workout: workouts.NewWorkoutService(repo),
	}
}

func NewUserService(repo *user.Repository) *user.StructUser {
	return &user.StructUser{
		User: user.NewUserService(repo),
	}
}

func NewMeasurementService(repo *measurement.Repository) *measurement.StructMeasurement {
	return &measurement.StructMeasurement{
		Measurement: measurement.NewMeasurementService(repo),
	}
}

func NewCalculatorService(repo *calculator.Repository) *calculator.StructCalculator {
	return &calculator.StructCalculator{
		Calculator: calculator.NewCalculatorService(repo),
	}
}

func NewActivityService(repo *activity.Repository) *activity.StructActivity {
	return &activity.StructActivity{
		Activity: activity.NewActivityService(repo),
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
		SSLMODE:  cnf.Database.SSLMODE,
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

	activityRepo, err := activity.NewRepository(s.db)
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

	measurementRepo, err := measurement.NewMeasurementRepository(s.db)
	if err != nil {
		_ = errors.New("error injecting calculator service")
	}

	workoutRepo, err := workouts.NewWorkoutsRepository(s.db)
	if err != nil {
		_ = errors.New("error injecting calculator service")
	}

	deps := &AppServices{
		ActivityService:    NewActivityService(activityRepo),
		UserService:        NewUserService(userRepo),
		CalculatorService:  NewCalculatorService(calculatorRepo),
		MeasurementService: NewMeasurementService(measurementRepo),
		WorkoutService:     NewWorkoutService(workoutRepo),
	}

	session := &auth.SessionDependencies{
		DB:    s.db,
		Redis: s.rdb,
	}

	Register(router, deps, session)

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
