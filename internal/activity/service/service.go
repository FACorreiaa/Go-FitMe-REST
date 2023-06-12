package service

import (
	"context"
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"github.com/FACorreiaa/Stay-Healthy-Backend/internal/activity/model"
	"github.com/FACorreiaa/Stay-Healthy-Backend/internal/activity/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}

func (s Service) GetAll(ctx context.Context) ([]model.Activity, error) {
	activities, err := s.repo.GetAll(ctx)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []model.Activity{}, db.ErrObjectNotFound{}
	default:
		return []model.Activity{}, err
	}
	return activities, nil
}
