package main

import (
	"context"
	"ecom/internal/client"
	"ecom/internal/config"
	"ecom/internal/logger"
	"ecom/internal/server"

	"go.uber.org/zap"
)

func main() {
	logger := logger.NewLogger()
	zap.ReplaceGlobals(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		zap.L().Fatal(
			"failed to load config",
			zap.String("action", "shutdown"),
			zap.Error(err),
		)
	}

	client, err := client.NewClient(context.Background(), cfg)
	if err != nil {
		zap.L().Fatal(
			"failed to create client",
			zap.String("action", "shutdown"),
			zap.Error(err),
		)
	}

	if err := server.NewServer(
		cfg.Server,
		cfg.Hashing,
		cfg.JWT,
		client,
	).Start(); err != nil {
		zap.L().Fatal(
			"failed to start the server",
			zap.String("action", "shutdown"),
			zap.Error(err),
		)
	}
}
