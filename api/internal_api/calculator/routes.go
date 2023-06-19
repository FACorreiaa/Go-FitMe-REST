package calculator

import "github.com/go-chi/chi/v5"

func RoutesCalculator() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/offline/measure={metric}/age={age}/gender={gender}/height={height}/weight={weight}/activity={activity}/objective={objective}/distribution={calories-distribution}", CalculateMacros)

	return router
}
