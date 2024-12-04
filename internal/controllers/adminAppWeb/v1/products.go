package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/dmRusakov/tonoco/pkg/common/pagination"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/dmRusakov/tonoco/pkg/utils/standart"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// URL params for products page

func (c Controller) RenderShopPage(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
	page string,
) {
	var wg sync.WaitGroup
	var errs []error

	/* first round loading */

	wg.Add(3)

	// make template
	var tmpl *template.Template
	go func() {
		defer wg.Done()
		tmpl = c.makeTemplate("products.page.gohtml")
	}()

	// read url params
	var url *pages.ProductsPageUrl
	go func() {
		defer wg.Done()
		url = c.readProductUrlParam(r)
	}()

	// add user to context
	go func() {
		defer wg.Done()
		ctx = c.addUserToContext(ctx)
	}()

	wg.Wait()

	/* second round loading */

	// get shop page
	wg.Add(1)
	productPage := &pages.ProductPage{}
	go func() {
		defer wg.Done()
		shopPage, e := c.shopPageUseCase.GetShopPage(ctx, page)
		if e != nil {
			errs = append(errs, e...)
		}

		productPage.Name = shopPage.Name
		productPage.SeoTitle = shopPage.SeoTitle
		productPage.Url = shopPage.Url
		productPage.ConsoleMessage = pages.ConsoleMessage{}

		// page
		if url.Params.Page == nil {
			productPage.Page = shopPage.Page
			url.Params.Page = &shopPage.Page
		} else {
			productPage.Page = pointer.PtrToUint64(url.Params.Page)
		}

		// PerPage
		if url.Params.PerPage == nil {
			productPage.PerPage = shopPage.PerPage
			url.Params.PerPage = &shopPage.PerPage
		} else {
			productPage.PerPage = pointer.PtrToUint64(url.Params.PerPage)
		}

		// url
		productPage.ShopPageUrl = c.cfg.ShopPageUrl
	}()

	wg.Wait()

	// get products
	var products *map[uuid.UUID]*pages.ProductGridItem
	wg.Add(1)
	var ids *[]uuid.UUID
	go func() {
		defer wg.Done()
		var err error
		products, ids, err = c.shopPageUseCase.GetProductList(ctx, &url.Params)
		if err != nil {
			errs = append(errs, err)
		}
		productPage.TotalItems = pointer.PtrToUint64(url.Params.Count)
	}()

	wg.Wait()

	// check for errors
	if len(errs) > 0 {
		http.Error(w, errs[0].Error(), http.StatusInternalServerError)
		return
	}

	/* third round loading */
	wg.Add(2)

	// add a product
	productPage.Items = make([]pages.ProductGridItem, 0)
	go func() {
		defer wg.Done()
		for _, id := range *ids {
			product := (*products)[id]
			productPage.Items = append(productPage.Items, pages.ProductGridItem{
				Id:               product.Id,
				No:               product.No,
				Sku:              product.Sku,
				Brand:            product.Brand,
				Name:             product.Name,
				ShortDescription: product.ShortDescription,
				Url:              product.Url,
				SalePrice:        product.SalePrice,
				Price:            product.Price,
				Currency:         product.Currency,
				Quantity:         product.Quantity,
				Status:           product.Status,
				IsTaxable:        product.IsTaxable,
				SeoTitle:         product.SeoTitle,
				SeoDescription:   product.SeoDescription,
				Categories:       product.Categories,
				Tags:             product.Tags,
				MainImage:        product.MainImage,
				HoverImage:       product.HoverImage,
			})
		}
	}()

	// total pages and pagination
	go func() {
		defer wg.Done()
		// total pages
		productPage.TotalPages = ((productPage.TotalItems) + (productPage.PerPage) - 1) / productPage.PerPage

		// pagination
		paginationPages := pagination.GetPagination((*productPage).Page, (*productPage).TotalPages, 5)
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
	}()

	wg.Wait()

	/* fourth round loading */

	// render template
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

func (c Controller) addUserToContext(ctx context.Context) context.Context {
	// TODO make it right
	return context.WithValue(ctx, "user_id", "0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1")
}
