package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" env-default:"account-service"`
	AppEnv      string `env:"APP_ENV" env-default:"development"`
	Host        string `env:"GRPC_HOST" env-default:"localhost"`
	Port        int    `env:"GRPC_PORT" env-default:"50051"`
	LogLevel    string `env:"LOG_LEVEL" env-default:"info"`

	DbDsn string `env:"DB_DSN"`
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
