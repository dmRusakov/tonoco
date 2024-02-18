package model

import "time"

// Product is a struct that maps to the 'public.Product' table in the PostgreSQL database.
type Product struct {
	ID                    string
	SKU                   string
	Name                  string
	ShortDescription      string
	Description           string
	SortOrder             uint32
	StatusID              string
	Slug                  string
	RegularPrice          float32
	SalePrice             float32
	FactoryPrice          float32
	IsTaxable             bool
	Quantity              int32
	ReturnToStockDate     time.Time
	IsTrackStock          bool
	ShippingClassID       string
	ShippingWeight        float32
	ShippingWidth         float32
	ShippingHeight        float32
	ShippingLength        float32
	SeoTitle              string
	SeoDescription        string
	GTIN                  string
	GoogleProductCategory string
	GoogleProductType     string
	CreatedAt             time.Time
	CreatedBy             string
	UpdatedAt             time.Time
	UpdatedBy             string
}

// fieldMap
var fieldMap = map[string]string{
	"ID":                    "id",
	"SKU":                   "sku",
	"Name":                  "name",
	"ShortDescription":      "short_description",
	"Description":           "description",
	"SortOrder":             "sort_order",
	"StatusID":              "status_id",
	"Slug":                  "slug",
	"RegularPrice":          "regular_price",
	"SalePrice":             "sale_price",
	"FactoryPrice":          "factory_price",
	"IsTaxable":             "is_taxable",
	"Quantity":              "quantity",
	"ReturnToStockDate":     "return_to_stock_date",
	"IsTrackStock":          "is_track_stock",
	"ShippingClassID":       "shipping_class_id",
	"ShippingWeight":        "shipping_weight",
	"ShippingWidth":         "shipping_width",
	"ShippingHeight":        "shipping_height",
	"ShippingLength":        "shipping_length",
	"SeoTitle":              "seo_title",
	"SeoDescription":        "seo_description",
	"GTIN":                  "gtin",
	"GoogleProductCategory": "google_product_category",
	"GoogleProductType":     "google_product_type",
	"CreatedAt":             "created_at",
	"CreatedBy":             "created_by",
	"UpdatedAt":             "updated_at",
	"UpdatedBy":             "updated_by",
}
