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

type AppDependencies struct {
	ActivityService    *activity.ServiceActivity
	UserService        *user.ServiceUser
	CalculatorService  *calculator.ServiceCalculator
	MeasurementService *measurement.ServiceMeasurements
	WorkoutService     *workouts.ServiceWorkout
	// You may add other dependencies here if needed
}

func (ad *AppDependencies) GetActivityService() *activity.ServiceActivity {
	return ad.ActivityService
}

func (ad *AppDependencies) GetUserService() *user.ServiceUser {
	return ad.UserService
}

func (ad *AppDependencies) GetCalculatorService() *calculator.ServiceCalculator {
	return ad.CalculatorService
}

func (ad *AppDependencies) GetMeasurementService() *measurement.ServiceMeasurements {
	return ad.MeasurementService
}

func (ad *AppDependencies) GetWorkoutsService() *workouts.ServiceWorkout {
	return ad.WorkoutService
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
