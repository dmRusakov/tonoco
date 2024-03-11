package entity

type ShippingClass struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	SortOrder uint32 `json:"sort_order"`
	Active    bool   `json:"active"`
	CreatedAt uint32 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt uint32 `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

type ShippingClassFilter struct {
	Active    *bool
	Search    *string
	SortBy    *string
	SortOrder *string
	Page      *uint64
	PerPage   *uint64
}
