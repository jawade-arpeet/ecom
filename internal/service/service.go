package service

import (
	"ecom/internal/config"
	"ecom/internal/repo"
)

type Service struct {
}

func NewService(
	repo *repo.Repo,
	hashingCfg *config.HashingConfig,
	jwtCfg *config.JWTConfig,
) *Service {
	return &Service{}
}
