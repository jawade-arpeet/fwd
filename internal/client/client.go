package client

import (
	"context"
	"fmt"
	"fwd/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Client struct {
	Postgres *pgxpool.Pool
}

func CreateClient(cfg *config.PostgresConfig) (*Client, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		zap.L().Error(
			"postgres connection failed",
			zap.String("operation", "client.CreateClient"),
			zap.Error(err),
		)
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		zap.L().Error(
			"failed to ping the postgres",
			zap.String("operation", "client.CreateClient"),
			zap.Error(err),
		)
		return nil, err
	}

	zap.L().Info(
		"postgres connection successful",
		zap.String("operation", "client.CreateClient"),
	)

	return &Client{
		Postgres: pool,
	}, nil
}
