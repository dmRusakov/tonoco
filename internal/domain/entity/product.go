package entity

type Product struct {
	ID                    string  `json:"id"`
	SKU                   string  `json:"sku"`
	Name                  string  `json:"name"`
	ShortDescription      string  `json:"short_description"`
	Description           string  `json:"description"`
	Order                 int32   `json:"order"`
	StatusID              string  `json:"status_id"`
	Slug                  string  `json:"slug"`
	RegularPrice          float32 `json:"regular_price"`
	SalePrice             float32 `json:"sale_price"`
	FactoryPrice          float32 `json:"factory_price"`
	IsTaxable             bool    `json:"is_taxable"`
	Quantity              int32   `json:"quantity"`
	ReturnToStockDate     int64   `json:"return_to_stock_date"`
	IsTrackStock          bool    `json:"is_track_stock"`
	ShippingClassID       string  `json:"shipping_class_id"`
	ShippingWeight        float32 `json:"shipping_weight"`
	ShippingWidth         float32 `json:"shipping_width"`
	ShippingHeight        float32 `json:"shipping_height"`
	ShippingLength        float32 `json:"shipping_length"`
	SeoTitle              string  `json:"seo_title"`
	SeoDescription        string  `json:"seo_description"`
	GTIN                  string  `json:"gtin"`
	GoogleProductCategory string  `json:"google_product_category"`
	GoogleProductType     string  `json:"google_product_type"`
	CreatedAt             int64   `json:"created_at"`
	CreatedBy             string  `json:"created_by"`
	UpdatedAt             int64   `json:"updated_at"`
	UpdatedBy             string  `json:"updated_by"`
}
