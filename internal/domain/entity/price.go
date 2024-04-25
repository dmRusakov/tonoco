package entity

type Price struct {
	ID          string `json:"id"`
	ProductID   string `json:"product_id"`
	PriceTypeID string `json:"price_type_id"`
	CurrencyID  string `json:"currency_id"`
	WarehouseID string `json:"warehouse_id"`
	StoreId     string `json:"store_id"`

	Price uint64 `json:"price"`

	Active bool `json:"active"`

	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

type PriceFilter struct {
	IDs         *string `json:"id"`
	ProductID   *string `json:"product_id"`
	CurrencyID  *string `json:"currency_id"`
	WarehouseID *string `json:"warehouse_id"`
	StoreId     *string `json:"store_id"`

	Active *bool `json:"active"`
}
