package main

import (
	"account/internal/config"
	"account/internal/logger"
	_ "account/internal/migrations"
	"account/internal/repository"
	"account/internal/server"
	"account/internal/service"

	"database/sql"
	"fmt"
	"log"
	"net"

	accountpb "github.com/KTemka1234/go-micro/contracts/account/go"

	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log := logger.New(cfg)

	db, err := gorm.Open(postgres.Open(cfg.DbDsn), &gorm.Config{})
	if err != nil {
		log.Error().Msgf("Failed to connect to database: %v", err)
		return
	}
	log.Info().Msg("database connected")

	migrationConn, err := db.DB()
	if migrationConn == nil {
		log.Fatal().Err(err).Msg("failed to select migrations dialect")
	}

	dbGoose, err := sql.Open("postgres", cfg.DbDsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create sql connection")
	}

	if err := goose.Up(dbGoose, "internal/migrations"); err != nil {
		log.Fatal().Err(err).Msg("failed to run up migrations")
	}

	repo := repository.New(db, &log)
	service := service.New(repo, &log)
	server := server.New(service, &log)

	s := grpc.NewServer()
	accountpb.RegisterAccountServer(s, server)

	listenAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Error().Msgf("Failed to listen on addr %s: %v", listenAddr, err)
		return
	}

	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthSrv)

	log.Info().Msgf("gRPC server listening on %s", listenAddr)
	if err := s.Serve(lis); err != nil {
		log.Error().Msgf("Failed to serve gRPC: %v", err)
		return
	}

	log.Info().Msg("service starting up")
}
