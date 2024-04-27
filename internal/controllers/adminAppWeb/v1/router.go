package admin_app_web_v1

import "net/http"

func (s server) router() {
	// dashboard
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "dashboard.page.gohtml")
	})

	// order
	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "orders.page.gohtml")
	})
	http.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "order.page.gohtml")
	})

	// product
	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		s.RenderProducts(w, r)
	})
	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "product.page.gohtml")
	})

	// category
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "categories.page.gohtml")
	})
	http.HandleFunc("/category/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "category.page.gohtml")
	})

	// specification
	http.HandleFunc("/specifications/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "specifications.page.gohtml")
	})
	http.HandleFunc("/specification/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "specification.page.gohtml")
	})

	// page
	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "pages.page.gohtml")
	})
	http.HandleFunc("/page/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "page.page.gohtml")
	})

	// integration
	http.HandleFunc("/integration/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "integration.page.gohtml")
	})

	// coupon
	http.HandleFunc("/coupons/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "coupons.page.gohtml")
	})
	http.HandleFunc("/coupon/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "coupon.page.gohtml")
	})

	// media
	http.HandleFunc("/media/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "media.page.gohtml")
	})

	// settings
	http.HandleFunc("/settings/", func(w http.ResponseWriter, r *http.Request) {
		s.Render(w, "settings.page.gohtml")
	})
}
