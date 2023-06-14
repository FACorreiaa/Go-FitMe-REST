package internals

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type ServerConfig struct {
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

// refactor later for Viper
// https://github.com/techschool/simplebank/blob/master/util/config.go
// https://maneeshaindrachapa.medium.com/go-with-env-files-using-viper-1eb3d1d1d221
func LoadEnvVariables() (ServerConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment config! %v", err)
	} else {
		fmt.Println("Config loaded!")
	}
	cnf := ServerConfig{}
	err = envconfig.Process("", &cnf)
	return cnf, err
}
