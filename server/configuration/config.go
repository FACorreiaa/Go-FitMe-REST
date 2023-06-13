package configuration

import "github.com/spf13/viper"

type Config struct {
	Mode     string `mapstructure:"mode"`
	Dotenv   string `mapstructure:"dotenv"`
	Handlers struct {
		ExternalApi struct {
			Port      string `mapstrucutre:"port"`
			CertFile  string `mapstructure:"certFile"`
			KeyFile   string `mapstructure:"keyFile"`
			EnableTLS bool   `mapstracture:"enableTLS"`
		} `mapstructure:"externalAPI"`
		Pprof struct {
			Port      string `mapstructure:"port"`
			CertFile  string `mapstructure:"certFile"`
			KeyFile   string `mapstructure:"keyFile"`
			EnableTLS bool   `mapstructure:"enableTLS"`
		}
		Prometheus struct {
			Port      string `mapstructure:"port"`
			CertFile  string `mapstructure:"certFile"`
			KeyFile   string `mapstructure:"keyFile"`
			EnableTLS bool   `mapstructure:"enableTLS"`
		}
	} `mapstructure:"handlers"`
	Repositories struct {
		Postgres struct {
			Host               string `mapstructure:"host"`
			Port               string `mapstructure:"port"`
			Username           string `mapstructure:"username"`
			DB                 string `mapstructure:"db"`
			SSLMode            string `mapstructure:"sslmode"`
			MaxConnWaitingTime int    `mapstructure:"maxConnWaitingTime"`
		}
	}
}

func InitConfig() (Config, error) {

	var config Config
	v := viper.New()
	v.AddConfigPath("/server/configuration")
	v.AddConfigPath("/app/server/configuration")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return config, err
	}
	if err := v.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
