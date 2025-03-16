package configs

import "github.com/spf13/viper"

type conf struct {
	PostgresHost         string `mapstructure:"POSTGRES_HOST"`
	PostgresPort         string `mapstructure:"POSTGRES_PORT"`
	PostgresUser         string `mapstructure:"POSTGRES_USER"`
	PostgresPassword     string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase     string `mapstructure:"POSTGRES_DATABASE"`
	JwtSecret            string `mapstructure:"JWT_SECRET"`
	RabbitMQUrl          string `mapstructure:"RABBITMQ_URL"`
	FootballApiUrl       string `mapstructure:"FOOTBALL_API_URL"`
	FootballApiAuthToken string `mapstructure:"FOOTBALL_API_AUTH_TOKEN"`
}

func LoadConfig() (*conf, error) {
	var cfg *conf

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./app/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
