package admin_app_web_v1

import (
	"context"
	"fmt"
	"net/http"
)

type Route struct {
	page     string
	template string
}

func (c Controller) Start(ctx context.Context) error {
	// static files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up a custom HTTP Controller to handle .wasm.js files
	http.HandleFunc("/assets/wasm/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// router
	c.router(ctx)

	// start Controller
	err := http.ListenAndServe(fmt.Sprintf(":%s", c.cfg.AppWebPort), nil)
	if err != nil {
		return err
	}

	return nil
}
