package workouts

type ServiceWorkout struct {
	repo *RepositoryWorkouts
}

func NewWorkoutService(repo *RepositoryWorkouts) *ServiceWorkout {
	return &ServiceWorkout{
		repo: repo,
	}
}
