package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity"
	"html/template"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

// URL params for products page

func (s server) RenderProducts(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
) {

	var wg sync.WaitGroup
	var tmpl *template.Template
	var params *entity.ProductsPgeUrlParams

	wg.Add(2)

	go func() {
		defer wg.Done()
		tmpl = s.makeTemplate("products.page.gohtml")
	}()

	go func() {
		defer wg.Done()
		params = s.ReadProductsUrlParam(r)
	}()

	wg.Wait()

	// get products
	products, err := s.productUseCase.GetProductList(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productPage := entity.ProductPage{
		Name: "Range Hoods",

		Products:   products,
		ProductUrl: "range-hood",
	}

	// render page
	if err := tmpl.Execute(w, productPage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

// ReadProductsUrlParam read page parameters from url
func (s server) ReadProductsUrlParam(r *http.Request) *entity.ProductsPgeUrlParams {
	params := &entity.ProductsPgeUrlParams{}
	v := reflect.ValueOf(params).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		value := r.URL.Query().Get(strings.ToLower(fieldName))
		if value != "" {
			field := v.Field(i)
			if field.Kind() == reflect.Ptr && !field.IsNil() {
				field = field.Elem()
			}
			if field.Kind() == reflect.String {
				field.SetString(value)
			}
		}
	}

	// default Currency
	if params.Currency == nil {
		params.Currency = entity.StringPtr("usd")
	}

	// default Page
	if params.Page == nil {
		params.Page = entity.Uint64Ptr(1)
	}

	// default PerPage
	if params.PerPage == nil {
		params.PerPage = entity.Uint64Ptr(18)
	}

	return params
}
