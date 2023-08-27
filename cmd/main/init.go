package main

import "github.com/dmRusakov/monkeysmoon-admin/internal/config"

func init() {
	// get config
	app.cfg = config.GetConfig()
}
