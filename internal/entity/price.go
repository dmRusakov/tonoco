package entity

type Price struct {
	Id          string `json:"id" db:"id"`
	ProductID   string `json:"product_id" db:"product_id"`
	PriceTypeID string `json:"price_type_id" db:"price_type_id"`
	CurrencyID  string `json:"currency_id" db:"currency_id"`
	WarehouseID string `json:"warehouse_id" db:"warehouse_id"`
	StoreId     string `json:"store_id" db:"store_id"`

	Price  uint64 `json:"price" db:"price"`
	Active bool   `json:"active" db:"active"`

	CreatedAt string `json:"created_at" db:"created_at"`
	CreatedBy string `json:"created_by" db:"created_by"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	UpdatedBy string `json:"updated_by" db:"updated_by"`
}

type PriceFilter struct {
	Ids          *[]string `json:"id"`
	ProductIds   *[]string `json:"product_id"`
	PriceTypeIds *[]string `json:"price_type_id"`
	CurrencyIds  *[]string `json:"currency_id"`
	WarehouseIds *[]string `json:"warehouse_id"`
	StoreIds     *[]string `json:"store_id"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`
}
