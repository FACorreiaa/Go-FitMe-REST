package activity

import (
	"database/sql"
	"time"
)

type Activity struct {
	ID              int            `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID          sql.NullString `json:"user_id,string" db:"user_id"`
	Name            string         `json:"name" db:"name"`
	CaloriesPerHour float32        `json:"calories_per_hour" db:"calories_per_hour"`
	DurationMinutes float32        `json:"duration_minutes" db:"duration_minutes"`
	TotalCalories   float32        `json:"total_calories" db:"total_calories"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at" db:"updated_at"`
}

type ExerciseSession struct {
	ID             int        `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID         int        `json:"user_id" db:"user_id"`
	ActivityID     int        `json:"activity_id" db:"activity_id"`
	SessionName    string     `json:"session_name" db:"session_name"`
	StartTime      time.Time  `json:"start_time" db:"start_time"`
	EndTime        time.Time  `json:"end_time" db:"end_time"`
	Duration       int        `json:"duration" db:"duration"`
	CaloriesBurned int        `json:"calories_burned" db:"calories_burned"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
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
