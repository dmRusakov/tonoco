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

	// dashboard
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "dashboard.page.gohtml", appData)
	})

	// order
	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "orders.page.gohtml", appData)
	})
	http.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "order.page.gohtml", appData)
	})

	// product
	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		c.RenderProducts(r.Context(), w, r, appData)
	})
	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "product.page.gohtml", appData)
	})

	// category
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "categories.page.gohtml", appData)
	})
	http.HandleFunc("/category/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "category.page.gohtml", appData)
	})

	// tag
	http.HandleFunc("/specifications/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "specifications.page.gohtml", appData)
	})
	http.HandleFunc("/specification/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "tag.page.gohtml", appData)
	})

	// page
	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "pages.page.gohtml", appData)
	})
	http.HandleFunc("/page/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "page.page.gohtml", appData)
	})

	// integration
	http.HandleFunc("/integration/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "integration.page.gohtml", appData)
	})

	// coupon
	http.HandleFunc("/coupons/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "coupons.page.gohtml", appData)
	})
	http.HandleFunc("/coupon/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "coupon.page.gohtml", appData)
	})

	// media
	http.HandleFunc("/media/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "media.page.gohtml", appData)
	})

	// settings
	http.HandleFunc("/settings/", func(w http.ResponseWriter, r *http.Request) {
		c.Render(w, "settings.page.gohtml", appData)
	})
}
