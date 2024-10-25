package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/pagination"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
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
		params, _ = c.ReadProductParam(r)
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

		Items: products,
		Url:   "range-hood",

		Page:       uint64(*params.Page),
		PerPage:    uint64(*params.PerPage),
		TotalItems: uint64(*params.Count),
		TotalPages: uint64(((*params.Count) + (*params.PerPage) - 1) / *params.PerPage),

		ConsoleMessage: appData.ConsoleMessage,
	}

	productPage.Pagination = pagination.GetPagination(productPage.Page, productPage.TotalPages, 5)

	// render page
	if err := tmpl.Execute(w, productPage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

// ReadProductParam read page parameters from url
func (c Controller) ReadProductParam(r *http.Request) (*entity.ProductsPageUrlParams, string) {
	// read url urlParams
	urlParams := &entity.ProductsPageUrlParams{}
	v := reflect.ValueOf(urlParams).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		value := r.URL.Query().Get(strings.ToLower(fieldName))
		if value != "" {
			field := v.Field(i)
			if field.Kind() == reflect.Ptr {
				if field.IsNil() {
					field.Set(reflect.New(field.Type().Elem()))
				}
				field = field.Elem()
			}
			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Uint64:
				if parsedValue, err := strconv.ParseUint(value, 10, 64); err == nil {
					field.SetUint(parsedValue)
				}
			}
		}
	}

	// get ase url without excitation
	url := strings.Split(r.URL.String(), "?")[0]

	urlExcitation := "&"

	// default Currency
	if urlParams.Currency == nil {
		urlParams.Currency = entity.StringPtr("usd")
	} else {
		urlExcitation = urlExcitation + "currency=" + *urlParams.Currency + "&"
	}

	// default Page
	if urlParams.Page == nil {
		urlParams.Page = pointer.Uint64Ptr(1)
	} else {
		urlExcitation = urlExcitation + "page=" + strconv.FormatUint(uint64(*urlParams.Page), 10) + "&"
	}

	// default PerPage
	if urlParams.PerPage == nil {
		urlParams.PerPage = pointer.Uint64Ptr(18)
	} else {
		urlExcitation = urlExcitation + "perPage=" + strconv.FormatUint(uint64(*urlParams.Page), 10) + "&"
	}

	return urlParams, url + urlExcitation
}

func (c Controller) addUserToContext(
	ctx context.Context,
	r *http.Request,
) (context.Context, error) {
	ctx = context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
	return ctx, nil
}
