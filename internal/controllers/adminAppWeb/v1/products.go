package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/dmRusakov/tonoco/pkg/common/pagination"
	"github.com/dmRusakov/tonoco/pkg/utils/html"
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

	// read url params
	wg.Add(1)
	var url *pages.ProductsPageUrl
	go func() {
		defer wg.Done()
		url = c.readProductUrlParam(r)
	}()

	// add user to context
	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx = c.addUserToContext(ctx)
	}()

	wg.Wait()

	/* second round loading */

	// get shop page
	wg.Add(1)
	shopPage := &pages.ShopPage{}
	go func() {
		defer wg.Done()
		shop, e := c.shopUseCase.GetShopPage(ctx, page)
		if e != nil {
			errs = append(errs, e...)
		}

		shopPage.Id = shop.Id
		shopPage.Name = html.GetTemplate(shop.Name)
		shopPage.SeoTitle = html.GetTemplate(shop.SeoTitle)
		shopPage.Url = shop.Url
		shopPage.ShortDescription = html.GetTemplate(shop.ShortDescription)
		shopPage.Description = html.GetTemplate(shop.Description)
		shopPage.ConsoleMessage = pages.ConsoleMessage{}

		// page
		if url.Params.Page == nil {
			shopPage.Page = shop.Page
			url.Params.Page = &shop.Page
		} else {
			shopPage.Page = pointer.PtrToUint64(url.Params.Page)
		}

		// PerPage
		if url.Params.PerPage == nil {
			shopPage.PerPage = shop.PerPage
			url.Params.PerPage = &shop.PerPage
		} else {
			shopPage.PerPage = pointer.PtrToUint64(url.Params.PerPage)
		}

		// url
		shopPage.ShopPageUrl = c.cfg.ShopPageUrl
	}()

	wg.Wait()

	/* third round loading */

	// get products
	var products *map[uuid.UUID]*pages.ProductGridItem
	wg.Add(1)
	var ids *[]uuid.UUID
	go func() {
		defer wg.Done()
		var err error
		products, ids, err = c.shopUseCase.GetProductList(ctx, &url.Params, &shopPage.Id)
		if err != nil {
			errs = append(errs, err)
		}
		shopPage.TotalItems = pointer.PtrToUint64(url.Params.Count)
	}()

	// get filters tag for shop page
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		shopPage.Filter, err = c.shopUseCase.GetShopPageFilter(ctx, &shopPage.Id)
		if err != nil {
			errs = append(errs, err)
		}
	}()

	wg.Wait()

	/* fourth round loading */

	// check for errors
	if len(errs) > 0 {
		http.Error(w, errs[0].Error(), http.StatusInternalServerError)
		return
	}

	// add a product
	wg.Add(1)
	shopPage.Items = make([]pages.ProductGridItem, 0)
	go func() {
		defer wg.Done()
		for _, id := range *ids {
			product := (*products)[id]
			shopPage.Items = append(shopPage.Items, pages.ProductGridItem{
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
	wg.Add(1)
	go func() {
		defer wg.Done()
		// total pages
		shopPage.TotalPages = ((shopPage.TotalItems) + (shopPage.PerPage) - 1) / shopPage.PerPage

		// pagination
		paginationPages := pagination.GetPagination((*shopPage).Page, (*shopPage).TotalPages, 5)
		shopPage.Pagination = make(map[uint64]pages.PaginationItem)
		for _, page := range paginationPages {
			newUrl := *url
			newUrl.Params.Page = &page
			shopPage.Pagination[page] = pages.PaginationItem{
				Page:        page,
				Url:         c.MakeProductPageUrl(newUrl),
				CurrentPage: *url.Params.Page,
			}
		}
	}()

	wg.Wait()
	/* fourth round loading */

	// make template
	var tmpl *template.Template
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		tmpl, err = c.makeTemplate("shop.gohtml")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}()

	wg.Wait()

	// render template
	if err := tmpl.Execute(w, shopPage); err != nil {
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
