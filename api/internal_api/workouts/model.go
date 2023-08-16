package workouts

import (
	"github.com/google/uuid"
	"time"
)

type ExerciseList struct {
	ID           uuid.UUID  `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	Name         string     `json:"name" db:"name"`
	ExerciseType string     `json:"type" db:"type"`
	MuscleGroup  string     `json:"muscle" db:"type"`
	Equipment    string     `json:"equipment" db:"equipment"`
	Difficulty   string     `json:"difficulty" db:"difficulty"`
	Instructions string     `json:"instruction" db:"instructions"`
	Video        string     `json:"video" db:"video"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}
