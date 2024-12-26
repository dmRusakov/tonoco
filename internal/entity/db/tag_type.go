package db

import (
	"github.com/google/uuid"
	"time"
)

type TagType struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Url       string    `json:"url" db:"url"`
	Active    bool      `json:"active" db:"active"`
	SortOrder uint64    `json:"sort_order" db:"sort_order"`
	Type      string    `json:"type" db:"type"`
	Prefix    string    `json:"prefix" db:"prefix"`
	Suffix    string    `json:"suffix" db:"suffix"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy uuid.UUID `db:"updated_by" json:"updated_by"`
}

type TagTypeFilter struct {
	Ids  *[]uuid.UUID
	Urls *[]string

	Active *bool
	Type   *string

	Search *string

	OrderBy  *string
	OrderDir *string
	Page     *uint64
	PerPage  *uint64

	Count *uint64 `json:"count"`

	IsIdsOnly      *bool `json:"is_ids_only"`
	IsCount        *bool `json:"is_count"`
	IsUpdateFilter *bool `json:"is_update_filter"`
	IsKeepIdsOrder *bool `json:"is_keep_ids_order"`
}
