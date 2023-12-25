package appCacheService

import (
	"github.com/go-redis/redis/v8"
)

type AppCacheService struct {
	client     *redis.Client
	authPrefix *string
}

type service interface {
	GetAuthPrefix() string
}

func NewCacheService(client *redis.Client, authPrefix string) (*AppCacheService, error) {
	return &AppCacheService{
		client:     client,
		authPrefix: &authPrefix,
	}, nil
}
