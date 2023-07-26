package activity

import (
	"github.com/go-chi/chi/v5"
)

func RoutesActivity(deps DependenciesActivity) *chi.Mux {
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
	router.Get("/user/session/total/user={user_id}", h.GetUserExerciseTotalData)

	return router
}
