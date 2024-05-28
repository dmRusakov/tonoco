package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
)

type Item = entity.Tag
type Filter = entity.TagFilter

// fieldMap
var fieldMap = map[string]string{
	"ID":               "id",
	"TagTypeId":        "tag_type_id",
	"Name":             "name",
	"ShortDescription": "short_description",
	"Description":      "description",
	"Url":              "url",
	"Active":           "active",
	"SortOrder":        "sort_order",
	"CreatedAt":        "created_at",
	"CreatedBy":        "created_by",
	"UpdatedAt":        "updated_at",
	"UpdatedBy":        "updated_by",
}

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		fieldMap["ID"],
		fieldMap["TagTypeId"],
		fieldMap["Name"],
		fieldMap["ShortDescription"],
		fieldMap["Description"],
		fieldMap["Url"],
		fieldMap["Active"],
		fieldMap["SortOrder"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).From(m.table + " p")
}

// make Get statement
func (m *Model) makeGetStatement(id *string, url *string) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if id != nil {
		statement = statement.Where(fieldMap["ID"]+" = ?", *id)
	}

	// url
	if url != nil {
		statement = statement.Where(fieldMap["Url"]+" = ?", *url)
	}

	return statement
}

// makeStatementByFilter
func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = new(string)
		*filter.OrderBy = "SortOrder"
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = new(string)
		*filter.OrderDir = "ASC"
	}

	// Page
	if filter.Page == nil {
		filter.Page = new(uint64)
		*filter.Page = 1
	}

	// PerPage
	if filter.PerPage == nil {
		filter.PerPage = new(uint64)
		*filter.PerPage = 10
	}

	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.IDs != nil {
		countIds := len(*filter.IDs)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{fieldMap["ID"]: *filter.IDs})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countIds) {
			*filter.PerPage = uint64(countIds)
		}
	}

	// Urls
	if filter.Urls != nil {
		countUrls := len(*filter.Urls)

		if countUrls > 0 {
			statement = statement.Where(sq.Eq{fieldMap["Url"]: *filter.Urls})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countUrls) {
			*filter.PerPage = uint64(countUrls)
		}
	}

	// Type
	if filter.TagTypeId != nil {
		statement = statement.Where(sq.Eq{fieldMap["Type"]: *filter.TagTypeId})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{fieldMap["Active"]: *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+fieldMap["Name"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["Url"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["ShortDescription"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["Description"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(fieldMap[*filter.OrderBy] + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = &Item{}
	err := rows.Scan(
		&item.ID,
		&item.TagTypeId,
		&item.Name,
		&item.ShortDescription,
		&item.Description,
		&item.Url,
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

// makeInsertStatement
func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *string) {

	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if ID is not set, generate a new UUID
	if item.ID == "" {
		item.ID = uuid.New().String()
	}

	// set ID to context
	ctx = context.WithValue(ctx, "itemId", item.ID)

	insertItem := m.qb.Insert(m.table).Columns(
		fieldMap["ID"],
		fieldMap["TagTypeId"],
		fieldMap["Name"],
		fieldMap["ShortDescription"],
		fieldMap["Description"],
		fieldMap["Url"],
		fieldMap["Active"],
		fieldMap["SortOrder"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).Values(
		item.ID,
		item.TagTypeId,
		item.Name,
		item.ShortDescription,
		item.Description,
		item.Url,
		item.Active,
		item.SortOrder,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	return &insertItem, &item.ID
}

// makeUpdateStatement
func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(fieldMap["TagTypeId"], item.TagTypeId).
		Set(fieldMap["Name"], item.Name).
		Set(fieldMap["ShortDescription"], item.ShortDescription).
		Set(fieldMap["Description"], item.Description).
		Set(fieldMap["Url"], item.Url).
		Set(fieldMap["Active"], item.Active).
		Set(fieldMap["SortOrder"], item.SortOrder).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by)
}

// makePatchStatement
func (m *Model) makePatchStatement(ctx context.Context, id *string, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		field = fieldMap[field]
		statement = statement.Set(field, value)
	}

	return statement.Set(fieldMap["UpdatedAt"], "NOW()").Set(fieldMap["UpdatedBy"], by)
}