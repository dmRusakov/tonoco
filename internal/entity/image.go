package entity

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	ID            uuid.UUID `db:"id" json:"id"`
	Title         string    `db:"title" json:"title"`
	AltText       string    `db:"alt_text" json:"alt_text"`
	OriginPath    string    `db:"origin_path" json:"origin_path"`
	FullPath      string    `db:"full_path" json:"full_path"`
	LargePath     string    `db:"large_path" json:"large_path"`
	MediumPath    string    `db:"medium_path" json:"medium_path"`
	GridPath      string    `db:"grid_path" json:"grid_path"`
	ThumbnailPath string    `db:"thumbnail_path" json:"thumbnail_path"`
	SortOrder     uint64    `db:"sort_order" json:"sort_order"`
	IsWebp        bool      `db:"is_webp" json:"is_webp"`
	ImageType     string    `db:"image_type" json:"image_type"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	CreatedBy     uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy     uuid.UUID `db:"updated_by" json:"updated_by"`
}

type ImageFilter struct {
	Ids *[]uuid.UUID `json:"id"`

	IsWebp *bool   `json:"is_webp"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
