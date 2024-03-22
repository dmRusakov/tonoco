package entity

type SpecificationType struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Unit      string `json:"unit"`
	Active    bool   `json:"active"`
	Order     int32  `json:"order"`
	CreatedAt int64  `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

type SpecificationTypeFilter struct {
	Active    *bool   `json:"active"`
	Unit      *string `json:"unit"`
	Search    *string `json:"search"`
	SortBy    *string `json:"sort_by"`
	SortOrder *string `json:"sort_order"`
	Page      *uint64 `json:"page"`
	PerPage   *uint64 `json:"per_page"`
}
