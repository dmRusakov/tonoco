package shop_page

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
)

func (u *UseCase) GetGridItem(ctx context.Context) *pages.ProductGridItem {
	item := pages.ProductGridItem{
		No:               0,
		Sku:              "",
		Brand:            "",
		Name:             "",
		ShortDescription: "",
		Url:              "",
		SalePrice:        "",
		Price:            "",
		Currency:         "",
		Quantity:         0,
		Status:           "",
		IsTaxable:        false,
		SeoTitle:         "",
		SeoDescription:   "",
		Categories:       nil,
		Tags:             nil,
		MainImage:        nil,
		HoverImage:       nil,
	}

	return &item
}
