package entity

import "time"

type Specification struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Type      string    `json:"specification_type"`
	Active    bool      `json:"active"`
	SortOrder uint64    `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type SpecificationFilter struct {
	IDs      *[]string `json:"ids"`
	Urls     *[]string `json:"urls"`
	Active   *bool     `json:"active"`
	Type     *string   `json:"specification_type"`
	Search   *string   `json:"search"`
	OrderBy  *string   `json:"sort_by"`
	OrderDir *string   `json:"sort_order"`
	Page     *uint64   `json:"page"`
	PerPage  *uint64   `json:"per_page"`
}
