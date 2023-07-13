package activity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ActivityRepository struct {
	db *sqlx.DB
}

//

func NewActivityRepository(db *sqlx.DB) (*ActivityRepository, error) {
	return &ActivityRepository{db: db}, nil
}

func (r ActivityRepository) GetAll(ctx context.Context) ([]Activity, error) {
	var activities = make([]Activity, 0)
	query := `SELECT id, user_id, name,
					duration_minutes, total_calories, calories_per_hour,
					created_at, updated_at
			FROM activity`

	err := r.db.SelectContext(ctx, &activities, query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return activities, fmt.Errorf("activities not found %w", err)
		}
		return activities, fmt.Errorf("failed to scan activities: %w", err)
	}

	return activities, nil
}

//func (r ActivityRepository) GetAll(ctx context.Context) ([]Activity, error) {
//	var activities []Activity
//	query := `SELECT id, user_id, name,
//                    duration_minutes, total_calories, calories_per_hour,
//                    created_at, updated_at
//            FROM activity`
//
//	rows, err := r.db.Query(query)
//	for rows.Next() {
//		var a Activity
//		err := rows.Scan(
//			&a.ID, &a.UserID, &a.Name, &a.DurationMinutes, &a.TotalCalories,
//			&a.CaloriesPerHour, &a.CreatedAt, &a.UpdatedAt)
//
//		if err != nil {
//			return nil, err
//		}
//		activities = append(activities, a)
//	}
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return activities, fmt.Errorf("activities not found %w", err)
//		}
//		return activities, fmt.Errorf("failed to scan activities: %w", err)
//	}
//
//	return activities, nil
//}

func (r ActivityRepository) GetExerciseByName(ctx context.Context, name string) ([]Activity, error) {
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

func (r ActivityRepository) GetExerciseByID(ctx context.Context, id int) (Activity, error) {
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
