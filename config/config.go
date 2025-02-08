package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort        string
	PostgreSqlDsn     string
	BasicAuthUsername string
	BasicAuthPassword string
}

func LoadConfig() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return &Config{
		ServerPort:        viper.GetString("SERVER_PORT"),
		PostgreSqlDsn:     viper.GetString("POSTGRESQL_DSN"),
		BasicAuthUsername: viper.GetString("BASIC_AUTH_USERNAME"),
		BasicAuthPassword: viper.GetString("BASIC_AUTH_PASSWORD"),
	}
}
