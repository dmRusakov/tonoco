package entity

import "time"

type TagSelect struct {
	Id               string    `json:"id" db:"id"`
	TagTypeId        string    `json:"tag_type_id" db:"tag_type_id"`
	Name             string    `json:"name" db:"name"`
	Url              string    `json:"url" db:"url"`
	ShortDescription string    `json:"short_description" db:"short_description"`
	Description      string    `json:"description" db:"description"`
	Active           bool      `json:"active" db:"active"`
	SortOrder        uint64    `json:"sort_order" db:"sort_order"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	CreatedBy        string    `json:"created_by" db:"created_by"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy        string    `json:"updated_by" db:"updated_by"`
}

type TagSelectFilter struct {
	Ids        *[]string `json:"ids"`
	Urls       *[]string `json:"urls"`
	TagTypeIds *[]string `json:"tag_type_ids"`
	Active     *bool     `json:"active"`
	Search     *string   `json:"search"`
	OrderBy    *string   `json:"order_by"`
	OrderDir   *string   `json:"order_dir"`
	Page       *uint64   `json:"page"`
	PerPage    *uint64   `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
