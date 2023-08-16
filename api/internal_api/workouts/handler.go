package workouts

type DependenciesWorkouts interface {
	GetWorkoutsService() *ServiceWorkout
}

type Handler struct {
	dependencies DependenciesWorkouts
}

func NewActivityHandler(deps DependenciesWorkouts) *Handler {
	return &Handler{
		dependencies: deps,
	}
}
