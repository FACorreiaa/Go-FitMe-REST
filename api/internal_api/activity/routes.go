package activity

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/dependencies"
	"github.com/go-chi/chi/v5"
)

func RoutesActivity(deps dependencies.Dependencies) *chi.Mux {
	h := NewActivityHandler(deps)

	router := chi.NewRouter()

	router.Get("/", h.GetActivities)
	router.Get("/id={id}", h.GetActivitiesById)
	router.Get("/name={name}", h.GetActivitiesByName)
	router.Get("/user/exercises/user={user_id}", h.GetUserExerciseSession)
	router.Post("/start/session/id={id}", h.StartActivityTracker)
	router.Post("/pause/session/id={id}", h.PauseActivityTracker)
	router.Post("/resume/session/id={id}", h.ResumeActivityTracker)
	router.Post("/stop/session/id={id}", h.StopActivityTracker)

	return router
}
