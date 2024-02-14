package main

import "github.com/dmRusakov/tonoco/pkg/common/logging"

func main() {
	// start adminAppWeb server
	err = app.WebServer.Start(ctx, app.Cfg.WebPort)
	if err != nil {
		logging.WithError(ctx, err).Fatal("webServer.Start")
	}
}
