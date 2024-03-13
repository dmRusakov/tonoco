package model

import "github.com/dmRusakov/tonoco/internal/domain/entity"

type Item = entity.ShippingClass
type Filter = entity.ShippingClassFilter

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
