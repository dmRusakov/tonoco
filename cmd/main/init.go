package main

import (
	"github.com/dmRusakov/monkeysmoon-admin/internal/config"
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
)

func init() {
	// get config (read ENV variables)
	app.cfg = config.GetConfig()

	// make logger (logrus)
	logrus.Init()
	logger := logrus.GetLogrus()
	app.logger = &logger
	app.logger.Info("logger initialized")

}
