package appInit

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	product_policy "github.com/dmRusakov/tonoco/internal/domain/policy/product"
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
	WebServer web.Server

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
	productPolicy *product_policy.Policy
	//ProductApi    pb_prod_products.TonocoProductServiceClient
}
