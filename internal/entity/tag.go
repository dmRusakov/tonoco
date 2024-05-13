package entity

import "time"

type Tag struct {
	ID               string    `json:"id"`
	TagTypeId        string    `json:"tag_type_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	Url              string    `json:"url"`
	Active           bool      `json:"active"`
	SortOrder        uint64    `json:"sort_order"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedAt        time.Time `json:"updated_at"`
	UpdatedBy        string    `json:"updated_by"`
}

type TagFilter struct {
	IDs       *[]string `json:"ids"`
	URLs      *[]string `json:"urls"`
	TagTypeId *[]string `json:"tag_type_id"`
	Active    *bool     `json:"active"`
	Search    *string   `json:"search"`
	OrderBy   *string   `json:"order_by"`
	OrderDir  *string   `json:"order_dir"`
	Page      *uint64   `json:"page"`
	PerPage   *uint64   `json:"per_page"`
}
