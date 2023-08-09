package measurement

import (
	"github.com/google/uuid"
	"time"
)

type Weight struct {
	ID          uuid.UUID  `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID      int        `json:"user_id" db:"user_id"`
	WeightValue float32    `json:"weight_value" db:"weight_value"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

type WaterIntake struct {
	ID        uuid.UUID  `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID    int        `json:"user_id" db:"user_id"`
	Quantity  float32    `json:"value" db:"value"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

type WaistLine struct {
	ID             uuid.UUID  `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID         int        `json:"user_id" db:"user_id"`
	WaistLineValue float32    `json:"value" db:"value"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
}
