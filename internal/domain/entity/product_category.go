package entity

import "time"

// ProductCategory is a struct that contains the fields of the product_category table.
type ProductCategory struct {
	ID               string
	Name             string
	Url              string
	ShortDescription string
	Description      string
	SortOrder        uint32
	Prime            bool
	Active           bool
	CreatedAt        time.Time
	CreatedBy        string
	UpdatedAt        time.Time
	UpdatedBy        string
}

type ProductCategoryFilter struct {
	Active    *bool
	Prime     *bool
	Search    *string
	SortBy    *string
	SortOrder *string
	Page      *uint64
	PerPage   *uint64
}
