package workouts

import "github.com/go-chi/chi/v5"

func RoutesWorkouts(deps DependenciesWorkouts) *chi.Mux {
	h := NewExerciseHandler(deps)

	router := chi.NewRouter()

	router.Get("/exercises", h.GetExercises)
	router.Get("/exercises/{id}", h.GetExerciseByID)
	router.Post("/exercises/exercise", h.InsertExercise)
	router.Delete("/exercises/exercise/{id}", h.DeleteExercise)
	router.Patch("/exercises/exercise/{id}", h.UpdateExercise)

	router.Post("/exercises/workout/plan", h.CreateWorkoutPlan)

	return router
}
