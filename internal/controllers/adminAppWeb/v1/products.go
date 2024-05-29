package admin_app_web_v1

import (
	"context"
	"fmt"
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
	var params *entity.ProductsUrlParameters

	wg.Add(2)

	go func() {
		defer wg.Done()
		tmpl = s.makeTemplate("products.page.gohtml")
	}()

	go func() {
		defer wg.Done()
		params = s.ReadProductsUrlPara(r)
	}()

	wg.Wait()

	// get products
	products, err := s.productUseCase.GetProductList(ctx, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get tag_types with `list_item` type

	fmt.Println(products[2], "products:37")

	// Use tmpl and params here
	// render page
	if err := tmpl.Execute(w, params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

// ReadProductsUrlPara read page parameters from url
func (s server) ReadProductsUrlPara(r *http.Request) *entity.ProductsUrlParameters {
	params := &entity.ProductsUrlParameters{}
	v := reflect.ValueOf(params).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		value := r.URL.Query().Get(strings.ToLower(fieldName))
		if value != "" {
			v.Field(i).SetString(value)
		}
	}

	return params
}
