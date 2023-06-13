package activity

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/internal/activity/handler"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func ActivityRoutes(lg *logrus.Logger, db *sqlx.DB) *chi.Mux {
	h := handler.NewActivityHandler(lg, db)
	router := chi.NewRouter()

	router.Get("/activities", h.GetActivities)
	//router.Get("/api/v1/tax/tax-name={tax_name}", taxHandler.GetTaxName)
	//router.Get("/api/v1/tax/count", taxHandler.GetTaxesCount)
	//
	//router.Route("/api/v1/tax/{id}", func(r chi.Router) {
	//	r.Get("/", taxHandler.GetTax)
	//	r.Delete("/", taxHandler.DeleteTax)
	//	r.Put("/", taxHandler.UpdateTax)
	//})
	return router
}
