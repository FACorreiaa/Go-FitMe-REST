package activity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
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

func (r RepositoryActivity) updateTotalExerciseSession(ctx context.Context, totalExerciseSession TotalExerciseSession, id int) error {
	// Check if the total_exercise_session row already exists in the database for the given user_id
	// If it exists, update the row; otherwise, insert a new row
	existingQuery := `SELECT COUNT(*) FROM total_exercise_session WHERE user_id = $1`
	var count int
	err := r.db.GetContext(ctx, &count, existingQuery, id)
	if err != nil {
		return fmt.Errorf("failed to check if total exercise session exists: %w", err)
	}

	if count > 0 {
		// Update the existing row
		updateQuery := `UPDATE total_exercise_session SET
                    	user_id = $1.
                        total_duration_hours = $2,
                        total_duration_minutes = $3,
                        total_duration_seconds = $4,
                        total_calories_burned = $5,
                        created_at = $6,
                        updated_at = $7
                        WHERE user_id = $8`

		_, err = r.db.ExecContext(ctx, updateQuery,
			totalExerciseSession.UserID,
			totalExerciseSession.TotalDurationHours,
			totalExerciseSession.TotalDurationMinutes,
			totalExerciseSession.TotalDurationSeconds,
			totalExerciseSession.TotalCaloriesBurned,
			time.Now(),
			time.Now(),
			id,
		)
		if err != nil {
			return fmt.Errorf("failed to update total exercise session: %w", err)
		}
	} else {
		// Insert a new row
		insertQuery := `INSERT INTO total_exercise_session (user_id, total_duration_hours, total_duration_minutes, total_duration_seconds, total_calories_burned, created_at, updated_at)
                        VALUES ($1, $2, $3, $4, $5)`

		_, err = r.db.ExecContext(ctx, insertQuery,
			totalExerciseSession.UserID,
			totalExerciseSession.TotalDurationHours,
			totalExerciseSession.TotalDurationMinutes,
			totalExerciseSession.TotalDurationSeconds,
			totalExerciseSession.TotalCaloriesBurned,
			totalExerciseSession.CreatedAt,
			totalExerciseSession.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("failed to insert total exercise session: %w", err)
		}
	}

	return nil
}

func (r RepositoryActivity) CalculateAndSaveTotalExerciseSession(ctx context.Context, userID int) (*TotalExerciseSession, error) {
	query := `SELECT duration_hours, duration_minutes, duration_seconds, calories_burned FROM exercise_session WHERE user_id = $1`

	var exerciseSessions []ExerciseSession
	err := r.db.SelectContext(ctx, &exerciseSessions, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error making calculations of data: %w", err)
		}
		return nil, fmt.Errorf("failed to fetch exercise sessions: %w", err)
	}

	totalDuration := Duration{}
	totalCaloriesBurned := 0

	for _, session := range exerciseSessions {
		totalDuration.Hours += session.DurationHours
		totalDuration.Minutes += session.DurationMinutes
		totalDuration.Seconds += session.DurationSeconds
		totalCaloriesBurned += session.CaloriesBurned
	}

	// Create or update the total_exercise_session row in the database
	_, err = r.db.ExecContext(ctx, `
		INSERT INTO total_exercise_session (user_id, total_duration_hours, total_duration_minutes, total_duration_seconds, total_calories_burned)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id)
		DO UPDATE SET
			total_duration_hours = EXCLUDED.total_duration_hours,
			total_duration_minutes = EXCLUDED.total_duration_minutes,
			total_duration_seconds = EXCLUDED.total_duration_seconds,
			total_calories_burned = EXCLUDED.total_calories_burned,
			updated_at = NOW()
	`, userID, totalDuration.Hours, totalDuration.Minutes, totalDuration.Seconds, totalCaloriesBurned)

	if err != nil {
		return nil, fmt.Errorf("failed to save total exercise session: %w", err)
	}

	return &TotalExerciseSession{
		ID:                   uuid.New(),
		UserID:               userID,
		TotalDurationHours:   totalDuration.Hours,
		TotalDurationMinutes: totalDuration.Minutes,
		TotalDurationSeconds: totalDuration.Seconds,
		TotalCaloriesBurned:  totalCaloriesBurned,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}, nil
}
