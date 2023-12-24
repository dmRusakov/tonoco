package main

import (
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
)

var app App = App{}
var err error

type App struct {
	cfg       *config.Config
	router    *httprouter.Router
	webServer web.Server

	// cache
	cacheStorage     *redis.Client
	appCacheService  appCacheService.AppCacheService
	userCacheService userCacheService.UserCacheService
}
