package activity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (*Repository, error) {
	return &Repository{db: db}, nil
}

func (r Repository) GetAll(ctx context.Context) ([]Activity, error) {
	var activities []Activity
	err := r.db.SelectContext(ctx, &activities, `SELECT id, user_id, name,
										duration_minutes, total_calories, calories_per_hour,
										created_at, updated_at
									FROM activity`)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return activities, fmt.Errorf("activities not found %w", err)
		}
		return activities, fmt.Errorf("failed to scan activities: %w", err)
	}

	return activities, nil
}

func (r Repository) GetExerciseByName(ctx context.Context, name string) ([]Activity, error) {
	var activities []Activity

	err := r.db.SelectContext(ctx, &activities, `SELECT id, user_id, name, duration_minutes,
											total_calories, calories_per_hour, created_at, updated_at
									FROM activity
									WHERE name LIKE '%' || $1 || '%'`, name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return activities, fmt.Errorf("activity with name %s not found: %w", name, err)
		}
		return activities, fmt.Errorf("failed to scan activities: %w", err)
	}

	return activities, nil
}

func (r Repository) GetExerciseByID(ctx context.Context, id int) (Activity, error) {
	var activity Activity

	err := r.db.GetContext(ctx, &activity,
		`SELECT id, user_id, name, duration_minutes, total_calories, calories_per_hour, created_at, updated_at
			   FROM activity
			   WHERE name = $1`, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return activity, fmt.Errorf("activity with id %s not found: %w", id, err)
		}
		return activity, fmt.Errorf("failed to scan activity: %w", err)
	}

	return activity, nil
}
