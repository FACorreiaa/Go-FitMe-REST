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

//weight

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

// Waist line

func (s ServiceMeasurements) InsertWaistLine(w WaistLine) (WaistLine, error) {
	waistLine, err := s.repo.InsertWaistLine(w)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return WaistLine{}, db.ErrObjectNotFound{}
	default:
		return WaistLine{}, err
	}

	return waistLine, nil
}

func (s ServiceMeasurements) UpdateWaistLine(id uuid.UUID, userID int, updates map[string]interface{}) error {
	err := s.repo.UpdateWaistLine(id, userID, updates)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}

	return nil
}

func (s ServiceMeasurements) DeleteWaistLine(id uuid.UUID, userID int) error {
	err := s.repo.DeleteWaistLine(id, userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}

	return nil
}

func (s ServiceMeasurements) GetWaistLine(id uuid.UUID, userID int) (WaistLine, error) {
	waistLine, err := s.repo.GetWaistLine(id, userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return WaistLine{}, db.ErrObjectNotFound{}
	default:
		return WaistLine{}, err
	}

	return waistLine, nil
}

func (s ServiceMeasurements) GetWaistLines(userID int) ([]WaistLine, error) {
	waistLines, err := s.repo.GetWaistLines(userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []WaistLine{}, db.ErrObjectNotFound{}
	default:
		return []WaistLine{}, err
	}

	return waistLines, nil
}

//Water Intake

func (s ServiceMeasurements) InsertWaterIntake(w WaterIntake) (WaterIntake, error) {
	waterIntake, err := s.repo.InsertWater(w)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return WaterIntake{}, db.ErrObjectNotFound{}
	default:
		return WaterIntake{}, err
	}

	return waterIntake, nil
}

func (s ServiceMeasurements) UpdateWaterIntake(id uuid.UUID, userID int, updates map[string]interface{}) error {
	err := s.repo.UpdateWater(id, userID, updates)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}

	return nil
}

func (s ServiceMeasurements) DeleteWaterIntake(id uuid.UUID, userID int) error {
	err := s.repo.DeleteWater(id, userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return db.ErrObjectNotFound{}
	default:
		return err
	}

	return nil
}

func (s ServiceMeasurements) GetWaterIntake(id uuid.UUID, userID int) (WaterIntake, error) {
	waterIntake, err := s.repo.GetWater(id, userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return WaterIntake{}, db.ErrObjectNotFound{}
	default:
		return WaterIntake{}, err
	}

	return waterIntake, nil
}

func (s ServiceMeasurements) GetWaterIntakes(userID int) ([]WaterIntake, error) {
	waterIntakes, err := s.repo.GetAllWater(userID)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return []WaterIntake{}, db.ErrObjectNotFound{}
	default:
		return []WaterIntake{}, err
	}

	return waterIntakes, nil
}
