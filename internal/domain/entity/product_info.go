package entity

import "time"

type ProductInfo struct {
	ID                    string    `json:"id"`
	SKU                   string    `json:"sku"`
	Name                  string    `json:"name"`
	ShortDescription      string    `json:"short_description"`
	Description           string    `json:"description"`
	SortOrder             uint64    `json:"sort_order"`
	StatusID              string    `json:"status_id"`
	Url                   string    `json:"url"`
	RegularPrice          float32   `json:"regular_price"`
	SalePrice             float32   `json:"sale_price"`
	FactoryPrice          float32   `json:"factory_price"`
	IsTaxable             bool      `json:"is_taxable"`
	Quantity              int32     `json:"quantity"`
	ReturnToStockDate     int64     `json:"return_to_stock_date"`
	IsTrackStock          bool      `json:"is_track_stock"`
	ShippingClassID       string    `json:"shipping_class_id"`
	ShippingWeight        float32   `json:"shipping_weight"`
	ShippingWidth         float32   `json:"shipping_width"`
	ShippingHeight        float32   `json:"shipping_height"`
	ShippingLength        float32   `json:"shipping_length"`
	SeoTitle              string    `json:"seo_title"`
	SeoDescription        string    `json:"seo_description"`
	GTIN                  string    `json:"gtin"`
	GoogleProductCategory string    `json:"google_product_category"`
	GoogleProductType     string    `json:"google_product_type"`
	CreatedAt             time.Time `json:"created_at"`
	CreatedBy             string    `json:"created_by"`
	UpdatedAt             time.Time `json:"updated_at"`
	UpdatedBy             string    `json:"updated_by"`
}

type ProductInfoFilter struct {
	IDs      *[]string `json:"ids"`
	Urls     *[]string `json:"urls"`
	SKUs     *[]string `json:"skus"`
	StatusID *[]string `json:"status_ids"`
	Active   *bool     `json:"active"`
	Search   *string   `json:"search"`
	OrderBy  *string   `json:"order_by"`
	OrderDir *string   `json:"order_dir"`
	Page     *uint64   `json:"page"`
	PerPage  *uint64   `json:"per_page"`
}
