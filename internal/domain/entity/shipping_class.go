package entity

import "time"

type ShippingClass struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	SortOrder uint32    `json:"sort_order"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type ShippingClassFilter struct {
	Active    *bool   `json:"active"`
	Search    *string `json:"search"`
	SortBy    *string `json:"sort_by"`
	SortOrder *string `json:"sort_order"`
	Page      *uint64 `json:"page"`
	PerPage   *uint64 `json:"per_page"`
}
