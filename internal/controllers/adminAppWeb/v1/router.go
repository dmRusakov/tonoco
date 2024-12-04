package admin_app_web_v1

import (
	"context"
	"fmt"
	"net/http"
)

func (c Controller) router(ctx context.Context) {
	routes := []struct {
		path    string
		handler func(http.ResponseWriter, *http.Request)
	}{
		// home page
		{"/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "dashboard.page.gohtml") }},

		// shop pages
		{fmt.Sprintf("/%s/", c.cfg.ShopPageUrl), func(w http.ResponseWriter, r *http.Request) { c.RenderShopPage(r.Context(), w, r, c.cfg.ShopPageUrl) }},

		{"/orders/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "orders.page.gohtml") }},
		{"/order/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "order.page.gohtml") }},

		{"/product/", func(w http.ResponseWriter, r *http.Request) { c.Render(w, "product.page.gohtml") }},
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
