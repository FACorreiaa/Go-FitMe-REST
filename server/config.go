package server

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/calculator"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/measurement"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/user"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/workouts"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type QueryExecMode uint

type AppServices struct {
	ActivityService    *activity.StructActivity
	UserService        *user.StructUser
	CalculatorService  *calculator.StructCalculator
	MeasurementService *measurement.StructMeasurement
	WorkoutService     *workouts.StructWorkout
}

type Server struct {
	logger      *zap.Logger
	router      *chi.Mux
	config      ServerConfig
	redisClient *redis.Client
	db          *sqlx.DB
}

type ServerConfig struct {
	Database   Database
	Redis      Redis
	ServerPort int    `envconfig:"SERVER_PORT" default:"80"`
	Env        string `envconfig:"STAY_HEALTHY_ENV"`
}

type Database struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"POSTGREST_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_DB" required:"true"`
	SSLMODE  string `envconfig:"DB_SSLMODE" required:"true"`
}

type Redis struct {
	Addr     string `envconfig:"REDIS_HOST" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" required:"true"`
	DB       int    `envconfig:"REDIS_DB" required:"true"`
}

// refactor later for Viper

func LoadEnvVariables() (ServerConfig, error) {
	cnf := ServerConfig{}
	err := envconfig.Process("", &cnf)
	return cnf, err
}
