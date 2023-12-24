package main

import (
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/logrus"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"github.com/go-redis/redis/v8"
)

type AppData struct {
	Cfg    *config.Config
	Logger *logrus.Logrus
	Router Router

	// cache
	CacheStorage     *redis.Client
	AppCacheService  appCacheService.AppCacheService
	UserCacheService userCacheService.UserCacheService
}

type Router struct {
	Web web.Server
}
