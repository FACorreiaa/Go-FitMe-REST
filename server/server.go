package server

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server/logs"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-chi/httprate"
	"github.com/go-chi/stampede"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"time"
)

func Register(r chi.Router, lg *logrus.Logger, db *sqlx.DB) {
	//swagger
	//logs.InitDefaultLogger()
	//logs.DefaultLogger.Info("Config was successfully imported")
	//configuration, err := configuration.InitConfig()
	//_, err = internals.NewConfig(
	//	configuration.Repositories.Postgres.Host,
	//	configuration.Repositories.Postgres.Port,
	//	configuration.Repositories.Postgres.Username,
	//	os.Getenv("POSTGRES_PASSWORD"),
	//	configuration.Repositories.Postgres.DB,
	//	configuration.Repositories.Postgres.SSLMode,
	//	10*time.Second,
	//	internals.CacheStatement,
	//)
	//logs.DefaultLogger.Info("Server was initialized")
	//
	//if err != nil {
	//	logs.DefaultLogger.WithError(err).Error("Config was not configure")
	//}
	//logs.DefaultLogger.Info("Config was successfully imported")
	//logs.DefaultLogger.ConfigureLogger(
	//	getLogFormatter(configuration.Mode),
	//)
	//logs.DefaultLogger.Info("Main logger was initialized successfully")
	swaggerRoute := SwaggerRoutes()
	//activity routes
	activityRoutes := activity.ActivityRoutes(lg, db)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	cached := stampede.Handler(512, 1*time.Second)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// Logger
	logger := httplog.NewLogger("StayHealthy API", httplog.Options{
		JSON:            true,
		Concise:         true,
		TimeFieldFormat: "Mon, 02 Jan 2006 15:04:05 MST",
	})
	r.Use(httplog.RequestLogger(logger))

	r.Use(middleware.Heartbeat("/ping"))

	InitPprof()
	InitPrometheus(r)
	r.Mount("/api/docs", swaggerRoute)
	r.With(cached).Mount("/api/v1", activityRoutes)

	//h := handler.NewHandler(lg, db)
	//app.Use(handler.MiddlewareLogger())

	//r.Get("/activities", h.GetActivities)
}

func getLogFormatter(mode string) logs.Formatter {
	switch mode {
	case "prod":
		return logs.JSONFormatter
	case "test":
		return logs.DefaultFormatter
	case "dev":
		return logs.DefaultFormatter
	default:
		logs.DefaultLogger.Fatal("Mode has no match")
		os.Exit(1)
		return 0
	}
}
