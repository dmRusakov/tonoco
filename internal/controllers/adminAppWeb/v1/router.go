package admin_app_web_v1

import (
	"context"
	"net/http"
)

func (c Controller) router(ctx context.Context) {
	routes := []struct {
		path    string
		handler func(http.ResponseWriter, *http.Request)
	}{
		{"/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "dashboard.page.gohtml") }},
		{"/orders/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "orders.page.gohtml") }},
		{"/order/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "order.page.gohtml") }},
		{"/products/", func(w http.ResponseWriter, r *http.Request) { c.RenderProducts(r.Context(), w, r) }},
		{"/range-hood", func(w http.ResponseWriter, r *http.Request) { c.RenderProducts(r.Context(), w, r) }},
		{"/product/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "product.page.gohtml") }},
		{"/categories/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "categories.page.gohtml") }},
		{"/category/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "category.page.gohtml") }},
		{"/specifications/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "specifications.page.gohtml") }},
		{"/specification/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "tag.page.gohtml") }},
		{"/pages/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "pages.page.gohtml") }},
		{"/page/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "page.page.gohtml") }},
		{"/integration/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "integration.page.gohtml") }},
		{"/coupons/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "coupons.page.gohtml") }},
		{"/coupon/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "coupon.page.gohtml") }},
		{"/media/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "media.page.gohtml") }},
		{"/settings/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "settings.page.gohtml") }},
	}

	for _, route := range routes {
		http.HandleFunc(route.path, route.handler)
	}
}
