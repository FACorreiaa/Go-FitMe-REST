package calculator

import (
	"context"
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"github.com/google/uuid"
)

type ServiceCalculator struct {
	repo *Repository
}

func NewCalculatorService(repo *Repository) *ServiceCalculator {
	return &ServiceCalculator{
		repo: repo,
	}
}

func (s ServiceCalculator) Create(user UserMacroDistribution) (UserMacroDistribution, error) {
	diet, err := s.repo.InsertDietGoals(user)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return UserMacroDistribution{}, db.ErrObjectNotFound{}
	default:
		return UserMacroDistribution{}, err
	}

	return diet, nil
}

func (s ServiceCalculator) GetAll(ctx context.Context, userID int) ([]UserMacroDistribution, error) {
	userMacros, err := s.repo.GetUserDietGoals(ctx, userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []UserMacroDistribution{}, db.ErrObjectNotFound{}
	default:
		return []UserMacroDistribution{}, err
	}
	return userMacros, err
}

func (s ServiceCalculator) Get(ctx context.Context, planID uuid.UUID) (UserMacroDistribution, error) {
	userMacros, err := s.repo.GetUserDietGoal(ctx, planID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return UserMacroDistribution{}, db.ErrObjectNotFound{}
	default:
		return UserMacroDistribution{}, err
	}
	return userMacros, err
}
