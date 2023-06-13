package internals

import (
	"context"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server"

	"github.com/FACorreiaa/Stay-Healthy-Backend/server/logs"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

//type Config struct {
//	host                 string
//	port                 string
//	username             string
//	password             string
//	dbName               string
//	sslMode              string
//	maxConnWaitingTime   time.Duration
//	defaultQueryExecMode QueryExecMode
//}
//
//func NewConfig(
//	host string,
//	port string,
//	username string,
//	password string,
//	dbName string,
//	sslMode string,
//	maxConnWaitingTime time.Duration,
//	defaultQueryExecMode QueryExecMode,
//) (Config, error) {
//	return Config{
//		host:                 host,
//		port:                 port,
//		username:             username,
//		password:             password,
//		dbName:               dbName,
//		sslMode:              sslMode,
//		maxConnWaitingTime:   maxConnWaitingTime,
//		defaultQueryExecMode: defaultQueryExecMode,
//	}, nil
//}

type Server struct {
	logger *logrus.Logger
	router *chi.Mux
	config ServerConfig
}

func NewServer() (*Server, error) {
	cnf, err := LoadEnvVariables()
	if err != nil {
		return nil, err
	}

	database, err := db.Connect(db.ConfingDB{
		Host:     cnf.Database.Host,
		Port:     cnf.Database.Port,
		User:     cnf.Database.User,
		Password: cnf.Database.Password,
		Name:     cnf.Database.Name,
		SslMode:  cnf.Database.SslMode,
	})
	if err != nil {
		return nil, err
	}

	log := NewLogger()
	router := chi.NewRouter()
	server.Register(router, log, database)

	s := Server{
		logger: log,
		config: cnf,
		router: router,
	}
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
		s.logger.Printf("REST API listening on port %d", s.config.ServerPort)
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
