package main

import (
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/controllers/web"
	"github.com/dmRusakov/tonoco/pkg/logrus"
)

type AppData struct {
	Cfg    *config.Config
	Logger *logrus.Logrus
	Router Router
}

type Router struct {
	Web web.Server
}
