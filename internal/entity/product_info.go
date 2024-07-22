package entity

import "time"

type ProductInfo struct {
	Id                    string    `json:"id" db:"id"`
	Sku                   string    `json:"sku" db:"sku"`
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
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	CreatedBy             string    `json:"created_by" db:"created_by"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy             string    `json:"updated_by" db:"updated_by"`
}

type ProductInfoFilter struct {
	Ids  *[]string `json:"ids"`
	Urls *[]string `json:"urls"`
	Skus *[]string `json:"skus"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
