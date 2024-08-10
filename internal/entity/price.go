package entity

import "time"

type Price struct {
	Id          string `json:"id" db:"id"`
	ProductID   string `json:"product_id" db:"product_id"`
	PriceTypeID string `json:"price_type_id" db:"price_type_id"`
	CurrencyID  string `json:"currency_id" db:"currency_id"`
	WarehouseID string `json:"warehouse_id" db:"warehouse_id"`
	StoreID     string `json:"store_id" db:"store_id"`

	Price     float64 `json:"price" db:"price"`
	SortOrder uint64  `json:"sort_order" db:"sort_order"`
	Active    bool    `json:"active" db:"active"`

	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy string    `json:"updated_by" db:"updated_by"`
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

	IsCount *bool `json:"is_count"`
}
