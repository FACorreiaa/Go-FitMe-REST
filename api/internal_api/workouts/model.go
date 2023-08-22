package workouts

import (
	"github.com/google/uuid"
	"time"
)

type Exercises struct {
	ID            uuid.UUID  `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	Name          string     `json:"name" db:"name"`
	ExerciseType  string     `json:"type" db:"type"`
	MuscleGroup   string     `json:"muscle" db:"muscle"`
	Equipment     string     `json:"equipment" db:"equipment"`
	Difficulty    string     `json:"difficulty" db:"difficulty"`
	Instructions  string     `json:"instructions" db:"instructions"`
	Video         string     `json:"video" db:"video"`
	CustomCreated bool       `json:"custom_created" db:"custom_created"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
}

type WorkoutPlan struct {
	ID          uuid.UUID    `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID      int          `json:"user_id" db:"user_id"`
	Description string       `json:"description" db:"description"`
	Notes       string       `json:"notes" db:"notes"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at" db:"updated_at"`
	Rating      int          `json:"rating" db:"rating"`
	WorkoutDays []WorkoutDay `json:"workoutDays" db:"-"`
}

type WorkoutDay struct {
	ID            uuid.UUID   `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	WorkoutPlanID uuid.UUID   `json:"workout_plan_id" db:"workout_plan_id"`
	Day           string      `json:"day" db:"day"`
	CreatedAt     time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at" db:"updated_at"`
	Exercises     []Exercises `json:"exercises" db:"-"`
}

type WorkoutDayExercise struct {
	WorkoutDayID uuid.UUID `db:"workout_day_id"`
	ExerciseID   uuid.UUID `db:"exercise_id"`
}

type PlanDay struct {
	Day         string      `json:"day"`
	ExerciseIDs []uuid.UUID `json:"exerciseIDs"`
}

type CreateWorkoutPlanRequest struct {
	WorkoutPlan WorkoutPlan `json:"workoutPlan"`
	Plan        []PlanDay   `json:"plan"`
}

type WorkoutPlanDetail struct {
	ID            uuid.UUID   `db:"id"`
	WorkoutPlanID uuid.UUID   `db:"workout_plan_id"`
	Day           string      `db:"day"`
	Exercises     []uuid.UUID `db:"exercises"`
	CreatedAt     time.Time   `db:"created_at"`
}
