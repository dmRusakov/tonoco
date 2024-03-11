package entity

import "time"

type ProductStatus struct {
	ID        string
	Name      string
	Url       string
	SortOrder uint32
	Active    bool
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

type ProductStatusFilter struct {
	Active    *bool
	Search    *string
	SortBy    *string
	SortOrder *string
	Page      *uint64
	PerPage   *uint64
}
