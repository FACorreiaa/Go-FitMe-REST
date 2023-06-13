package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database
	ServerPort int `envconfig:"SERVER_PORT" default:"80"`
}

type Database struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     int    `envconfig:"POSTGREST_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Name     string `envconfig:"POSTGRES_DB" required:"true"`
	SslMode  string `envconfig:"POSTGRES_SSLMODE" required:"true"`
}

func LoadEnvVariables() (Config, error) {
	_ = godotenv.Load("app.env")
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	return cnf, err
}
