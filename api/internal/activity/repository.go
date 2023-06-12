package activity

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ActivityQueries struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) *ActivityQueries {
	return &ActivityQueries{db: db}
}

func (a *ActivityQueries) GetAll() ([]domain.Activity, error) {
	activities := []domain.Activity{}
	query := `select * from activity`
	err := a.db.Get(&activities, query)
	if err != nil {
		// Return empty object and error.
		return activities, err
	}
	return activities, nil
}
