package pages

import (
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/google/uuid"
	"html/template"
)

type ShopPage struct {
	Id               uuid.UUID
	Name             template.HTML
	SeoTitle         template.HTML
	ShortDescription template.HTML
	Description      template.HTML
	Items            []ProductGridItem
	Filter           *ShopPageFilter
	Url              string
	Page             uint64
	PerPage          uint64
	TotalPages       uint64
	TotalItems       uint64
	Pagination       map[uint64]PaginationItem
	ConsoleMessage   ConsoleMessage

	// url
	ShopPageUrl string
}

type ProductGridItem struct {
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
	Quantity         int64                         `json:"quantity" db:"quantity"`
	Status           string                        `json:"status" db:"status"`
	IsTaxable        bool                          `json:"is_taxable" db:"is_taxable"`
	SeoTitle         string                        `json:"seo_title" db:"seo_title"`
	SeoDescription   string                        `json:"seo_description" db:"seo_description"`
	Categories       []string                      `json:"categories" db:"categories"`
	Tags             map[uint64]ProductListItemTag `json:"tags" db:"tags"`
	MainImage        *db.Image                     `json:"main_image" db:"main_image"`
	HoverImage       *db.Image                     `json:"hover_image" db:"hover_image"`
}

type ProductsPageUrl struct {
	Params ProductsPageUrlParams
	Url    string
}

type ProductsPageUrlParams struct {
	Category *string
	Currency *string
	Page     *uint64
	PerPage  *uint64
	Count    *uint64
}

type ProductsPageParams struct {
	Currency            db.Currency
	RegularPriceTypeIds []uuid.UUID
	SpecialPriceTypeIds []uuid.UUID
	Page                *uint64
	PerPage             *uint64
	Count               *uint64
}

type ProductListItemTag struct {
	Name  string
	Url   string
	Value string
}

type ProductGridTagTypes struct {
	TagTypes    *map[uuid.UUID]db.TagType
	TagOrder    *map[uuid.UUID]uint64
	TagTypesIds *[]uuid.UUID
}
