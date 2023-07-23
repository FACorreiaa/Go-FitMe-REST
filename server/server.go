package server

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/calculator"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/go-chi/httprate"
	"github.com/go-chi/stampede"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"time"
)

func Register(r chi.Router, db *sqlx.DB, rdb *redis.Client) {
	swaggerRoute := SwaggerRoutes()
	sessionManager := auth.NewSessionManager(rdb, db)
	userRoutes := user.RoutesUser(db)
	calculatorRoute := calculator.RoutesCalculator()

	activityRoutes := activity.RoutesActivity(db)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	//r.Use(auth.SessionMiddleware(sessionManager))

	// Use middleware to add the session manager to the request context
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
	r.With(cached).Mount("/api/v1/activities", auth.SessionMiddleware(sessionManager)(activityRoutes))
	r.With(cached).Mount("/api/v1/users", userRoutes)
	r.With(cached).Mount("/api/v1/calories", calculatorRoute)
}
