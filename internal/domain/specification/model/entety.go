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
	"Type":      "type",
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
		"id",
		"name",
		"url",
		"'type' as type",
		"active",
		"sort_order",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
	).From(repo.table + " p")
}

// make insert
func (repo *Model) makeInsert(ctx context.Context, item *Item) (sq.InsertBuilder, error) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if ID is not set, generate a new UUID
	if item.ID == "" {
		item.ID = uuid.New().String()
	}

	// build query
	statement := repo.qb.Insert(repo.table).
		Columns(
			"id",
			"name",
			"url",
			"type",
			"active",
			"sort_order",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by").
		Values(
			item.ID,
			item.Name,
			item.Url,
			item.Type,
			item.Active,
			item.SortOrder,
			"NOW()",
			by,
			"NOW()",
			by,
		)

	return statement, nil
}

// make Scan
func (repo *Model) makeScan(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = &Item{}
	err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.Url,
		&item.Type,
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
