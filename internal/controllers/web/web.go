package web

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
)

var _ Server = &server{}

func NewWebServer() (*server, error) {
	return &server{
		tmlPath: "./assets/templates/",
	}, nil
}

type server struct {
	tmlPath string
}

type Server interface {
	Render(w http.ResponseWriter, t string)
	Start(ctx context.Context) error
}

func (s server) Render(w http.ResponseWriter, pageTemplate string) {
	// page template
	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s%s", s.tmlPath, pageTemplate))

	// static templates
	partials := []string{
		"base.layout",
		"element/head_file_imports.partial",
		"element/head.partial",
		"element/menu.partial",
		"element/footer.partial",
		"element/footer_file_imports.partial",
	}
	for _, x := range partials {
		templateSlice = append(templateSlice, fmt.Sprintf("%s%s.gohtml", s.tmlPath, x))
	}

	// parse templates
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// render page
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

func (s server) Start(ctx context.Context) error {
	// static files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up a custom HTTP server to handle .wasm.js files
	http.HandleFunc("/assets/wasm/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// pages
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "dashboard.page.gohtml")
	})
	http.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "order.page.gohtml")
	})
	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "orders.page.gohtml")
	})
	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "product.page.gohtml")
	})
	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "products.page.gohtml")
	})
	http.HandleFunc("/category/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "category.page.gohtml")
	})
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "categories.page.gohtml")
	})
	http.HandleFunc("/specification/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "specification.page.gohtml")
	})
	http.HandleFunc("/specifications/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "specifications.page.gohtml")
	})
	http.HandleFunc("/page/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "page.page.gohtml")
	})
	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "pages.page.gohtml")
	})
	http.HandleFunc("/integration/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "integration.page.gohtml")
	})
	http.HandleFunc("/coupon/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "coupon.page.gohtml")
	})
	http.HandleFunc("/coupons/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "coupons.page.gohtml")
	})
	http.HandleFunc("/media/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "media.page.gohtml")
	})
	http.HandleFunc("/settings/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "settings.page.gohtml")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}
