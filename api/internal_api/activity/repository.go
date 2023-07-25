package activity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RepositoryActivity struct {
	db *sqlx.DB
}

func NewRepositoryActivity(db *sqlx.DB) (*RepositoryActivity, error) {
	return &RepositoryActivity{db: db}, nil
}

func (r RepositoryActivity) GetAll(ctx context.Context) ([]Activity, error) {
	activities := make([]Activity, 0)
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

func (r RepositoryActivity) GetExerciseByName(ctx context.Context, name string) ([]Activity, error) {
	activities := make([]Activity, 0)
	query := `SELECT id, user_id, name, duration_minutes,
					total_calories, calories_per_hour, created_at, updated_at
			  	FROM activity
				WHERE name LIKE '%' || $1 || '%'`
	err := r.db.SelectContext(ctx, &activities, query, name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return activities, fmt.Errorf("activity with name %s not found: %w", name, err)
		}
		return activities, fmt.Errorf("failed to scan activities: %w", err)
	}

	return activities, nil
}

func (r RepositoryActivity) GetExerciseById(ctx context.Context, id int) (Activity, error) {
	var activity Activity
	query := `SELECT 	id, user_id, name, duration_minutes,
       					total_calories, calories_per_hour, created_at,
       					updated_at
			   FROM activity
			   WHERE id = $1`
	err := r.db.GetContext(ctx, &activity, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return activity, fmt.Errorf("activity with id %s not found: %w", id, err)
		}
		return activity, fmt.Errorf("failed to scan activity: %w", err)
	}

	return activity, nil
}

func (r RepositoryActivity) Save(ctx context.Context, exerciseSession *ExerciseSession) error {
	query := `
		INSERT INTO exercise_session
		    (user_id, activity_id, session_name, start_time,
		     end_time, duration_hours, duration_minutes, duration_seconds,
		     calories_burned, created_at)
		VALUES (:user_id, :activity_id, :session_name, :start_time,
		        :end_time, :duration_hours, :duration_minutes, :duration_seconds,
		        :calories_burned, :created_at)
		RETURNING id;
	`

	_, err := r.db.NamedExecContext(ctx, query, exerciseSession)
	if err != nil {
		return fmt.Errorf("failed to insert exercise session: %w", err)
	}

	return nil
}

func (r RepositoryActivity) GetExerciseSessions(ctx context.Context, id int) ([]ExerciseSession, error) {
	exerciseSessions := make([]ExerciseSession, 0)
	query := `SELECT user_id, activity_id, session_name, start_time,
		     		end_time, duration_hours, duration_minutes,
		     		duration_seconds,calories_burned, created_at
				FROM exercise_session
				WHERE user_id = $1`

	err := r.db.SelectContext(ctx, &exerciseSessions, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exerciseSessions, fmt.Errorf("exercises not found %w", err)
		}
		return exerciseSessions, fmt.Errorf("failed to scan exercises: %w", err)
	}

	return exerciseSessions, nil
}
