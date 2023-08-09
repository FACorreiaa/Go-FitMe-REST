package measurement

import "github.com/go-chi/chi/v5"

func RoutesMeasurements(deps DependenciesMeasurements) *chi.Mux {
	h := NewMeasurementHandler(deps)
	router := chi.NewRouter()
	router.Get("/weights", h.GetWeights)
	router.Get("/weight/{id}", h.GetWeight)
	router.Delete("/weight/{id}", h.DeleteWeight)
	router.Patch("/weight/{id}", h.UpdateWeight)

	router.Post("/weights", h.InsertWeight)

	return router
}
