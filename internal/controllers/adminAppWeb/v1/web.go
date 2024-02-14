package admin_app_web_v1

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
	Start(ctx context.Context, port string) error
}

type Route struct {
	Path string
	File string
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

func (s server) Start(ctx context.Context, port string) error {
	// static files
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Set up a custom HTTP server to handle .wasm.js files
	http.HandleFunc("/assets/wasm/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	routes := []Route{
		{Path: "/", File: "dashboard.page.gohtml"},
		{Path: "/order/", File: "order.page.gohtml"},
		{Path: "/orders/", File: "orders.page.gohtml"},
		{Path: "/product/", File: "product.page.gohtml"},
		{Path: "/products/", File: "products.page.gohtml"},
		{Path: "/category/", File: "category.page.gohtml"},
		{Path: "/categories/", File: "categories.page.gohtml"},
		{Path: "/specification/", File: "specification.page.gohtml"},
		{Path: "/specifications/", File: "specifications.page.gohtml"},
		{Path: "/page/", File: "page.page.gohtml"},
		{Path: "/pages/", File: "pages.page.gohtml"},
		{Path: "/integration/", File: "integration.page.gohtml"},
		{Path: "/coupon/", File: "coupon.page.gohtml"},
		{Path: "/coupons/", File: "coupons.page.gohtml"},
		{Path: "/media/", File: "media.page.gohtml"},
		{Path: "/settings/", File: "settings.page.gohtml"},
	}

	for _, route := range routes {
		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			s.Render(w, route.File)
		})
	}

	addr := fmt.Sprintf(":%s", port)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}

	return nil
}
