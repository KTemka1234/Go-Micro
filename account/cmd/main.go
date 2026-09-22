package main

import (
	"account/cmd/app"
	"account/internal/config"
	"account/internal/logger"
	_ "account/internal/migrations"
	"context"

	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log := logger.New(cfg)

	ctx := context.Background()
	app := app.New(&log, cfg)
	if err := app.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to start service")
	}
}
