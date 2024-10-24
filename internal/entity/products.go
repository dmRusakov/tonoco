package entity

import (
	"github.com/google/uuid"
)

type ProductListItem struct {
	Id               uuid.UUID                     `json:"id" db:"id"`
	No               int32                         `json:"no" db:"no"`
	Sku              string                        `json:"sku" db:"sku"`
	Brand            string                        `json:"brand" db:"brand"`
	Name             string                        `json:"name" db:"name"`
	ShortDescription string                        `json:"short_description" db:"short_description"`
	Url              string                        `json:"url" db:"url"`
	SalePrice        string                        `json:"sale_price" db:"price"`
	Price            string                        `json:"price" db:"price"`
	Currency         string                        `json:"currency" db:"currency"`
	Quantity         int32                         `json:"quantity" db:"quantity"`
	Status           string                        `json:"status" db:"status"`
	IsTaxable        bool                          `json:"is_taxable" db:"is_taxable"`
	SeoTitle         string                        `json:"seo_title" db:"seo_title"`
	SeoDescription   string                        `json:"seo_description" db:"seo_description"`
	Categories       []string                      `json:"categories" db:"categories"`
	Tags             map[uint32]ProductListItemTag `json:"tags" db:"tags"`
	MainImage        *Image                        `json:"main_image" db:"main_image"`
	HoverImage       *Image                        `json:"hover_image" db:"hover_image"`
}

type ProductsPageUrlParams struct {
	Category *string
	Currency *string
	Page     *uint64
	PerPage  *uint64
	Count    *uint64
}

type ProductsPageParams struct {
	Currency            Currency
	RegularPriceTypeIds []uuid.UUID
	SpecialPriceTypeIds []uuid.UUID
	Page                *uint64
	PerPage             *uint64
	Count               *uint64
}

type ProductListItemTag struct {
	Name  string `json:"name" db:"name"`
	Url   string `json:"url" db:"url"`
	Value string `json:"value" db:"value"`
}
