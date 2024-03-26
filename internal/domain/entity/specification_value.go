package entity

type SpecificationValue struct {
	ID              string `json:"id"`
	SpecificationID string `json:"specification_id"`
	Name            string `json:"name"`
	Url             string `json:"url"`
	Active          bool   `json:"active"`
	Order           int32  `json:"order"`
	CreatedAt       int64  `json:"created_at"`
	CreatedBy       string `json:"created_by"`
	UpdatedAt       int64  `json:"updated_at"`
	UpdatedBy       string `json:"updated_by"`
}

type SpecificationValueFilter struct {
	Active          *bool   `json:"active"`
	SpecificationID *string `json:"unit"`
	Search          *string `json:"search"`
	OrderBy         *string `json:"order_by"`
	SortOrder       *string `json:"sort_order"`
	Page            *uint64 `json:"page"`
	PerPage         *uint64 `json:"per_page"`
}
