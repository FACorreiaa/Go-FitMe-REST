package activity

import (
	"context"
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s Service) GetAll(ctx context.Context) ([]Activity, error) {
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

func (s Service) Get(ctx context.Context) ([]Activity, error) {
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

func (s Service) GetByName(ctx context.Context, name string) ([]Activity, error) {
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

func (s Service) GetByID(ctx context.Context, id int) (Activity, error) {
	activity, err := s.repo.GetExerciseByID(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return Activity{}, db.ErrObjectNotFound{}
	default:
		return Activity{}, err
	}
	return activity, nil
}
