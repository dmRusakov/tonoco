package appInit

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web/v1"
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
	WebServer v1.Server

	// helpers
	clock     clock.Clock
	generator *identity.Generator

	// db
	cacheDB *redis.Client
	sqlDB   *pgxpool.Pool

	// cache
	AppCacheService  *appCacheService.AppCacheService
	UserCacheService *userCacheService.UserCacheService

	// product
	//ProductApi    pb_prod_products.TonocoProductServiceClient
}
