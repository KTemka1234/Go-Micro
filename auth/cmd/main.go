package main

import (
	"auth/cmd/app"
	"auth/internal/config"
	"auth/internal/logger"
	_ "auth/internal/migrations"
	"context"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.New(cfg)

	application := app.New(&log, cfg)

	if err := application.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("error")
	}
}
