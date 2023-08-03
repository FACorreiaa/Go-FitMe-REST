package calculator

import (
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
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
