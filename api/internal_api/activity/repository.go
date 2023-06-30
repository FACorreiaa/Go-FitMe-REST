package activity

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{db: db}
}

func (r Repository) GetAll(ctx context.Context) ([]Activity, error) {
	rows, err := r.db.Query(`SELECT id, user_id, name,
										duration_minutes, total_calories, calories_per_hour,
										created_at, updated_at FROM activity`)

	if err != nil {
		return nil, ctx.Err()
	}
	defer rows.Close()
	result := []Activity{}

	for rows.Next() {
		activity := Activity{}
		if err := rows.Scan(
			&activity.ID, &activity.UserID, &activity.Name,
			&activity.DurationMinutes,
			&activity.TotalCalories, &activity.CaloriesPerHour,
			&activity.CreatedAt, &activity.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error inserting values: %w", err) // Exit if we get an error
		}

		// Append Employee to Employees
		result = append(result, activity)
	}
	// Return Employees in JSON format
	return result, nil
}
