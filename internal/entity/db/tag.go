package db

import (
	"github.com/google/uuid"
	"time"
)

type Tag struct {
	Id          uuid.UUID `json:"id" db:"id"`
	ProductId   uuid.UUID `json:"product_id" db:"product_id"`
	TagTypeId   uuid.UUID `json:"tag_type_id" db:"tag_type_id"`
	TagSelectId uuid.UUID `json:"tag_select_id" db:"tag_select_id"`
	Value       string    `json:"value" db:"value"`
	Active      bool      `json:"active" db:"active"`
	SortOrder   uint64    `json:"sort_order" db:"sort_order"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	CreatedBy   uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy   uuid.UUID `db:"updated_by" json:"updated_by"`
}

type TagFilter struct {
	Ids          *[]uuid.UUID `json:"ids"`
	ProductIds   *[]uuid.UUID `json:"product_ids"`
	TagTypeIds   *[]uuid.UUID `json:"tag_type_ids"`
	TagSelectIds *[]uuid.UUID `json:"tag_select_ids"`

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
