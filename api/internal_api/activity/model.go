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

type Exercise struct {
	Name             string
	CaloriesBurnedPM float64
}

type ExerciseUserHistory struct {
	ExerciceName   string
	Duration       time.Duration
	CaloriesBurned float64
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
