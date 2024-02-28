package model

import "time"

type ProductStatus struct {
	ID        string
	Name      string
	Slug      string
	SortOrder uint32
	Active    bool
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

// fieldMap
var fieldMap = map[string]string{
	"ID":        "id",
	"Name":      "name",
	"Slug":      "slug",
	"SortOrder": "sort_order",
	"Active":    "active",
	"CreatedAt": "created_at",
	"CreatedBy": "created_by",
	"UpdatedAt": "updated_at",
	"UpdatedBy": "updated_by",
}

// publicFields
var publicFields = []string{
	"ID",
	"Name",
	"Slug",
	"SortOrder",
}

// makeDbRequestColumns
func (repo *ProductStatusModel) makeDbRequestColumns() []string {
	columns := []string{}
	for _, field := range publicFields {
		columns = append(columns, fieldMap[field])
	}

	return columns
}
