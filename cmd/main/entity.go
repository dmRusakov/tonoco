package main

import "github.com/dmRusakov/monkeysmoon-admin/internal/config"

type App struct {
	cfg *config.Config
}

var app App = App{}
