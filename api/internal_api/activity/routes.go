package activity

import (
	"github.com/go-chi/chi/v5"
)

func RoutesActivity(s *StructActivity) *chi.Mux {
	h := NewActivityHandler(s)

	router := chi.NewRouter()

	router.Get("/", h.GetActivities)
	router.Get("/{id}", h.GetActivitiesById)
	router.Get("/{name}", h.GetActivitiesByName)
	router.Get("/user/exercises/{user_id}", h.GetUserExerciseSession)
	router.Get("/user/session/total/{user_id}", h.GetUserExerciseTotalData)

	router.Get("/user/session/stats/{user_id}", h.GetUserExerciseSessionStats)
	router.Get("/user/session/stats/total/{user_id}", h.GetExerciseSessionStats)

	router.Post("/start/session/{id}", h.StartActivityTracker)
	router.Post("/pause/session/{id}", h.PauseActivityTracker)
	router.Post("/resume/session/{id}", h.ResumeActivityTracker)
	router.Post("/stop/session/{id}", h.StopActivityTracker)

	return router
}
