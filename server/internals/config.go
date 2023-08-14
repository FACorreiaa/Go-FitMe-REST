package internals

import (
	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Database   Database
	Redis      Redis
	ServerPort int    `envconfig:"SERVER_PORT" default:"80"`
	Env        string `envconfig:"STAY_HEALTHY_ENV"`
}

type Database struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     int    `envconfig:"POSTGREST_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Name     string `envconfig:"POSTGRES_DB" required:"true"`
	SSLMODE  string `envconfig:"POSTGRES_SSLMODE" required:"true"`
}

type Redis struct {
	Addr     string `envconfig:"REDIS_HOST" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" required:"true"`
	DB       int    `envconfig:"REDIS_DB" required:"true"`
}

// refactor later for Viper
// https://github.com/techschool/simplebank/blob/master/util/config.go
// https://maneeshaindrachapa.medium.com/go-with-env-files-using-viper-1eb3d1d1d221

func LoadEnvVariables() (ServerConfig, error) {
	cnf := ServerConfig{}
	err := envconfig.Process("", &cnf)
	return cnf, err
}
