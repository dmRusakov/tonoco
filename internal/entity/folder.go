package entity

import (
	"github.com/google/uuid"
	"time"
)

type Folder struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	ParentID  string    `json:"parent_id"`
	SortOrder uint64    `json:"sort_order"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy uuid.UUID `db:"updated_by" json:"updated_by"`
}

type FolderFilter struct {
	Ids  *[]uuid.UUID `json:"ids"`
	Urls *[]string    `json:"urls"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

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
