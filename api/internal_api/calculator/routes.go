package calculator

import "github.com/go-chi/chi/v5"

func RoutesCalculator() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/calculator/offline", CalculateMacros)

	return router
}
