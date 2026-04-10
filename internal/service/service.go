package service

import "fwd/internal/repo"

type Service struct {
	Account  *AccountService
	Platform *PlatformService
}

func NewService(repo *repo.Repo) *Service {
	return &Service{
		Account:  NewAccountService(repo.Account),
		Platform: NewPlatformService(repo.Platform),
	}
}
