package appInit

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	apiController "github.com/dmRusakov/tonoco/internal/controllers/api/v1"
	webController "github.com/dmRusakov/tonoco/internal/controllers/web/v1"
	productUsecase "github.com/dmRusakov/tonoco/internal/domain/useCase/product"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/common/core/clock"
	"github.com/dmRusakov/tonoco/pkg/common/core/identity"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	Ctx       context.Context
	Cfg       *config.Config
	Router    *httprouter.Router
	WebServer webController.Server

	// helpers
	clock     clock.Clock
	generator *identity.Generator

	// db
	cacheDB *redis.Client
	sqlDB   *pgxpool.Pool

	// cache
	AppCacheService  *appCacheService.AppCacheService
	UserCacheService *userCacheService.UserCacheService

	// usaCase
	ProductUseCase *productUsecase.UseCase

	// api controllers
	ProductController *apiController.ProductController
}