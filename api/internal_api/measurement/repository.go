package measurement

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewMeasurementRepository(db *sqlx.DB) (*Repository, error) {
	return &Repository{db: db}, nil
}

//weight

func (r *Repository) InsertWeight(w Weight) (Weight, error) {
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

func (r *Repository) UpdateWeight(id string, userID int, updates map[string]interface{}) error {
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

func (r *Repository) DeleteWeight(id string, userID int) error {
	query := `
		DELETE FROM weight_measure
		WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.Exec(query, id, userID)
	return err
}

func (r *Repository) GetWeight(id string, userID int) (Weight, error) {
	query := `
		SELECT id, user_id, weight_value, created_at, updated_at FROM weight_measure
		WHERE id = $1 AND user_id = $2
	`
	var weight Weight

	err := r.db.Get(&weight, query, id, userID)
	return weight, err
}

func (r *Repository) GetWeights(userID int) ([]Weight, error) {
	weights := make([]Weight, 0)

	query := `
		SELECT id, user_id, weight_value, created_at, updated_at FROM weight_measure
		WHERE user_id = $1

	`

	err := r.db.Select(&weights, query, userID)
	return weights, err
}

//water

func (r *Repository) InsertWater(w WaterIntake) (WaterIntake, error) {
	query := `
		INSERT INTO water_intake
		    (id, user_id, quantity, created_at, updated_at)
		VALUES (:id, :user_id, :quantity, :created_at, :updated_at)
		RETURNING *;
	`

	result, err := r.db.PrepareNamed(query)
	if err != nil {
		return WaterIntake{}, fmt.Errorf("failed to insert exercise session: %w", err)
	}

	var insertedData WaterIntake
	err = result.Get(&insertedData, w)
	if err != nil {
		return WaterIntake{}, err
	}

	return insertedData, nil
}

func (r *Repository) UpdateWater(id string, userID int, updates map[string]interface{}) error {
	query := `
		UPDATE water_intake
		SET quantity = :quantity, updated_at = :updated_at
		WHERE id = :id AND user_id = :user_id
	`

	namedParams := map[string]interface{}{
		"id":           id,
		"user_id":      userID,
		"weight_value": updates["quantity"],
		"updated_at":   updates["UpdatedAt"],
	}

	_, err := r.db.NamedExec(query, namedParams)
	return err
}

func (r *Repository) DeleteWater(id string, userID int) error {
	query := `
		DELETE FROM water_intake
		WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.Exec(query, id, userID)
	return err
}

func (r *Repository) GetWater(id string, userID int) (WaterIntake, error) {
	query := `
		SELECT id, user_id, quantity, created_at, updated_at FROM water_intake
		WHERE id = $1 AND user_id = $2
	`
	var water WaterIntake

	err := r.db.Get(&water, query, id, userID)
	return water, err
}

func (r *Repository) GetAllWater(userID int) ([]WaterIntake, error) {
	water := make([]WaterIntake, 0)

	query := `
		SELECT id, user_id, quantity, created_at, updated_at FROM water_intake
		WHERE user_id = $1

	`

	err := r.db.Select(&water, query, userID)
	return water, err
}

//waist line

func (r *Repository) InsertWaistLine(w WaistLine) (WaistLine, error) {
	query := `
		INSERT INTO waist_line
		    (id, user_id, quantity, created_at, updated_at)
		VALUES (:id, :user_id, :quantity, :created_at, :updated_at)
		RETURNING *;
	`

	result, err := r.db.PrepareNamed(query)
	if err != nil {
		return WaistLine{}, fmt.Errorf("failed to insert exercise session: %w", err)
	}

	var insertedData WaistLine
	err = result.Get(&insertedData, w)
	if err != nil {
		return WaistLine{}, err
	}

	return insertedData, nil
}

func (r *Repository) UpdateWaistLine(id string, userID int, updates map[string]interface{}) error {
	query := `
		UPDATE waist_line
		SET quantity = :quantity, updated_at = :updated_at
		WHERE id = :id AND user_id = :user_id
	`

	namedParams := map[string]interface{}{
		"id":           id,
		"user_id":      userID,
		"weight_value": updates["quantity"],
		"updated_at":   updates["UpdatedAt"],
	}

	_, err := r.db.NamedExec(query, namedParams)
	return err
}

func (r *Repository) DeleteWaistLine(id string, userID int) error {
	query := `
		DELETE FROM waist_line
		WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.Exec(query, id, userID)
	return err
}

func (r *Repository) GetWaistLine(id string, userID int) (WaistLine, error) {
	query := `
		SELECT id, user_id, quantity, created_at, updated_at FROM waist_line
		WHERE id = $1 AND user_id = $2
	`
	var w WaistLine

	err := r.db.Get(&w, query, id, userID)
	return w, err
}

func (r *Repository) GetWaistLines(userID int) ([]WaistLine, error) {
	w := make([]WaistLine, 0)

	query := `
		SELECT id, user_id, quantity, created_at, updated_at FROM waist_line
		WHERE user_id = $1

	`

	err := r.db.Select(&w, query, userID)
	return w, err
}
