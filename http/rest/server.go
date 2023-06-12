package rest

import (
	"context"
	"fmt"
	configs "github.com/FACorreiaa/Stay-Healthy-Backend/config"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	logger *logrus.Logger
	router *chi.Mux
	config configs.Config
}

func NewServer() (*Server, error) {
	cnf, err := configs.NewParsedConfig()
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
	server := http.Server{
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
		serverErrors <- server.ListenAndServe()
	}(&wg)

	// blocking run and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("error: starting REST API server: %w", err)
	case <-stopServer:
		s.logger.Warn("server received STOP signal")
		// asking listener to shutdown
		err := server.Shutdown(ctx)
		if err != nil {
			return fmt.Errorf("graceful shutdown did not complete: %w", err)
		}
		wg.Wait()
		s.logger.Info("server was shut down gracefully")
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	fmt.Sprintf("%s:%s", serverHost, serverPort)
	s.router.ServeHTTP(w, r)
}
