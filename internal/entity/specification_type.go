package entity

import "time"

type SpecificationType struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Unit      string    `json:"unit"`
	SortOrder uint64    `json:"sort_order"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type SpecificationTypeFilter struct {
	IDs      *[]string `json:"ids"`
	Urls     *[]string `json:"urls"`
	Active   *bool     `json:"active"`
	Unit     *string   `json:"unit"`
	Search   *string   `json:"search"`
	OrderBy  *string   `json:"order_by"`
	OrderDir *string   `json:"order_dir"`
	Page     *uint64   `json:"page"`
	PerPage  *uint64   `json:"per_page"`
}
