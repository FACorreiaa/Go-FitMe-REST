package measurement

import (
	"errors"

	"github.com/FACorreiaa/Stay-Healthy-Backend/db"
)

type ServiceMeasurement struct {
	repo *Repository
}

func NewMeasurementService(repo *Repository) *ServiceMeasurement {
	return &ServiceMeasurement{
		repo: repo,
	}
}

type IMeasurement interface {
	InsertWeight(w Weight) (Weight, error)
	UpdateWeight(id string, userID int, updates map[string]interface{}) error
	DeleteWeight(id string, userID int) error
	GetWeight(id string, userID int) (Weight, error)
	GetWeights(userID int) ([]Weight, error)

	InsertWaistLine(w WaistLine) (WaistLine, error)
	UpdateWaistLine(id string, userID int, updates map[string]interface{}) error
	DeleteWaistLine(id string, userID int) error
	GetWaistLine(id string, userID int) (WaistLine, error)
	GetWaistLines(userID int) ([]WaistLine, error)

	InsertWaterIntake(w WaterIntake) (WaterIntake, error)
	UpdateWaterIntake(id string, userID int, updates map[string]interface{}) error
	DeleteWaterIntake(id string, userID int) error
	GetWaterIntake(id string, userID int) (WaterIntake, error)
	GetWaterIntakes(userID int) ([]WaterIntake, error)
}
type StructMeasurement struct {
	Measurement IMeasurement
}

func (s ServiceMeasurement) InsertWeight(w Weight) (Weight, error) {
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

func (s ServiceMeasurement) UpdateWeight(id string, userID int, updates map[string]interface{}) error {
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

func (s ServiceMeasurement) DeleteWeight(id string, userID int) error {
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

func (s ServiceMeasurement) GetWeight(id string, userID int) (Weight, error) {
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

func (s ServiceMeasurement) GetWeights(userID int) ([]Weight, error) {
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

func (s ServiceMeasurement) InsertWaistLine(w WaistLine) (WaistLine, error) {
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

func (s ServiceMeasurement) UpdateWaistLine(id string, userID int, updates map[string]interface{}) error {
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

func (s ServiceMeasurement) DeleteWaistLine(id string, userID int) error {
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

func (s ServiceMeasurement) GetWaistLine(id string, userID int) (WaistLine, error) {
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

func (s ServiceMeasurement) GetWaistLines(userID int) ([]WaistLine, error) {
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

func (s ServiceMeasurement) InsertWaterIntake(w WaterIntake) (WaterIntake, error) {
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

func (s ServiceMeasurement) UpdateWaterIntake(id string, userID int, updates map[string]interface{}) error {
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

func (s ServiceMeasurement) DeleteWaterIntake(id string, userID int) error {
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

func (s ServiceMeasurement) GetWaterIntake(id string, userID int) (WaterIntake, error) {
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

func (s ServiceMeasurement) GetWaterIntakes(userID int) ([]WaterIntake, error) {
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
