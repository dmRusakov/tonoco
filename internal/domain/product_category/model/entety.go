package model

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

// fieldMap
var fieldMap = map[string]string{
	"ID":               "id",
	"Name":             "name",
	"Url":              "url",
	"ShortDescription": "short_description",
	"Description":      "description",
	"SortOrder":        "sort_order",
	"Prime":            "prime",
	"Active":           "active",
	"CreatedAt":        "created_at",
	"CreatedBy":        "created_by",
	"UpdatedAt":        "updated_at",
	"UpdatedBy":        "updated_by",
}
