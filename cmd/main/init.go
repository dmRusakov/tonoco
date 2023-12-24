package main

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/redisdb"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"

	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/pkg/logrus"
)

var App = &AppData{}

func init() {
	var err error

	// config (read ENV variables)
	App.Cfg = config.GetConfig()

	// logger (logrus)
	logrus.Init()
	logger := logrus.GetLogrus()
	App.Logger = &logger
	App.Logger.Info("logger initialized")

	// cache storage
	App.CacheStorage, err = redisdb.Connect(context.Background(), App.Cfg.CacheStorage.ToRedisConfig())
	if err != nil {
		App.Logger.Fatal(err)
	}
	App.Logger.Info("CacheStorage initialized")

	// AppCacheService
	App.AppCacheService, err = appCacheService.NewCacheService(App.Logger, "app")
	if err != nil {
		App.Logger.Fatal(err)
	}
	App.Logger.Info("AppCacheService initialized")

	// UserCacheService
	App.UserCacheService, err = userCacheService.NewCacheService(App.Logger, "user")
	if err != nil {
		App.Logger.Fatal(err)
	}
	App.Logger.Info("UserCacheService initialized")

	// web router
	App.Router.Web, _ = web.NewWebServer(App.Logger)

}
