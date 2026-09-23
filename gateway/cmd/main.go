package main

import (
	"context"
	"gateway/cmd/app"
	"gateway/internal/config"
	"gateway/internal/logger"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	l := logger.New(cfg)

	application := app.New(&l, cfg)

	if err := application.Run(ctx); err != nil {
		l.Fatal().Err(err).Msg("error")
	}
}
