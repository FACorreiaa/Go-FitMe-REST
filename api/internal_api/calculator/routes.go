package calculator

import "github.com/go-chi/chi/v5"

func RoutesCalculatorOffline() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/offline", CalculateMacrosOffline)

	return router
}

func RoutesCalculatorSession(deps DependenciesCalculator) *chi.Mux {
	h := NewCalculatorHandler(deps)
	router := chi.NewRouter()
	router.Post("/{user_id}", h.CalculateMacros)
	router.Get("/{user_id}", h.GetAllDietMacros)
	router.Get("/plan/{id}", h.GetDietMacros)

	return router
}
