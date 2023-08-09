package measurement

import (
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/helpers/db"
	"github.com/google/uuid"
)

type ServiceMeasurements struct {
	repo *RepositoryMeasurement
}

func NewMeasurementService(repo *RepositoryMeasurement) *ServiceMeasurements {
	return &ServiceMeasurements{
		repo: repo,
	}
}

func (s ServiceMeasurements) InsertWeight(w Weight) (Weight, error) {
	weight, err := s.repo.InsertWeight(w)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return Weight{}, db.ErrObjectNotFound{}
	default:
		return Weight{}, err
	}

	return weight, nil
}

func (s ServiceMeasurements) UpdateWeight(id uuid.UUID, userID int, updates map[string]interface{}) error {
	err := s.repo.UpdateWeight(id, userID, updates)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}

	return nil
}

func (s ServiceMeasurements) DeleteWeight(id uuid.UUID, userID int) error {
	err := s.repo.DeleteWeight(id, userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}

	return nil
}

func (s ServiceMeasurements) GetWeight(id uuid.UUID, userID int) (Weight, error) {
	weight, err := s.repo.GetWeight(id, userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return Weight{}, db.ErrObjectNotFound{}
	default:
		return Weight{}, err
	}

	return weight, nil
}

func (s ServiceMeasurements) GetWeights(userID int) ([]Weight, error) {
	weights, err := s.repo.GetWeights(userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []Weight{}, db.ErrObjectNotFound{}
	default:
		return []Weight{}, err
	}

	return weights, nil
}
