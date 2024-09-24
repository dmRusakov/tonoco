package entity

import (
	"github.com/google/uuid"
	"time"
)

type StockQuantity struct {
	Id          uuid.UUID `json:"id" db:"id"`
	ProductId   uuid.UUID `json:"product_id" db:"product_id"`
	Quality     int32     `json:"quality" db:"quantity"`
	WarehouseId uuid.UUID `json:"warehouse_id" db:"warehouse_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy uuid.UUID `db:"updated_by" json:"updated_by"`
}

type StockQuantityFilter struct {
	Ids          *[]uuid.UUID `json:"ids"`
	ProductIds   *[]uuid.UUID `json:"product_ids"`
	WarehouseIds *[]uuid.UUID `json:"warehouse_id"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	Count *uint64 `json:"count"`

	IsIdsOnly          *bool `json:"is_ids_only"`
	IsCount            *bool `json:"is_count"`
	IsUpdateFilter     *bool `json:"is_update_filter"`
	IsRemoveDuplicates *bool `json:"is_remove_duplicates"`
}
