package admin_app_web_v1

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/config"
	"net/http"
)

type Route struct {
	page     string
	template string
}

func (s Service) Start(ctx context.Context, cfg *config.Config) error {
	// static files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up a custom HTTP Service to handle .wasm.js files
	http.HandleFunc("/assets/wasm/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// router
	s.router(ctx, cfg)

	// start Service
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.cfg.WebPort), nil)
	if err != nil {
		return err
	}

	return nil
}
