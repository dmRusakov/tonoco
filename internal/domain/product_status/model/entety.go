package model

import "time"

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

type Filter struct {
	Active    *bool
	Search    *string
	SortBy    *string
	SortOrder *string
	Page      *uint64
	PerPage   *uint64
}
