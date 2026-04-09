package repo

import (
	"context"
	"errors"
	"fwd/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AccountRepo struct {
	pgPool *pgxpool.Pool
}

func NewAccountRepo(pgPool *pgxpool.Pool) *AccountRepo {
	return &AccountRepo{
		pgPool: pgPool,
	}
}

func (r *AccountRepo) GetAccountByEmail(
	ctx context.Context,
	email string,
) (*model.AccountModel, error) {
	query := `
		SELECT *
		FROM accounts
		WHERE email = @email
	`

	rows, err := r.pgPool.Query(ctx, query, pgx.NamedArgs{
		"email": email,
	})
	if err != nil {
		zap.L().Error(
			"postgres query failed",
			zap.String("operation", "AccountRepo.GetAccountByEmail"),
			zap.Error(err),
		)
		return nil, err
	}

	defer rows.Close()

	acc, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[model.AccountModel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			zap.L().Info(
				"account not found",
				zap.String("operation", "AccountRepo.GetAccountByEmail"),
				zap.String("email", email),
			)
			return nil, nil
		}
		zap.L().Error(
			"error collecting rows",
			zap.String("operation", "AccountRepo.GetAccountByEmail"),
			zap.Error(err),
		)
		return nil, err
	}

	return acc, nil
}

func (r *AccountRepo) CreateAccount(
	ctx context.Context,
	email string, passwordHash string,
) error {
	query := `
		INSERT INTO accounts (email, password)
		VALUES (@email, @passwordHash)
	`

	cmdTg, err := r.pgPool.Exec(context.Background(), query, pgx.NamedArgs{
		"email":        email,
		"passwordHash": passwordHash,
	})
	if err != nil {
		zap.L().Error(
			"postgres query failed",
			zap.String("operation", "AccountRepo.CreateAccount"),
			zap.Error(err),
		)
		return err
	}

	if cmdTg.RowsAffected() != 1 {
		zap.L().Error(
			"unexpected number of rows affected",
			zap.String("operation", "AccountRepo.CreateAccount"),
			zap.Int64("rowsAffected", cmdTg.RowsAffected()),
			zap.Error(err),
		)
		return err
	}

	return nil
}
