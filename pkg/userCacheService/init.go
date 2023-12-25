package userCacheService

import "github.com/go-redis/redis/v8"

type UserCacheService struct {
	client     *redis.Client
	authPrefix string
}

type userCacheService interface {
	GetAuthPrefix() string
}

func NewCacheService(client *redis.Client, authPrefix string) (*UserCacheService, error) {
	return &UserCacheService{
		client:     client,
		authPrefix: authPrefix,
	}, nil
}
