package server

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/calculator"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/measurement"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/user"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/workouts"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AppServices struct {
	ActivityService    *activity.StructActivity
	UserService        *user.StructUser
	CalculatorService  *calculator.StructCalculator
	MeasurementService *measurement.StructMeasurement
	WorkoutService     *workouts.StructWorkout
}

type SessionDependencies struct {
	DB  *sqlx.DB
	RDB *redis.Client
	// You may add other dependencies here if needed
}

func (ad *SessionDependencies) GetDB() *sqlx.DB {
	return ad.DB
}

func (ad *SessionDependencies) GetRedisClient() *redis.Client {
	return ad.RDB
}
