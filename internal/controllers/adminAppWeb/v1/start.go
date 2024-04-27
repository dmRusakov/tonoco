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

func (s server) Start(ctx context.Context, port string) error {
	// static files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up a custom HTTP server to handle .wasm.js files
	http.HandleFunc("/assets/wasm/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// router
	s.router()

	// start server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		return err
	}

	return nil
}
