package cacheService

import (
	"github.com/dmRusakov/tonoco/internal/config"
)

var _ CacheService = &cacheService{}

type cacheService struct {
	App        *config.AppData
	authPrefix *string
}

type CacheService interface {
	GetAuthPrefix() string
}

func NewCacheService(app *config.AppData, authPrefix string) (*cacheService, error) {
	return &cacheService{
		app,
		&authPrefix,
	}, nil
}
