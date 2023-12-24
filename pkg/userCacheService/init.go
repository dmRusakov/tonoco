package userCacheService

import (
	"github.com/dmRusakov/tonoco/pkg/logrus"
)

var _ UserCacheService = &userCacheService{}

type userCacheService struct {
	log        *logrus.Logrus
	authPrefix *string
}

type UserCacheService interface {
	GetAuthPrefix() string
}

func NewCacheService(log *logrus.Logrus, authPrefix string) (*userCacheService, error) {
	return &userCacheService{
		log,
		&authPrefix,
	}, nil
}
