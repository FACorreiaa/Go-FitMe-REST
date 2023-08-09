package measurement

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RepositoryMeasurement struct {
	db *sqlx.DB
}

func NewMeasurementRepository(db *sqlx.DB) (*RepositoryMeasurement, error) {
	return &RepositoryMeasurement{db: db}, nil
}

func (r *RepositoryMeasurement) InsertWeight(w Weight) (Weight, error) {
	query := `
		INSERT INTO weight_measure
		    (id, user_id, weight_value, created_at, updated_at)
		VALUES (:id, :user_id, :weight_value, :created_at, :updated_at)
		RETURNING *;
	`

	result, err := r.db.PrepareNamed(query)
	if err != nil {
		return Weight{}, fmt.Errorf("failed to insert exercise session: %w", err)
	}

	var insertedData Weight
	err = result.Get(&insertedData, w)
	if err != nil {
		return Weight{}, err
	}

	return insertedData, nil
}

func (r *RepositoryMeasurement) UpdateWeight(id uuid.UUID, userID int, updates map[string]interface{}) error {
	query := `
		UPDATE weight_measure
		SET weight_value = :weight_value, updated_at = :updated_at
		WHERE id = :id AND user_id = :user_id
	`

	namedParams := map[string]interface{}{
		"id":           id,
		"user_id":      userID,
		"weight_value": updates["weight_value"],
		"updated_at":   updates["UpdatedAt"],
	}

	_, err := r.db.NamedExec(query, namedParams)
	return err
}

func (r *RepositoryMeasurement) DeleteWeight(id uuid.UUID, userID int) error {
	query := `
		DELETE FROM weight_measure
		WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.Exec(query, id, userID)
	return err
}

func (r *RepositoryMeasurement) GetWeight(id uuid.UUID, userID int) (Weight, error) {
	query := `
		SELECT id, user_id, weight_value, created_at, updated_at FROM weight_measure
		WHERE id = $1 AND user_id = $2
	`
	var weight Weight

	err := r.db.Get(&weight, query, id, userID)
	return weight, err
}

func (r *RepositoryMeasurement) GetWeights(userID int) ([]Weight, error) {
	weights := make([]Weight, 0)

	query := `
		SELECT id, user_id, weight_value, created_at, updated_at FROM weight_measure
		WHERE user_id = $1

	`

	err := r.db.Select(&weights, query, userID)
	return weights, err
}
