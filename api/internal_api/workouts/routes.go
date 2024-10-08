package workouts

import "github.com/go-chi/chi/v5"

func RoutesWorkouts(s *StructWorkout) *chi.Mux {
	h := NewExerciseHandler(s)

	router := chi.NewRouter()

	router.Get("/exercises", h.GetExercises)
	router.Get("/exercises/{id}", h.GetExerciseByID)
	router.Post("/exercises/exercise", h.InsertExercise)
	router.Delete("/exercises/exercise/{id}", h.DeleteExercise)
	router.Patch("/exercises/exercise/{id}", h.UpdateExercise)

	router.Get("/exercises/workout/plan/exercise", h.GetWorkoutPlanExercises)
	router.Get("/exercises/workout/plan/exercise/{id}", h.GetExerciseByIdWorkoutPlan)
	router.Delete("/exercises/workout/plan/{workoutPlanID}/day/{workoutDay}/exercise/{exerciseID}", h.DeleteExerciseByIdWorkoutPlan)
	router.Patch("/exercises/workout/plan/{workoutPlanID}/day/{workoutDay}/exercise/{prevExerciseID}/{exerciseID}", h.UpdateExerciseByIdWorkoutPlan)
	router.Post("/exercises/workout/plan/{workoutPlanID}/day/{workoutDay}/exercise/{exerciseID}", h.CreateExerciseWorkoutPlan)

	router.Get("/exercises/workout/plan", h.GetWorkoutPlans)
	router.Get("/exercises/workout/plan/{id}", h.GetWorkoutPlan)
	router.Delete("/exercises/workout/plan/{id}", h.DeleteWorkoutPlan)
	router.Patch("/exercises/workout/plan/{id}", h.UpdateWorkoutPlan)
	router.Post("/exercises/workout/plan", h.CreateWorkoutPlan)

	router.Get("/exercises/workout/plan/{workoutPlanID}/data", h.GetFullWorkoutPlan)
	//router.Get("/export-pdf", h.ExportWorkoutToPDF)
	return router
}
