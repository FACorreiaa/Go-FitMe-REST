package activity

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func RoutesActivity(lg *logrus.Logger, db *sqlx.DB, sm *auth.SessionManager) *chi.Mux {
	h := NewActivityHandler(lg, db, sm)

	router := chi.NewRouter()

	router.Get("/", h.GetActivities)
	router.Post("/start", h.StartTracker)
	router.Post("/stop", h.StopTracker)
	router.Post("/resume", h.ResumeTracker)

	return router
}
