package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/entity"
	"net/http"
)

func (c Controller) router(ctx context.Context, cfg *config.Config) {

	appData := entity.AppData{
		IsProd:  cfg.IsProd,
		IsDebug: cfg.IsDebug,
	}

	routes := []struct {
		path    string
		handler func(http.ResponseWriter, *http.Request)
	}{
		{"/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "dashboard.page.gohtml", appData) }},
		{"/orders/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "orders.page.gohtml", appData) }},
		{"/order/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "order.page.gohtml", appData) }},
		{"/products/", func(w http.ResponseWriter, r *http.Request) { c.RenderProducts(r.Context(), w, r, appData) }},
		{"/range-hood/", func(w http.ResponseWriter, r *http.Request) { c.RenderProducts(r.Context(), w, r, appData) }},
		{"/product/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "product.page.gohtml", appData) }},
		{"/categories/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "categories.page.gohtml", appData) }},
		{"/category/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "category.page.gohtml", appData) }},
		{"/specifications/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "specifications.page.gohtml", appData) }},
		{"/specification/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "tag.page.gohtml", appData) }},
		{"/pages/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "pages.page.gohtml", appData) }},
		{"/page/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "page.page.gohtml", appData) }},
		{"/integration/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "integration.page.gohtml", appData) }},
		{"/coupons/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "coupons.page.gohtml", appData) }},
		{"/coupon/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "coupon.page.gohtml", appData) }},
		{"/media/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "media.page.gohtml", appData) }},
		{"/settings/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "settings.page.gohtml", appData) }},
	}

	for _, route := range routes {
		http.HandleFunc(route.path, route.handler)
	}
}
