package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"time"
)

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

// make select
func (repo *Model) makeSelect() sq.SelectBuilder {
	return repo.qb.Select(
		fieldMap["ID"],
		fieldMap["Name"],
		fieldMap["Url"],
		fieldMap["SortOrder"],
		fieldMap["Active"],
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

	return repo.qb.Insert(repo.table).
		Columns(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["SortOrder"],
			fieldMap["Active"],
			fieldMap["CreatedAt"],
			fieldMap["CreatedBy"],
			fieldMap["UpdatedAt"],
			fieldMap["UpdatedBy"],
		).
		Values(
			item.ID,
			item.Name,
			item.Url,
			item.SortOrder,
			item.Active,
			"NOW()",
			by,
			"NOW()",
			by,
		)
}

// make update
func (repo *Model) makeUpdate(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return repo.qb.Update(repo.table).
		Set(fieldMap["Name"], item.Name).
		Set(fieldMap["Url"], item.Url).
		Set(fieldMap["SortOrder"], item.SortOrder).
		Set(fieldMap["Active"], item.Active).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by)
}

// make patch
func (repo *Model) makePatch(ctx context.Context, id string, fields map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := repo.qb.Update(repo.table).Where("id = ?", id)

	for field, value := range fields {
		field = fieldMap[field]
		statement = statement.Set(field, value)
	}

	return statement.Set(fieldMap["UpdatedAt"], "NOW()").Set(fieldMap["UpdatedBy"], by)
}

// make updated at
func (repo *Model) makeUpdatedAt(id string) sq.SelectBuilder {
	return repo.qb.Select(fieldMap["UpdatedAt"]).From(repo.table).Where("id = ?", id)
}

// make table updated
func (repo *Model) makeTableUpdated() sq.SelectBuilder {
	return repo.qb.Select("max(updated_at)").From(repo.table)
}

// scan get
func (repo *Model) scanGet(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = &Item{}
	err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.Url,
		&item.SortOrder,
		&item.Active,
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

// scan updated at
func (repo *Model) scanUpdatedAt(ctx context.Context, rows sq.RowScanner) (*time.Time, error) {
	var updatedAt time.Time
	err := rows.Scan(&updatedAt)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	return &updatedAt, nil
}
