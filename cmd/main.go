package main

import (
	"fwd/internal/client"
	"fwd/internal/config"
	"fwd/internal/logger"
	"fwd/internal/server"

	"go.uber.org/zap"
)

func main() {
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		zap.L().Fatal(
			"failed to load config file",
			zap.String("operation", "config.LoadConfig"),
			zap.String("impact", "stoping the server"),
			zap.Error(err))
	}

	client, err := client.CreateClient(cfg.Postgres)
	if err != nil {
		zap.L().Fatal("failed to create client", zap.String("operation", "Client.CreateClient"), zap.String("impact", "stoping the server"), zap.Error(err))
	}

	if err := server.NewServer(cfg.Server, client).Start(); err != nil {
		zap.L().Fatal(
			"failed to start the server",
			zap.String("operation", "server.Start"),
			zap.String("impact", "stoping the server"),
			zap.Error(err),
		)
	}
}
