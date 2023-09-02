package main

import (
	appData "github.com/dmRusakov/monkeysmoon-admin/internal/app"
	"github.com/dmRusakov/monkeysmoon-admin/internal/config"
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
)

var app = appData.AppData{}

func init() {
	//var err error

	// get config (read ENV variables)
	app.Cfg = config.GetConfig()

	// make logger (logrus)
	logrus.Init()
	logger := logrus.GetLogrus()
	app.Logger = &logger
	app.Logger.Info("logger initialized")

	// connect to DataStorage (postgres)
	//app.DataStorage, err = postgresdb.Connect(context.Background(), app.Cfg)

}
