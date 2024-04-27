package appCacheService

import (
	"github.com/go-redis/redis/v8"
)

type Service struct {
	client     *redis.Client
	authPrefix *string
}

type service interface {
	GetAuthPrefix() string
}

func NewCacheService(client *redis.Client, authPrefix string) (*Service, error) {
	return &Service{
		client:     client,
		authPrefix: &authPrefix,
	}, nil
}
