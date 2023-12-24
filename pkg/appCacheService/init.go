package appCacheService

import (
	"github.com/dmRusakov/tonoco/pkg/logrus"
)

var _ AppCacheService = &appCacheService{}

type appCacheService struct {
	log        *logrus.Logrus
	authPrefix *string
}

type AppCacheService interface {
	GetAuthPrefix() string
}

func NewCacheService(log *logrus.Logrus, authPrefix string) (*appCacheService, error) {
	return &appCacheService{
		log,
		&authPrefix,
	}, nil
}
