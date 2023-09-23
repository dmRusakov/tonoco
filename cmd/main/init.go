package main

import (
	"github.com/dmRusakov/monkeysmoon-admin/controllers/web"
	"github.com/dmRusakov/monkeysmoon-admin/internal/config"
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
)

var App = AppData{}

func init() {
	//var err error

	// get config (read ENV variables)
	App.Cfg = config.GetConfig()

	// make logger (logrus)
	logrus.Init()
	logger := logrus.GetLogrus()
	App.Logger = &logger
	App.Logger.Info("logger initialized")

	// web router
	App.Router.Web, _ = web.NewWebServer(App.Logger)

	// connect to DataStorage (postgres)
	//app.DataStorage, err = postgresdb.Connect(context.Background(), app.Cfg)

}
