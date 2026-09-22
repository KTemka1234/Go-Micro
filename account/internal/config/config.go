package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" env-default:"account-service"`
	AppEnv string `env:"APP_ENV" env-default:"development"`
	Host string `env:"HTTP_HOST" env-default:"localhost"`
	Port int `env:"HTTP_PORT" env-default:"9000"`
	LogLevel string `envv:"LOG_LEVEL" env-default:"info"`

	DbDsn       string `env:"DB_DSN"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}