package main

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/common/logging"
)

var app *appInit.App
var err error
var ctx context.Context

func init() {
	// make context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	logging.L(ctx).Info("config initializing")
	cfg := config.GetConfig(ctx)

	// save logger to context
	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	// new app init
	logging.L(ctx).Info("app initializing")
	app = appInit.NewAppInit(ctx, cfg)

	// appCacheService
	logging.L(ctx).Info("appCacheService initializing")
	err = app.AppCacheServiceInit()
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.AppCacheServiceInit")
	}

	// userCacheService
	logging.L(ctx).Info("UserCacheService initializing")
	err = app.UserCacheServiceInit()
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.UserCacheServiceInit")
	}

	// productPolicy
	logging.L(ctx).Info("product API initializing")
	//err = app.ProductAPIInit()
	//if err != nil {
	//	logging.WithError(ctx, err).Fatal("app.ProductPolicyInit")
	//}

	//fmt.Println(productPolicy)

	// web router
	app.WebServer, err = web.NewWebServer()
}
