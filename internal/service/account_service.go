package service

import (
	"context"
	"fwd/internal/errs"
	"fwd/internal/model"
	"fwd/internal/repo"

	"go.uber.org/zap"
)

type AccountService struct {
	accountRepo *repo.AccountRepo
}

func NewAccountService(repo *repo.AccountRepo) *AccountService {
	return &AccountService{
		accountRepo: repo,
	}
}

func (s *AccountService) SignUp(
	ctx context.Context,
	payload *model.SignUpPayload,
) error {
	acc, err := s.accountRepo.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		zap.L().Error(
			"error fetching account by email",
			zap.String("operation", "AccountService.SignUp"),
			zap.Error(err),
		)
		return err
	}

	if acc != nil {
		zap.L().Info(
			"account already exists with email",
			zap.String("operation", "AccountService.SignUp"),
			zap.String("email", payload.Email),
		)
		return errs.ErrAccountAlreadyExists
	}

	return s.accountRepo.CreateAccount(ctx, payload.Email, payload.Password)
}

func (s *AccountService) SignIn(
	ctx context.Context,
	payload *model.SignInPayload,
) error {
	acc, err := s.accountRepo.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		zap.L().Error(
			"error fetching account by email",
			zap.String("operation", "AccountService.SignIn"),
			zap.Error(err),
		)
		return err
	}

	if acc == nil {
		zap.L().Info(
			"account not found during sign-in",
			zap.String("operation", "AccountService.SignIn"),
		)
		return errs.ErrAccountDoesNotExists
	}

	if acc.Password != payload.Password {
		zap.L().Info(
			"invalid password during sign-in",
			zap.String("operation", "AccountService.SignIn"),
		)
		return errs.ErrInvalidPassword
	}

	return nil
}
