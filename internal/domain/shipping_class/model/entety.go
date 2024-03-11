package model

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

// fieldMap
var fieldMap = map[string]string{
	"ID":        "id",
	"Name":      "name",
	"Url":       "url",
	"SortOrder": "sort_order",
	"Active":    "active",
	"CreatedAt": "created_at",
	"CreatedBy": "created_by",
	"UpdatedAt": "updated_at",
	"UpdatedBy": "updated_by",
}
