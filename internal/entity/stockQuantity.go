package entity

import "time"

type StockQuantity struct {
	Id          string `json:"id" db:"id"`
	ProductId   string `json:"product_id" db:"product_id"`
	Quality     int32  `json:"quality" db:"quantity"`
	WarehouseId string `json:"warehouse_id" db:"warehouse_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy string    `json:"updated_by" db:"updated_by"`
}

type StockQuantityFilter struct {
	Ids          *[]string `json:"ids"`
	ProductIds   *[]string `json:"product_ids"`
	WarehouseIds *[]string `json:"warehouse_id"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
