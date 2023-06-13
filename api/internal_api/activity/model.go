package activity

import (
	"database/sql"
	"time"
)

type Activity struct {
	ID              int            `json:"id,string" pg:"default:gen_random_uuid()"`
	UserID          sql.NullString `json:"user_id,string"`
	Name            string         `json:"name"`
	CaloriesPerHour float32        `json:"calories_per_hour"`
	DurationMinutes float32        `json:"duration_minutes"`
	TotalCalories   float32        `json:"total_calories"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       *time.Time     `json:"updated_at"`
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
