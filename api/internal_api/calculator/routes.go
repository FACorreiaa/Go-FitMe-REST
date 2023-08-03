package calculator

import "github.com/go-chi/chi/v5"

func RoutesCalculatorOffline() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/calculator/offline", CalculateMacrosOffline)

	return router
}

func RoutesCalculatorSession(deps DependenciesCalculator) *chi.Mux {
	h := NewCalculatorHandler(deps)
	router := chi.NewRouter()
	router.Post("/calculator/{id}", h.CalculateMacros)

	return router
}
