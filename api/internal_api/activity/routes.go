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
	router.Get("/id={id}", h.GetActivitiesById)
	router.Get("/name={name}", h.GetActivitiesByName)
	router.Post("/start/session/id={id}", h.StartActivityTracker)
	router.Post("/pause/session/id={id}", h.PauseActivityTracker)
	router.Post("/resume/session/id={id}", h.ResumeActivityTracker)
	router.Post("/stop/session/id={id}", h.StopActivityTracker)

	return router
}
