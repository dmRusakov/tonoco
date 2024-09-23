package entity

import (
	"github.com/google/uuid"
	"time"
)

type ProductImage struct {
	Id        uuid.UUID `db:"id" json:"id"`
	ProductId uuid.UUID `db:"product_id" json:"product_id"`
	ImageId   uuid.UUID `db:"image_id" json:"image_id"`
	Type      string    `db:"type" json:"type"`
	SortOrder uint64    `db:"sort_order" json:"sort_order"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy uuid.UUID `db:"updated_by" json:"updated_by"`
}

type ProductImageFilter struct {
	Ids        *[]uuid.UUID `json:"id"`
	ProductIds *[]uuid.UUID `json:"product_id"`
	ImageIds   *[]uuid.UUID `json:"image_id"`

	Type   *[]string `json:"type"`
	Search *string   `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
