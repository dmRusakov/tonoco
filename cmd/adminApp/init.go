package main

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/common/logging"
	"github.com/dmRusakov/tonoco/pkg/redisdb"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
)

func init() {
	// make context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	logging.L(ctx).Info("config initializing")
	app.cfg = config.GetConfig(ctx)

	// save logger to context
	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	// cache storage
	logging.L(ctx).Info("cache storage initializing")
	app.cacheStorage, err = redisdb.Connect(context.Background(), app.cfg.CacheStorage.ToRedisConfig())
	if err != nil {
		logging.WithError(ctx, err).Fatal("redisdb.Connect")
	}

	// appCacheService
	logging.L(ctx).Info("appCacheService initializing")

	app.appCacheService, err = appCacheService.NewCacheService("app")
	if err != nil {
		logging.WithError(ctx, err).Fatal("appCacheService.NewCacheService")
	}

	// UserCacheService
	logging.L(ctx).Info("UserCacheService initializing")
	app.userCacheService, err = userCacheService.NewCacheService("user")
	if err != nil {
		logging.WithError(ctx, err).Fatal("userCacheService.NewCacheService")
	}

	// web router
	app.webServer, err = web.NewWebServer()
}
