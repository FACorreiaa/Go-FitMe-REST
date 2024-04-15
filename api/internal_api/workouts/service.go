package workouts

import (
	"context"
	"errors"

	"github.com/FACorreiaa/Stay-Healthy-Backend/db"
)

type ServiceWorkout struct {
	repo *Repository
}

func NewWorkoutService(repo *Repository) *ServiceWorkout {
	return &ServiceWorkout{
		repo: repo,
	}
}

type IWorkout interface {
	GetAllExercises(ctx context.Context) ([]Exercises, error)
	GetExerciseByID(ctx context.Context, id string) (Exercises, error)
	InsertExercise(id int, exercise Exercises) (Exercises, error)
	DeleteExercise(userID int, exerciseID string) error
	UpdateExercise(id string, updates map[string]interface{}) error
	CreateWorkoutPlan(newPlan WorkoutPlan, plan []PlanDay) (WorkoutPlan, error)
	GetWorkoutPlans(ctx context.Context) ([]WorkoutPlanResponse, error)
	DeleteWorkoutPlan(userID int, workoutPlanID string) error
	GetWorkoutPlan(ctx context.Context, id string) (WorkoutPlanResponse, error)
	UpdateWorkoutPlan(id string, updates map[string]interface{}) error
	GetExerciseByIdWorkoutPlan(ctx context.Context, id string) (WorkoutExerciseDay, error)
	GetWorkoutPlanExercises(ctx context.Context) ([]WorkoutExerciseDay, error)
	DeleteWorkoutPlanIdExercises(workoutDay string, workoutPlanID string, exerciseID string) error
	CreateExerciseWorkoutPlan(workoutDay string, workoutPlanID string, exerciseID string) error
	UpdateExerciseByIdWorkoutPlan(workoutDay string, workoutPlanID string, exerciseID string, prevExerciseID string) error
	GetFullWorkoutPlan(ctx context.Context, workoutPlanID string, userID int) (WorkoutPlan, error)
}
type StructWorkout struct {
	Workout IWorkout
}

func (s ServiceWorkout) GetAllExercises(ctx context.Context) ([]Exercises, error) {
	exercises, err := s.repo.GetAllExercises(ctx)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []Exercises{}, db.ErrObjectNotFound{}
	default:
		return []Exercises{}, err
	}
	return exercises, nil
}

func (s ServiceWorkout) GetExerciseByID(ctx context.Context, id string) (Exercises, error) {
	exercise, err := s.repo.GetExerciseByID(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return Exercises{}, db.ErrObjectNotFound{}
	default:
		return Exercises{}, err
	}
	return exercise, nil
}

func (s ServiceWorkout) InsertExercise(id int, exercise Exercises) (Exercises, error) {
	exercise, err := s.repo.InsertNewExercise(id, exercise)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return Exercises{}, db.ErrObjectNotFound{}
	default:
		return Exercises{}, err
	}
	return exercise, nil
}

func (s ServiceWorkout) DeleteExercise(userID int, exerciseID string) error {
	err := s.repo.DeleteUserExercise(userID, exerciseID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}
	return nil
}

func (s ServiceWorkout) UpdateExercise(id string, updates map[string]interface{}) error {
	err := s.repo.UpdateExercise(id, updates)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}

	return nil
}

func (s ServiceWorkout) CreateWorkoutPlan(newPlan WorkoutPlan, plan []PlanDay) (WorkoutPlan, error) {
	workoutPlan, err := s.repo.CreateWorkoutPlan(newPlan, plan)
	if err != nil {
		return WorkoutPlan{}, err
	}

	return workoutPlan, nil
}

func (s ServiceWorkout) GetWorkoutPlans(ctx context.Context) ([]WorkoutPlanResponse, error) {
	workoutPlan, err := s.repo.GetWorkoutPlans(ctx)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []WorkoutPlanResponse{}, db.ErrObjectNotFound{}
	default:
		return []WorkoutPlanResponse{}, err
	}
	return workoutPlan, nil
}

func (s ServiceWorkout) GetWorkoutPlan(ctx context.Context, id string) (WorkoutPlanResponse, error) {
	workoutPlan, err := s.repo.GetWorkoutPlan(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return WorkoutPlanResponse{}, db.ErrObjectNotFound{}
	default:
		return WorkoutPlanResponse{}, err
	}
	return workoutPlan, nil
}

func (s ServiceWorkout) DeleteWorkoutPlan(userID int, workoutPlanID string) error {
	err := s.repo.DeleteWorkoutPlan(userID, workoutPlanID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}
	return nil
}

func (s ServiceWorkout) UpdateWorkoutPlan(id string, updates map[string]interface{}) error {
	err := s.repo.UpdateWorkoutPlan(id, updates)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}

	return nil
}

func (s ServiceWorkout) GetExerciseByIdWorkoutPlan(ctx context.Context, id string) (WorkoutExerciseDay, error) {
	workoutPlanExercise, err := s.repo.GetExerciseByIdWorkoutPlan(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return WorkoutExerciseDay{}, db.ErrObjectNotFound{}
	default:
		return WorkoutExerciseDay{}, err
	}
	return workoutPlanExercise, nil
}

func (s ServiceWorkout) GetWorkoutPlanExercises(ctx context.Context) ([]WorkoutExerciseDay, error) {
	workoutPlanExercises, err := s.repo.GetWorkoutPlanExercises(ctx)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []WorkoutExerciseDay{}, db.ErrObjectNotFound{}
	default:
		return []WorkoutExerciseDay{}, err
	}
	return workoutPlanExercises, nil
}

func (s ServiceWorkout) DeleteWorkoutPlanIdExercises(workoutDay string, workoutPlanID string, exerciseID string) error {
	err := s.repo.DeleteWorkoutPlanIdExercises(workoutDay, workoutPlanID, exerciseID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}
	return nil
}

func (s ServiceWorkout) CreateExerciseWorkoutPlan(workoutDay string, workoutPlanID string, exerciseID string) error {
	err := s.repo.CreateExerciseWorkoutPlan(workoutDay, workoutPlanID, exerciseID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}
	return nil
}

func (s ServiceWorkout) UpdateExerciseByIdWorkoutPlan(workoutDay string, workoutPlanID string, exerciseID string, prevExerciseID string) error {
	err := s.repo.UpdateExerciseByIdWorkoutPlan(workoutDay, workoutPlanID, exerciseID, prevExerciseID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}
	return nil
}

func (s ServiceWorkout) GetFullWorkoutPlan(ctx context.Context, workoutPlanID string, userID int) (WorkoutPlan, error) {
	workoutPlan, err := s.repo.GetFullWorkoutPlan(ctx, workoutPlanID, userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return WorkoutPlan{}, db.ErrObjectNotFound{}
	default:
		return WorkoutPlan{}, err
	}
	return workoutPlan, nil
}

//func (s ServiceWorkout) ExportWorkoutToPDF() {}
