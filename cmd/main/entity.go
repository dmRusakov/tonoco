package main

import (
	"github.com/dmRusakov/monkeysmoon-admin/controllers/web"
	"github.com/dmRusakov/monkeysmoon-admin/internal/config"
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
)

type AppData struct {
	Cfg    *config.Config
	Logger *logrus.Logrus
	Router Router
}

type Router struct {
	Web web.Server
}
