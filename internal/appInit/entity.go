package appInit

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	Ctx       *context.Context
	Cfg       *config.Config
	Router    *httprouter.Router
	WebServer web.Server

	// cache
	AppCacheService  *appCacheService.AppCacheService
	UserCacheService *userCacheService.UserCacheService

	// data
}
