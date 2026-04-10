package service

import (
	"context"
	"fwd/internal/model"
	"fwd/internal/repo"
)

type PlatformService struct {
	platformRepo *repo.PlatformRepo
}

func NewPlatformService(platformRepo *repo.PlatformRepo) *PlatformService {
	return &PlatformService{platformRepo: platformRepo}
}

func (s *PlatformService) GetAllPlatforms(
	ctx context.Context,
) ([]*model.PlatformModel, error) {
	platforms, err := s.platformRepo.GetAllPlatforms(ctx)
	if err != nil {
		return nil, err
	}

	return platforms, nil
}
