package model

import "github.com/dmRusakov/tonoco/internal/domain/entity"

type Item = entity.Specification
type Filter = entity.SpecificationFilter

// fieldMap
var fieldMap = map[string]string{
	"ID":        "id",
	"Name":      "name",
	"Url":       "url",
	"Type":      "type",
	"Active":    "active",
	"SortOrder": "sort_order",
	"CreatedAt": "created_at",
	"CreatedBy": "created_by",
	"UpdatedAt": "updated_at",
	"UpdatedBy": "updated_by",
}
