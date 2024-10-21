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

func (c Controller) RenderProducts(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	appData entity.AppData,
) {
	// add user to context TODO make it right
	ctx = context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

	var wg sync.WaitGroup
	var tmpl *template.Template
	var params *entity.ProductsPageUrlParams

	wg.Add(2)

	go func() {
		defer wg.Done()
		tmpl = c.makeTemplate("products.page.gohtml")
	}()

	go func() {
		defer wg.Done()
		params = c.ReadProductParam(r)
	}()

	wg.Wait()

	// get products
	appData.ConsoleMessage = entity.ConsoleMessage{}

	// get products
	products, err := c.productUseCase.GetProductList(ctx, params, &appData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// page info
	productPage := entity.ProductPage{
		Name: "Range Hoods",

		Products:   products,
		ProductUrl: "range-hood",

		Page:           params.Page,
		CountItems:     params.Count,
		ConsoleMessage: appData.ConsoleMessage,
	}

	// render page
	if err := tmpl.Execute(w, productPage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

// ReadProductParam read page parameters from url
func (c Controller) ReadProductParam(r *http.Request) *entity.ProductsPageUrlParams {
	// read url urlParams
	urlParams := &entity.ProductsPageUrlParams{}
	v := reflect.ValueOf(urlParams).Elem()
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
	if urlParams.Currency == nil {
		urlParams.Currency = entity.StringPtr("usd")
	}

	// default Page
	if urlParams.Page == nil {
		urlParams.Page = entity.Uint64Ptr(1)
	}

	// default PerPage
	if urlParams.PerPage == nil {
		urlParams.PerPage = entity.Uint64Ptr(18)
	}

	return urlParams
}

func (c Controller) addUserToContext(
	ctx context.Context,
	r *http.Request,
) (context.Context, error) {
	context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

	return ctx, nil
}
