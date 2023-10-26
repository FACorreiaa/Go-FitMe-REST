package activity

import (
	"database/sql"
	"time"
)

type Activity struct {
	ID              int            `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID          sql.NullString `json:"user_id,string" db:"user_id" swaggertype:"string"`
	Name            string         `json:"name" db:"name"`
	CaloriesPerHour float32        `json:"calories_per_hour" db:"calories_per_hour"`
	DurationMinutes float32        `json:"duration_minutes" db:"duration_minutes"`
	TotalCalories   float32        `json:"total_calories" db:"total_calories"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at" db:"updated_at"`
}

type ExerciseSession struct {
	ID              string     `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID          int        `json:"user_id" db:"user_id"`
	ActivityID      int        `json:"activity_id" db:"activity_id"`
	SessionName     string     `json:"session_name" db:"session_name"`
	StartTime       time.Time  `json:"start_time" db:"start_time"`
	EndTime         time.Time  `json:"end_time" db:"end_time"`
	DurationHours   int        `json:"duration_hours" db:"duration_hours"`
	DurationMinutes int        `json:"duration_minutes" db:"duration_minutes"`
	DurationSeconds int        `json:"duration_seconds" db:"duration_seconds"`
	CaloriesBurned  int        `json:"calories_burned" db:"calories_burned"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
}

type Duration struct {
	Hours   int
	Minutes int
	Seconds int
}

type TotalExerciseSession struct {
	ID                   string    `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID               int       `json:"user_id" db:"user_id"`
	ActivityID           int       `json:"activity_id" db:"activity_id"`
	TotalDurationHours   int       `json:"duration_hours" db:"total_duration_hours"`
	TotalDurationMinutes int       `json:"duration_minutes" db:"total_duration_minutes"`
	TotalDurationSeconds int       `json:"duration_seconds" db:"total_duration_seconds"`
	TotalCaloriesBurned  int       `json:"calories_burned" db:"total_calories_burned"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}

//type SessionStats struct {
//	ID                     string `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
//	TotalExerciseSessionID string `json:"total_exercise_session_id,string" db:"total_exercise_session_id"`
//	ActivityID             int       `json:"activity_id" db:"activity_id"`
//	UserID                 int       `json:"user_id" db:"user_id"`
//	SessionName            string    `json:"session_name" db:"session_name"`
//	NumberOfTimes          int       `json:"number_of_times" db:"number_of_times"`
//	TotalDurationHours     int       `json:"duration_hours" db:"total_duration_hours"`
//	TotalDurationMinutes   int       `json:"duration_minutes" db:"total_duration_minutes"`
//	TotalDurationSeconds   int       `json:"duration_seconds" db:"total_duration_seconds"`
//	TotalCaloriesBurned    int       `json:"calories_burned" db:"total_calories_burned"`
//	CreatedAt              time.Time `json:"created_at" db:"created_at"`
//	UpdatedAt              time.Time `json:"updated_at" db:"updated_at"`
//}

type ExerciseCountStats struct {
	ID                           string    `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	ActivityID                   int       `json:"activity_id,string" db:"activity_id"`
	UserID                       int       `json:"user_id,string" db:"user_id"`
	SessionName                  string    `json:"session_name" db:"session_name"`
	NumberOfTimes                int       `json:"number_of_times" db:"number_of_times"`
	TotalExerciseDurationHours   int       `json:"total_duration_hours" db:"total_duration_hours"`
	TotalExerciseDurationMinutes int       `json:"total_duration_minutes" db:"total_duration_minutes"`
	TotalExerciseDurationSeconds int       `json:"total_duration_seconds" db:"total_duration_seconds"`
	TotalExerciseCaloriesBurned  int       `json:"total_calories_burned" db:"total_calories_burned"`
	CreatedAt                    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at" db:"updated_at"`
}

type Status int

const (
	StatusPending Status = iota + 1
	StatusInProgress
	StatusDone
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPending:
		return true
	case StatusInProgress:
		return true
	case StatusDone:
		return true
	}
	return false
}
