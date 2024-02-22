package model

import "time"

// ProductCategory is a struct that contains the fields of the product_category table.
type ProductCategory struct {
	ID               string
	Name             string
	Slug             string
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

// fieldMap
var fieldMap = map[string]string{
	"ID":               "id",
	"Name":             "name",
	"Slug":             "slug",
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
