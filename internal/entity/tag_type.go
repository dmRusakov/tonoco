package entity

import (
	"github.com/google/uuid"
	"time"
)

type TagType struct {
	Id               uuid.UUID `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Url              string    `json:"url" db:"url"`
	ShortDescription string    `json:"short_description" db:"short_description"`
	Description      string    `json:"description" db:"description"`
	Required         bool      `json:"required" db:"required"`
	Active           bool      `json:"active" db:"active"`
	Prime            bool      `json:"prime" db:"prime"`
	ListItem         bool      `json:"list_item" db:"list_item"`
	Filter           bool      `json:"filter" db:"filter"`
	SortOrder        uint64    `json:"sort_order" db:"sort_order"`
	Type             string    `json:"type" db:"type"`
	Prefix           string    `json:"prefix" db:"prefix"`
	Suffix           string    `json:"suffix" db:"suffix"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	CreatedBy        uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy        uuid.UUID `db:"updated_by" json:"updated_by"`
}

type TagTypeFilter struct {
	Ids  *[]uuid.UUID
	Urls *[]string

	Required *bool
	Active   *bool
	Prime    *bool
	ListItem *bool
	Filter   *bool
	Type     *string

	Search *string

	OrderBy  *string
	OrderDir *string
	Page     *uint64
	PerPage  *uint64

	Count *uint64 `json:"count"`

	IsIdsOnly          *bool `json:"is_ids_only"`
	IsCount            *bool `json:"is_count"`
	IsUpdateFilter     *bool `json:"is_update_filter"`
	IsRemoveDuplicates *bool `json:"is_remove_duplicates"`
}
