package main

import (
	"github.com/dmRusakov/monkeysmoon-admin/internal/config"
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
)

type App struct {
	cfg    *config.Config
	logger *logrus.Logrus
}

var app App = App{}
