package app

import (
	"github.com/dmRusakov/monkeysmoon-admin/internal/config"
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
	"github.com/go-redis/redis/v8"
)

type AppData struct {
	Cfg    *config.Config
	Logger *logrus.Logrus

	// Auth
	AuthStorage *redis.Client
	// AuthService authService.AuthService
}
