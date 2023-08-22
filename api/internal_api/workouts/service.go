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

//func (s ServiceWorkout) CreateWorkoutPlan(newPlan WorkoutPlan, exerciseIDs []uuid.UUID) (WorkoutPlan, error) {
//	workoutDays := make([]WorkoutDay, len(exerciseIDs))
//	for i, exerciseID := range exerciseIDs {
//		workoutDays[i] = WorkoutDay{
//			WorkoutPlanID: newPlan.ID,
//			ExerciseID:    exerciseID,
//		}
//	}
//
//	workoutPlan, err := s.repo.CreateWorkoutPlan(newPlan, workoutDays)
//	if err != nil {
//		return WorkoutPlan{}, err
//	}
//
//	return workoutPlan, nil
//}

// Service layer

func (s ServiceWorkout) CreateWorkoutPlan(newPlan WorkoutPlan, plan []PlanDay) (WorkoutPlan, error) {
	workoutPlan, err := s.repo.CreateWorkoutPlan(newPlan, plan)
	if err != nil {
		return WorkoutPlan{}, err
	}

	return workoutPlan, nil
}
