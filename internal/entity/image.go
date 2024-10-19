package entity

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	Id           uuid.UUID `db:"id"`
	FileName     string    `db:"filename"`
	Extension    string    `db:"extension"`
	IsCompressed bool      `db:"is_compressed"`
	IsWebp       bool      `db:"is_webp"`
	FolderId     uuid.UUID `db:"folder_id"`
	SortOrder    uint64    `db:"sort_order"`
	Title        string    `db:"title"`
	AltText      string    `db:"alt_text"`
	CopyRight    string    `db:"copyright"`
	Creator      string    `db:"creator"`
	Rating       float32   `db:"rating"`
	OriginPath   string    `db:"origin_path"`
	CreatedAt    time.Time `db:"created_at"`
	CreatedBy    uuid.UUID `db:"created_by"`
	UpdatedAt    time.Time `db:"updated_at"`
	UpdatedBy    uuid.UUID `db:"updated_by"`
}

type ImageFilter struct {
	Ids *[]uuid.UUID `json:"id"`

	IsWebp *bool   `json:"is_webp"`
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

type ImageCompression struct {
	Ids         *[]uuid.UUID `json:"id"`
	FileName    *string      `json:"file_name"`
	Title       *string      `json:"title"`
	AltText     *string      `json:"alt_text"`
	CopyRight   *string      `json:"copy_right"`
	Creator     *string      `json:"creator"`
	Rating      *float32     `json:"rating"`
	Compression *uint        `json:"compression"`
}
