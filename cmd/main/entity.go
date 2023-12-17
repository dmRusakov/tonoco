package main

import (
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/logrus"
	"github.com/go-redis/redis/v8"
)

type AppData struct {
	Cfg    *config.Config
	Logger *logrus.Logrus
	Router Router

	// cache
	CacheStorage *redis.Client
}

type Router struct {
	Web web.Server
}
