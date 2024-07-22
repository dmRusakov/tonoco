package entity

import "time"

type Tag struct {
	Id          string    `json:"id" db:"id"`
	ProductId   string    `json:"product_id" db:"product_id"`
	TagTypeId   string    `json:"tag_type_id" db:"tag_type_id"`
	TagSelectId string    `json:"tag_select_id" db:"tag_select_id"`
	Value       string    `json:"value" db:"value"`
	Active      bool      `json:"active" db:"active"`
	SortOrder   uint64    `json:"sort_order" db:"sort_order"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy   string    `json:"updated_by" db:"updated_by"`
}

type TagFilter struct {
	Ids          *[]string `json:"ids"`
	ProductIds   *[]string `json:"product_ids"`
	TagTypeIds   *[]string `json:"tag_type_ids"`
	TagSelectIds *[]string `json:"tag_select_ids"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
