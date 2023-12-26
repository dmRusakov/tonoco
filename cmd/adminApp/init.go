package main

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/common/logging"
	"github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/redisdb"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"time"
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

	// cache storage (Redis)
	logging.L(ctx).Info("cache storage initializing")
	cacheStorage, err := redisdb.Connect(context.Background(), app.cfg.CacheStorage.ToRedisConfig())
	if err != nil {
		logging.WithError(ctx, err).Fatal("redisdb.Connect")
	}

	// appCacheService
	logging.L(ctx).Info("appCacheService initializing")

	app.appCacheService, err = appCacheService.NewCacheService(cacheStorage, "app")
	if err != nil {
		logging.WithError(ctx, err).Fatal("appCacheService.NewCacheService")
	}

	// userCacheService
	logging.L(ctx).Info("UserCacheService initializing")
	app.userCacheService, err = userCacheService.NewCacheService(cacheStorage, "user")
	if err != nil {
		logging.WithError(ctx, err).Fatal("userCacheService.NewCacheService")
	}

	// data storage (PostgreSQL)
	logging.L(ctx).Info("data storage initializing")
	dataStorage, err := postgresql.NewClient(ctx, 5, 3*time.Second, app.cfg.DataStorage.ToPostgreSQLConfig(), false)
	fmt.Println(dataStorage)

	// web router
	app.webServer, err = web.NewWebServer()
}
