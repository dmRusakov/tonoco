package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
)

type Item = entity.Specification
type Filter = entity.SpecificationFilter

// fieldMap
var fieldMap = map[string]string{
	"ID":        "id",
	"Name":      "name",
	"Url":       "url",
	"Type":      "specification_type",
	"Active":    "active",
	"SortOrder": "sort_order",
	"CreatedAt": "created_at",
	"CreatedBy": "created_by",
	"UpdatedAt": "updated_at",
	"UpdatedBy": "updated_by",
}

// make select
func (repo *Model) makeSelect() sq.SelectBuilder {
	return repo.qb.Select(
		fieldMap["ID"],
		fieldMap["Name"],
		fieldMap["Url"],
		fieldMap["Type"],
		fieldMap["Active"],
		fieldMap["SortOrder"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).From(repo.table + " p")
}

// make insert
func (repo *Model) makeInsert(ctx context.Context, item *Item) sq.InsertBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if ID is not set, generate a new UUID
	if item.ID == "" {
		item.ID = uuid.New().String()
	}

	// build query
	return repo.qb.Insert(repo.table).
		Columns(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["Type"],
			fieldMap["Active"],
			fieldMap["SortOrder"],
			fieldMap["CreatedAt"],
			fieldMap["CreatedBy"],
			fieldMap["UpdatedAt"],
			fieldMap["UpdatedBy"]).
		Values(
			item.ID,
			item.Name,
			item.Url,
			item.SpecificationType,
			item.Active,
			item.SortOrder,
			"NOW()",
			by,
			"NOW()",
			by,
		)
}

// make Scan
func (repo *Model) makeScan(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = &Item{}
	err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.Url,
		&item.SpecificationType,
		&item.Active,
		&item.SortOrder,
		&item.CreatedAt,
		&item.CreatedBy,
		&item.UpdatedAt,
		&item.UpdatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	return item, nil
}
