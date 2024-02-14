package appInit

import (
	"context"
	productServire "github.com/dmRusakov/tonoco-grpc/gen/go/proto/service/v1"
	"github.com/dmRusakov/tonoco/internal/config"
	webController "github.com/dmRusakov/tonoco/internal/controllers/adminAppWeb/v1"
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
	CacheDB *redis.Client
	SqlDB   *pgxpool.Pool

	// cache
	AppCacheService  *appCacheService.AppCacheService
	UserCacheService *userCacheService.UserCacheService

	// usaCase
	ProductUseCase *productUsecase.UseCase

	// api controllers
	ProductController *productServire.ProductServiceServer
}
