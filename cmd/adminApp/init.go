package main

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	web_v1 "github.com/dmRusakov/tonoco/internal/controllers/adminAppWeb/v1"
	"github.com/dmRusakov/tonoco/pkg/common/logging"
)

var app *appInit.App
var err error
var ctx context.Context

func init() {
	// make context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logging.L(ctx).Info("Init App")

	// config
	cfg := config.GetConfig(ctx)
	logging.L(ctx).Info("Config initialized")

	// save logger to context
	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	// new app init
	app = appInit.NewAppInit(ctx, cfg)
	logging.L(ctx).Info("App initialized")

	// app cache service (redis)
	err = app.AppCacheServiceInit()
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.AppCacheServiceInit")
	}
	logging.L(ctx).Info("App Cache Service initialized")

	// user cache service (redis)
	err = app.UserCacheServiceInit()
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.UserCacheServiceInit")
	}
	logging.L(ctx).Info("UserCacheService initialized")

	// application database (sqlDB) (postgresql)
	err := app.SqlDBInit()
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.SqlDBInit")
	}
	logging.L(ctx).Info("Product DB initialized")

	// product use case
	err = app.ProductUseCaseInit(cfg)
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.ProductUseCaseInit")
	}
	logging.L(ctx).Info("Product UseCase initialized")

	// product api server (controller, getter)
	err = app.ProductControllerGetterInit()
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.ProductControllerGetterInit")
	}
	logging.L(ctx).Info("Product Controller (Getter) initialized")

	// adminAppWeb server
	app.WebServer, err = web_v1.NewWebServer(cfg, app.ProductUseCase, app.AppCacheService, app.UserCacheService)
	if err != nil {
		logging.WithError(ctx, err).Fatal("web_v1.NewWebServer")
	}
	logging.L(ctx).Info("Web Server initialized")
}
