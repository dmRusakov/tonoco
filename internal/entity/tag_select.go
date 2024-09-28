package entity

import (
	"github.com/google/uuid"
	"time"
)

type TagSelect struct {
	Id               uuid.UUID `json:"id" db:"id"`
	TagTypeId        uuid.UUID `json:"tag_type_id" db:"tag_type_id"`
	Name             string    `json:"name" db:"name"`
	Url              string    `json:"url" db:"url"`
	ShortDescription string    `json:"short_description" db:"short_description"`
	Description      string    `json:"description" db:"description"`
	Active           bool      `json:"active" db:"active"`
	SortOrder        uint64    `json:"sort_order" db:"sort_order"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	CreatedBy        uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy        uuid.UUID `db:"updated_by" json:"updated_by"`
}

type TagSelectFilter struct {
	Ids        *[]uuid.UUID `json:"ids"`
	Urls       *[]string    `json:"urls"`
	TagTypeIds *[]uuid.UUID `json:"tag_type_ids"`
	Active     *bool        `json:"active"`
	Search     *string      `json:"search"`
	OrderBy    *string      `json:"order_by"`
	OrderDir   *string      `json:"order_dir"`
	Page       *uint64      `json:"page"`
	PerPage    *uint64      `json:"per_page"`

	Count *uint64 `json:"count"`

	IsIdsOnly      *bool `json:"is_ids_only"`
	IsCount        *bool `json:"is_count"`
	IsUpdateFilter *bool `json:"is_update_filter"`
	IsKeepIdsOrder *bool `json:"is_keep_ids_order"`
}
