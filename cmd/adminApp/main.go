package main

import "github.com/dmRusakov/tonoco/pkg/common/logging"

func main() {
	// start web server
	err = app.webServer.Start(ctx)
	if err != nil {
		logging.WithError(ctx, err).Fatal("webServer.Start")
	}
}