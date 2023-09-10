package app

import (
	"github.com/dmRusakov/monkeysmoon-admin/internal/config"
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
)

type AppData struct {
	Cfg    *config.Config
	Logger *logrus.Logrus
}
