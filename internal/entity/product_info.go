package entity

import (
	"github.com/google/uuid"
	"time"
)

type ProductInfo struct {
	Id                    uuid.UUID `json:"id" db:"id"`
	Sku                   string    `json:"sku" db:"sku"`
	Brand                 string    `json:"brand" db:"brand"`
	Name                  string    `json:"name" db:"name"`
	ShortDescription      string    `json:"short_description" db:"short_description"`
	Description           string    `json:"description" db:"description"`
	SortOrder             uint64    `json:"sort_order" db:"sort_order"`
	Url                   string    `json:"url" db:"url"`
	IsTaxable             bool      `json:"is_taxable" db:"is_taxable"`
	IsTrackStock          bool      `json:"is_track_stock" db:"is_track_stock"`
	ShippingWeight        uint64    `json:"shipping_weight" db:"shipping_weight"`
	ShippingWidth         uint64    `json:"shipping_width" db:"shipping_width"`
	ShippingHeight        uint64    `json:"shipping_height" db:"shipping_height"`
	ShippingLength        uint64    `json:"shipping_length" db:"shipping_length"`
	SeoTitle              string    `json:"seo_title" db:"seo_title"`
	SeoDescription        string    `json:"seo_description" db:"seo_description"`
	GTIN                  string    `json:"gtin" db:"gtin"`
	GoogleProductCategory string    `json:"google_product_category" db:"google_product_category"`
	GoogleProductType     string    `json:"google_product_type" db:"google_product_type"`
	CreatedAt             time.Time `db:"created_at" json:"created_at"`
	CreatedBy             uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt             time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy             uuid.UUID `db:"updated_by" json:"updated_by"`
}

type ProductInfoFilter struct {
	Ids    *[]uuid.UUID `json:"ids"`
	Urls   *[]string    `json:"urls"`
	Skus   *[]string    `json:"skus"`
	Brands *[]string    `json:"brands"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
