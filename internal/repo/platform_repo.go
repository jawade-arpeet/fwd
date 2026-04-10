package repo

import (
	"context"
	"fwd/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PlatformRepo struct {
	pgPool *pgxpool.Pool
}

func NewPlatformRepo(pgPool *pgxpool.Pool) *PlatformRepo {
	return &PlatformRepo{
		pgPool: pgPool,
	}
}

func (r *PlatformRepo) GetAllPlatforms(ctx context.Context) ([]*model.PlatformModel, error) {
	query := `
		SELECT
			id,
			name,
			logo_url
		FROM platforms
	`

	rows, err := r.pgPool.Query(ctx, query)
	if err != nil {
		zap.L().Error(
			"postgres query failed",
			zap.String("operation", "PlatformRepo.GetAllPlatforms"),
			zap.Error(err),
		)
		return nil, err
	}

	defer rows.Close()

	platforms, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[model.PlatformModel])
	if err != nil {
		zap.L().Error(
			"error collecting rows",
			zap.String("operation", "PlatformRepo.GetAllPlatforms"),
			zap.Error(err),
		)
		return nil, err
	}

	return platforms, nil
}
