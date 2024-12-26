package db

import (
	"github.com/google/uuid"
	"time"
)

type Price struct {
	Id          uuid.UUID `json:"id" db:"id"`
	ProductID   uuid.UUID `json:"product_id" db:"product_id"`
	PriceTypeID uuid.UUID `json:"price_type_id" db:"price_type_id"`
	CurrencyID  uuid.UUID `json:"currency_id" db:"currency_id"`
	WarehouseID uuid.UUID `json:"warehouse_id" db:"warehouse_id"`
	StoreID     uuid.UUID `json:"store_id" db:"store_id"`

	Price     float64 `json:"price" db:"price"`
	SortOrder uint64  `json:"sort_order" db:"sort_order"`
	Active    bool    `json:"active" db:"active"`

	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy uuid.UUID `db:"updated_by" json:"updated_by"`
}

type PriceFilter struct {
	Ids *[]uuid.UUID `json:"id"`

	ProductIds   *[]uuid.UUID `json:"product_id"`
	PriceTypeIds *[]uuid.UUID `json:"price_type_id"`
	CurrencyIds  *[]uuid.UUID `json:"currency_id"`
	WarehouseIds *[]uuid.UUID `json:"warehouse_id"`
	StoreIds     *[]uuid.UUID `json:"store_id"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	Count *uint64 `json:"count"`

	IsIdsOnly      *bool `json:"is_ids_only"`
	IsCount        *bool `json:"is_count"`
	IsUpdateFilter *bool `json:"is_update_filter"`
	IsKeepIdsOrder *bool `json:"is_keep_ids_order"`
}
