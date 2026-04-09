package service

import "fwd/internal/repo"

type Service struct {
	Account *AccountService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Account: NewAccountService(repo.Account),
	}
}
