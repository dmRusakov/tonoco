package main

import "github.com/dmRusakov/tonoco/pkg/common/logging"

func main() {
	// start adminAppWeb server
	err = app.WebServer.Start(ctx)
	if err != nil {
		logging.WithError(ctx, err).Fatal("webServer.Start")
	}
}
