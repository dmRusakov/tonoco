package model

import "github.com/dmRusakov/tonoco/internal/domain/entity"

type Item = entity.ProductCategory
type Filter = entity.ProductCategoryFilter

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
