package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" env-default:"gateway-service"`
	AppEnv      string `env:"APP_ENV" env-default:"development"`
	Host        string `env:"GRPC_HOST" env-default:"localhost"`
	Port        int    `env:"GRPC_PORT" env-default:"50051"`
	LogLevel    string `env:"LOG_LEVEL" env-default:"info"`

	AccountGrpcHost     string `env:"ACCOUNT_GRPC_HOST"`
	AuthGrpcHost        string `env:"AUTH_GRPC_HOST"`
	TransactionGrpcHost string `env:"TRANSACTION_GRPC_HOST"`
	JwtSecret           string `env:"JWT_SECRET"`
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
