package activity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) (*Repository, error) {
	return &Repository{db: db}, nil
}

func (r Repository) GetAll(ctx context.Context) ([]Activity, error) {
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

func (r Repository) GetExerciseByName(ctx context.Context, name string) ([]Activity, error) {
	activities := make([]Activity, 0)
	query := `SELECT id, user_id, name, duration_minutes,
					total_calories, calories_per_hour, created_at, updated_at
			  	FROM activity
				WHERE name LIKE '%' || $1 || '%'`
	err := r.db.SelectContext(ctx, &activities, query, name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return activities, fmt.Errorf("activity name %s not found: %w", name, err)
		}
		return activities, fmt.Errorf("failed to scan activities: %w", err)
	}

	return activities, nil
}

func (r Repository) GetExerciseById(ctx context.Context, id int) (Activity, error) {
	var activity Activity
	query := `SELECT 	id, user_id, name, duration_minutes,
       					total_calories, calories_per_hour, created_at,
       					updated_at
			   FROM activity
			   WHERE id = $1`
	err := r.db.GetContext(ctx, &activity, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return activity, fmt.Errorf("activity id %d not found: %w", id, err)
		}
		return activity, fmt.Errorf("failed to scan activity: %w", err)
	}

	return activity, nil
}

func (r Repository) Save(ctx context.Context, exerciseSession *ExerciseSession) error {
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

func (r Repository) GetExerciseSessions(ctx context.Context, id int) ([]ExerciseSession, error) {
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

func (r Repository) CalculateAndSaveTotalExerciseSession(ctx context.Context, userID int) (*TotalExerciseSession, error) {
	query := `SELECT
    			duration_hours, duration_minutes, duration_seconds, calories_burned
			FROM exercise_session WHERE user_id = $1`

	var exerciseSessions []ExerciseSession
	err := r.db.SelectContext(ctx, &exerciseSessions, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error making calculations of data: %w", err)
		}
		return nil, fmt.Errorf("failed to fetch exercise sessions: %w", err)
	}

	if len(exerciseSessions) == 0 {
		return nil, nil
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

// GetTotalExerciseOccurrence need to save on db as a history <<<
// using CTEs, provides more general statistics for the most frequent activity across all users and
// the number of times this activity occurs for each user, along with additional information from the total_exercise_session table.
func (r Repository) GetTotalExerciseOccurrence(ctx context.Context, userID int) ([]ExerciseCountStats, error) {
	sessionStats := make([]ExerciseCountStats, 0)

	query := `WITH total_exercise_stats AS (
			  SELECT activity_id
			  FROM exercise_session
			  WHERE user_id = $1
			  GROUP BY activity_id
			  ORDER BY COUNT(*) DESC
			  LIMIT 1
			),
			activity_counts AS (SELECT
			                        user_id,
									activity_id,
									COUNT(*) AS number_of_times
							  FROM exercise_session
							  GROUP BY user_id, activity_id
								LIMIT 1)
			SELECT DISTINCT tes.id, tes.user_id,
							ac.number_of_times, tes.activity_id,
			  				es.session_name, tes.total_duration_hours,
			  				tes.total_duration_minutes,
			  				tes.total_duration_seconds,
			  				tes.total_calories_burned,
			  				tes.created_at, tes.updated_at
			FROM total_exercise_session tes
			JOIN total_exercise_stats mfa ON tes.activity_id = mfa.activity_id
			JOIN activity_counts ac ON tes.user_id = ac.user_id AND tes.activity_id = ac.activity_id
			JOIN exercise_session es ON tes.activity_id = es.activity_id AND tes.user_id = es.user_id`

	err := r.db.SelectContext(ctx, &sessionStats, query, userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sessionStats, fmt.Errorf("exercises not found %w", err)
		}
		return sessionStats, fmt.Errorf("failed to scan exercises: %w", err)
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return sessionStats, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rb := tx.Rollback(); rb != nil {
				log.Fatalf("query failed: %v, unable to abort: %v", err, rb)
			}
			log.Fatal(err)
		}

		if err := tx.Commit(); err != nil {
			log.Fatal(err)
		}
	}()

	stmt := `
		INSERT INTO total_exercise_stats (user_id, activity_id, session_name, number_of_times,
		                                  total_duration_hours, total_duration_minutes, total_duration_seconds,
		                                  total_calories_burned, created_at, updated_at)
		VALUES (:user_id, :activity_id, :session_name, :number_of_times,
		        :total_duration_hours, :total_duration_minutes, :total_duration_seconds,
		        :total_calories_burned, :created_at, :updated_at)
		ON CONFLICT (user_id) DO UPDATE SET
		    session_name = EXCLUDED.session_name,
    		number_of_times = EXCLUDED.number_of_times,
			total_duration_hours = EXCLUDED.total_duration_hours,
			total_duration_minutes = EXCLUDED.total_duration_minutes,
			total_duration_seconds = EXCLUDED.total_duration_seconds,
			total_calories_burned = EXCLUDED.total_calories_burned,
			updated_at = EXCLUDED.updated_at
	`

	for _, stats := range sessionStats {
		stats.CreatedAt = time.Now()
		stats.UpdatedAt = time.Now()

		_, err := tx.NamedExec(stmt, stats)
		if err != nil {
			return sessionStats, fmt.Errorf("failed to upsert exercise stats %w", err)
		}
	}

	return sessionStats, nil
}

// GetExerciseOccurrenceByUser detailed statistics about the most frequently occurring combination of
// session_name and activity_id for a specific user_id from the exercise_session table
func (r Repository) GetExerciseOccurrenceByUser(ctx context.Context, id int) ([]ExerciseCountStats, error) {
	sessionStats := make([]ExerciseCountStats, 0)
	query := `SELECT es.session_name, es.activity_id,
       				COUNT(*) as number_of_times,
       				SUM(es.duration_seconds) as total_duration_seconds,
					SUM(es.duration_minutes) as total_duration_minutes,
					SUM(es.duration_hours) as total_duration_hours,
					SUM(es.calories_burned) as total_calories_burned
              FROM exercise_session es
              WHERE user_id = $1
              GROUP BY session_name, activity_id
              ORDER BY number_of_times DESC
              LIMIT 1`

	err := r.db.SelectContext(ctx, &sessionStats, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sessionStats, fmt.Errorf("exercises not found %w", err)
		}
		return sessionStats, fmt.Errorf("failed to scan exercises: %w", err)
	}

	return sessionStats, nil
}
