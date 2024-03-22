package entity

import "time"

type Specification struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Url               string    `json:"url"`
	SpecificationType string    `json:"specification_type"`
	Active            bool      `json:"active"`
	SortOrder         uint32    `json:"order"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         string    `json:"updated_by"`
}

type SpecificationFilter struct {
	Active            *bool   `json:"active"`
	SpecificationType *string `json:"specification_type"`
	Search            *string `json:"search"`
	SortBy            *string `json:"sort_by"`
	SortOrder         *string `json:"sort_order"`
	Page              *uint64 `json:"page"`
	PerPage           *uint64 `json:"per_page"`
}
