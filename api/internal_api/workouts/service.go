package workouts

import (
	"context"
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"github.com/google/uuid"
)

type ServiceWorkout struct {
	repo *RepositoryWorkouts
}

func NewWorkoutService(repo *RepositoryWorkouts) *ServiceWorkout {
	return &ServiceWorkout{
		repo: repo,
	}
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

func (s ServiceWorkout) GetExerciseByID(ctx context.Context, id uuid.UUID) (Exercises, error) {
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

func (s ServiceWorkout) DeleteExercise(userID int, exerciseID uuid.UUID) error {
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

func (s ServiceWorkout) UpdateExercise(id uuid.UUID, updates map[string]interface{}) error {
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

func (s ServiceWorkout) GetWorkoutPlan(ctx context.Context, id uuid.UUID) (WorkoutPlanResponse, error) {
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

func (s ServiceWorkout) DeleteWorkoutPlan(userID int, workoutPlanID uuid.UUID) error {
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

func (s ServiceWorkout) UpdateWorkoutPlan(id uuid.UUID, updates map[string]interface{}) error {
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

func (s ServiceWorkout) GetWorkoutPlanIdExercises(ctx context.Context, id uuid.UUID) (WorkoutExerciseDay, error) {
	workoutPlanExercise, err := s.repo.GetWorkoutPlanIdExercises(ctx, id)
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

func (s ServiceWorkout) DeleteWorkoutPlanIdExercises(workoutDay string, workoutPlanID uuid.UUID, exerciseID uuid.UUID) error {
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

func (s ServiceWorkout) CreateWorkoutPlanExercise(workoutDay string, workoutPlanID uuid.UUID, exerciseID uuid.UUID) error {
	err := s.repo.CreateWorkoutPlanExercise(workoutDay, workoutPlanID, exerciseID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}
	return nil
}

func (s ServiceWorkout) UpdateWorkoutPlanExercise(workoutDay string, workoutPlanID uuid.UUID, exerciseID uuid.UUID, prevExerciseID uuid.UUID) error {
	err := s.repo.UpdateWorkoutPlanExercise(workoutDay, workoutPlanID, exerciseID, prevExerciseID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}
	return nil
}
