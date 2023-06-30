package activity

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func RoutesActivity(lg *logrus.Logger, db *sqlx.DB) *chi.Mux {
	h := NewActivityHandler(lg, db)
	
	router := chi.NewRouter()

	router.Get("/", h.GetActivities)

	return router
}
