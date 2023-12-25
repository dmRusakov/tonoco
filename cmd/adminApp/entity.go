package main

import (
	"context"
	"database/sql"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"github.com/julienschmidt/httprouter"
)

var app App = App{}
var err error
var ctx context.Context

type App struct {
	cfg       *config.Config
	router    *httprouter.Router
	webServer web.Server

	// cache
	appCacheService  *appCacheService.AppCacheService
	userCacheService *userCacheService.UserCacheService

	// data storage
	dataStorage *sql.DB
}
