package entity

import "time"

// ProductCategory is a struct that contains the fields of the product_category table.
type ProductCategory struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Url              string    `json:"url"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	SortOrder        uint32    `json:"sort_order"`
	Prime            bool      `json:"prime"`
	Active           bool      `json:"active"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedAt        time.Time `json:"updated_at"`
	UpdatedBy        string    `json:"updated_by"`
}

type ProductCategoryFilter struct {
	Active   *bool   `json:"active"`
	Prime    *bool   `json:"prime"`
	Search   *string `json:"search"`
	OrderBy  *string `json:"sort_by"`
	OrderDir *string `json:"sort_order"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`
}
