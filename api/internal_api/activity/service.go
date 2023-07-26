package activity

import (
	"context"
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"time"
)

type ServiceActivity struct {
	repo *RepositoryActivity
}

func NewActivityService(repo *RepositoryActivity) *ServiceActivity {
	return &ServiceActivity{
		repo: repo,
	}
}

func CalculateTotalFromJSONData(exerciseSessions []ExerciseSession, id int) TotalExerciseSession {
	totalDurationHours := 0
	totalDurationMinutes := 0
	totalDurationSeconds := 0
	totalCaloriesBurned := 0

	for _, session := range exerciseSessions {
		totalDurationHours += session.DurationHours
		totalDurationMinutes += session.DurationMinutes
		totalDurationSeconds += session.DurationSeconds
		totalCaloriesBurned += session.CaloriesBurned
	}

	return TotalExerciseSession{
		UserID:               id,
		TotalDurationHours:   totalDurationHours,
		TotalDurationMinutes: totalDurationMinutes,
		TotalDurationSeconds: totalDurationSeconds,
		TotalCaloriesBurned:  totalCaloriesBurned,
		CreatedAt:            time.Now(),
	}
}

func (s ServiceActivity) GetAll(ctx context.Context) ([]Activity, error) {
	activities, err := s.repo.GetAll(ctx)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []Activity{}, db.ErrObjectNotFound{}
	default:
		return []Activity{}, err
	}
	return activities, nil
}

func (s ServiceActivity) Get(ctx context.Context) ([]Activity, error) {
	activities, err := s.repo.GetAll(ctx)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []Activity{}, db.ErrObjectNotFound{}
	default:
		return []Activity{}, err
	}
	return activities, nil
}

func (s ServiceActivity) GetByName(ctx context.Context, name string) ([]Activity, error) {
	activities, err := s.repo.GetExerciseByName(ctx, name)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []Activity{}, db.ErrObjectNotFound{}
	default:
		return []Activity{}, err
	}
	return activities, nil
}

func (s ServiceActivity) GetByID(ctx context.Context, id int) (Activity, error) {
	activity, err := s.repo.GetExerciseById(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return Activity{}, db.ErrObjectNotFound{}
	default:
		return Activity{}, err
	}
	return activity, nil
}

func (s ServiceActivity) SaveExerciseSession(ctx context.Context, session *ExerciseSession) error {
	err := s.repo.Save(ctx, session)

	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}
	return nil
}

func (s ServiceActivity) GetExerciseSession(ctx context.Context, id int) ([]ExerciseSession, error) {
	exerciseSession, err := s.repo.GetExerciseSessions(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []ExerciseSession{}, db.ErrObjectNotFound{}
	default:
		return []ExerciseSession{}, err
	}
	return exerciseSession, nil

}

func (s ServiceActivity) GetExerciseTotalSession(ctx context.Context, userId int) (*TotalExerciseSession, error) {
	exerciseSessionTotalValues, err := s.repo.CalculateAndSaveTotalExerciseSession(ctx, userId)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return &TotalExerciseSession{}, db.ErrObjectNotFound{}
	default:
		return &TotalExerciseSession{}, err
	}
	return exerciseSessionTotalValues, nil
}
