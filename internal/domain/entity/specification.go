package entity

type Specification struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Type      string `json:"type"`
	Active    bool   `json:"active"`
	SortOrder int32  `json:"order"`
	CreatedAt int64  `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

type SpecificationFilter struct {
	Active    *bool
	Type      *string
	Search    *string
	SortBy    *string
	SortOrder *string
	Page      *uint64
	PerPage   *uint64
}
