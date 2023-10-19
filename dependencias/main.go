package dependencias

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/calculator"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/measurement"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/user"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/workouts"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthDependencies interface {
	GetDB() *sqlx.DB
	GetRedisClient() *redis.Client
}
type Dependencies interface {
	GetActivityService() *activity.ServiceActivity
	GetUserService() *user.ServiceUser
	GetCalculatorService() *calculator.ServiceCalculator
	GetMeasurementsService() *measurement.ServiceMeasurement
	GetWorkoutsService() *workouts.ServiceWorkout
}
