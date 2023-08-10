package measurement

import "github.com/go-chi/chi/v5"

func RoutesMeasurements(deps DependenciesMeasurements) *chi.Mux {
	h := NewMeasurementHandler(deps)
	router := chi.NewRouter()
	//weights
	router.Get("/weights", h.GetWeights)
	router.Get("/weight/{id}", h.GetWeight)
	router.Delete("/weight/{id}", h.DeleteWeight)
	router.Patch("/weight/{id}", h.UpdateWeight)
	router.Post("/weights", h.InsertWeight)

	//water
	router.Get("/water", h.GetWaterIntakes)
	router.Get("/water/{id}", h.GetWaterIntake)
	router.Delete("/water/{id}", h.DeleteWaterIntake)
	router.Patch("/water/{id}", h.UpdateWaterIntake)
	router.Post("/water", h.InsertWaterIntake)

	//waistline
	router.Get("/waistline", h.GetWaistLines)
	router.Get("/waistline/{id}", h.GetWaistLine)
	router.Delete("/waistline/{id}", h.DeleteWaistLine)
	router.Patch("/waistline/{id}", h.UpdateWaistLine)
	router.Post("/waistline", h.InsertWaistLine)
	return router
}
