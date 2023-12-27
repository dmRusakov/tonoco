package main

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/appInit"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	product_policy "github.com/dmRusakov/tonoco/internal/domain/policy/product"
	product_dao "github.com/dmRusakov/tonoco/internal/domain/product/dao"
	product_service "github.com/dmRusakov/tonoco/internal/domain/product/service"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/common/core/clock"
	"github.com/dmRusakov/tonoco/pkg/common/core/closer"
	"github.com/dmRusakov/tonoco/pkg/common/core/identity"
	"github.com/dmRusakov/tonoco/pkg/common/logging"
	"github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/redisdb"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"time"
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
	app = appInit.NewAppInit(&ctx, cfg)
	fmt.Println(app)

	// cache storage (Redis)
	logging.L(ctx).Info("cache storage initializing")
	cacheStorage, err := redisdb.Connect(context.Background(), app.Cfg.CacheStorage.ToRedisConfig())
	if err != nil {
		logging.WithError(ctx, err).Fatal("redisdb.Connect")
	}

	// appCacheService
	logging.L(ctx).Info("appCacheService initializing")

	app.AppCacheService, err = appCacheService.NewCacheService(cacheStorage, "app")
	if err != nil {
		logging.WithError(ctx, err).Fatal("appCacheService.NewCacheService")
	}

	// userCacheService
	logging.L(ctx).Info("UserCacheService initializing")
	app.UserCacheService, err = userCacheService.NewCacheService(cacheStorage, "user")
	if err != nil {
		logging.WithError(ctx, err).Fatal("userCacheService.NewCacheService")
	}

	// data storage (PostgreSQL)
	logging.L(ctx).Info("data storage initializing")
	dataClient, err := postgresql.NewClient(ctx, 5, 3*time.Second, app.Cfg.DataStorage.ToPostgreSQLConfig(), false)
	if err != nil {
		logging.WithError(ctx, err).Fatal("postgresql.NewClient")
	}

	closer.AddN(dataClient)

	cl := clock.New()
	generator := identity.NewGenerator()

	productStorage := product_dao.NewProductStorage(dataClient)
	productService := product_service.NewProductService(productStorage)
	productPolicy := product_policy.NewProductPolicy(productService, generator, cl)
	//productServiceServer := product_api_v1.NewServer(
	//	productPolicy,
	//	product_constact.UnimplementedTonocoProductServiceServer{},
	//)

	fmt.Println(productPolicy)

	// web router
	app.WebServer, err = web.NewWebServer()
}
