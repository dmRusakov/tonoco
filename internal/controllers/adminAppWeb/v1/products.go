package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/dmRusakov/tonoco/pkg/common/pagination"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/dmRusakov/tonoco/pkg/utils/standart"
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
) {
	// add user to context TODO make it right
	ctx = context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")

	var wg sync.WaitGroup
	var tmpl *template.Template
	var url *pages.ProductsPageUrl

	wg.Add(2)

	go func() {
		defer wg.Done()
		tmpl = c.makeTemplate("products.page.gohtml")
	}()

	go func() {
		defer wg.Done()
		url = c.readProductUrlParam(r)
	}()

	wg.Wait()

	// get products
	consoleMessage := pages.ConsoleMessage{}

	// get products
	products, err := c.productUseCase.GetProductList(ctx, &url.Params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// page info
	productPage := pages.ProductPage{
		Name: "Range Hoods",

		Items: products,
		Url:   url.Url,

		Page:       pointer.PtrToUint64(url.Params.Page),
		PerPage:    pointer.PtrToUint64(url.Params.PerPage),
		TotalItems: pointer.PtrToUint64(url.Params.Count),

		ConsoleMessage: consoleMessage,
	}

	// total pages
	productPage.TotalPages = ((productPage.TotalItems) + (productPage.PerPage) - 1) / productPage.PerPage

	paginationPages := pagination.GetPagination(productPage.Page, productPage.TotalPages, 5)
	productPage.Pagination = make(map[uint64]pages.PaginationItem)
	for _, page := range paginationPages {
		newUrl := *url
		newUrl.Params.Page = &page
		productPage.Pagination[page] = pages.PaginationItem{
			Page:        page,
			Url:         c.MakeProductPageUrl(newUrl),
			CurrentPage: *url.Params.Page,
		}
	}

	// render page
	if err := tmpl.Execute(w, productPage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

// readProductUrlParam read page parameters from url
func (c Controller) readProductUrlParam(r *http.Request) *pages.ProductsPageUrl {
	// read url urlParams
	url := pages.ProductsPageUrl{}
	v := reflect.ValueOf(&url.Params).Elem()
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
	url.Url = strings.Split(r.URL.String(), "?")[0]

	// default Currency
	if url.Params.Currency == nil {
		url.Params.Currency = &c.cfg.StoreCurrency
	}

	// default Page
	if url.Params.Page == nil {
		url.Params.Page = &c.cfg.AppDefaultPage
	}

	// default PerPage
	if url.Params.PerPage == nil {
		url.Params.PerPage = &c.cfg.AppDefaultPerPAge
	}

	return &url
}

func (c Controller) MakeProductPageUrl(urlParams pages.ProductsPageUrl) string {
	url := urlParams.Url + "?"

	addParam := func(key, value string) {
		if value != "" {
			url += key + "=" + value + "&"
		}
	}

	addParam("currency", standart.GetStringValue(urlParams.Params.Currency, &c.cfg.StoreCurrency))
	addParam("page", standart.GetUint64Value(urlParams.Params.Page, &c.cfg.AppDefaultPage))
	addParam("perpage", standart.GetUint64Value(urlParams.Params.PerPage, &c.cfg.AppDefaultPerPAge))

	return strings.TrimRight(url, "&?")
}

func (c Controller) addUserToContext(
	ctx context.Context,
	r *http.Request,
) (context.Context, error) {
	ctx = context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
	return ctx, nil
}
