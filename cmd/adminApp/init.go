package main

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	web_v1 "github.com/dmRusakov/tonoco/internal/controllers/web/v1"
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
	logging.L(ctx).Info("config initialized")

	// save logger to context
	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	// new app init
	app = appInit.NewAppInit(ctx, cfg)
	logging.L(ctx).Info("app initialized")

	// appCacheService
	err = app.AppCacheServiceInit()
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.AppCacheServiceInit")
	}
	logging.L(ctx).Info("appCacheService initialized")

	// userCacheService
	err = app.UserCacheServiceInit()
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.UserCacheServiceInit")
	}
	logging.L(ctx).Info("UserCacheService initialized")

	// productPolicy

	//fmt.Println(productPolicy)

	// web router
	app.WebServer, err = web_v1.NewWebServer()
	if err != nil {
		logging.WithError(ctx, err).Fatal("web_v1.NewWebServer")
	}
}
