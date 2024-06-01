package model

import (
	"context"
	"database/sql"
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
	"ID":          "id",
	"ProductId":   "product_id",
	"TagTypeId":   "tag_type_id",
	"TagSelectId": "tag_select_id",
	"Value":       "value",
	"Active":      "active",
	"SortOrder":   "sort_order",
	"CreatedAt":   "created_at",
	"CreatedBy":   "created_by",
	"UpdatedAt":   "updated_at",
	"UpdatedBy":   "updated_by",
}

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		fieldMap["ID"],
		fieldMap["ProductId"],
		fieldMap["TagTypeId"],
		fieldMap["TagSelectId"],
		fieldMap["Value"],
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

	// PerPage
	if filter.PerPage == nil {
		filter.PerPage = new(uint64)
		if filter.Page == nil {
			*filter.PerPage = 999999999999999999
		} else {
			*filter.PerPage = 10
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = new(uint64)
		*filter.Page = 1
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

	// ProductIDs
	if filter.ProductIDs != nil {
		statement = statement.Where(sq.Eq{fieldMap["ProductId"]: *filter.ProductIDs})
	}

	// TagTypeId
	if filter.TagTypeIDs != nil {
		statement = statement.Where(sq.Eq{fieldMap["TagTypeId"]: *filter.TagTypeIDs})
	}

	// TagSelectIDs
	if filter.TagSelectIDs != nil {
		statement = statement.Where(sq.Eq{fieldMap["TagSelectId"]: *filter.TagSelectIDs})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{fieldMap["Active"]: *filter.Active})
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(fieldMap[*filter.OrderBy] + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = Item{}
	var productId, tagTypeId, tagSelectId, value sql.NullString

	err := rows.Scan(
		&item.ID,
		&productId,
		&tagTypeId,
		&tagSelectId,
		&value,
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

	// if productId if null, set it to empty string
	if productId.Valid {
		item.ProductId = productId.String
	} else {
		item.ProductId = ""
	}

	// if tagTypeId if null, set it to empty string
	if tagTypeId.Valid {
		item.TagTypeId = tagTypeId.String
	} else {
		item.TagTypeId = ""
	}

	// if tagSelectId if null, set it to empty string
	if tagSelectId.Valid {
		item.TagSelectId = tagSelectId.String
	} else {
		item.TagSelectId = ""
	}

	// if value if null, set it to empty string
	if value.Valid {
		item.Value = value.String
	} else {
		item.Value = ""
	}

	return &item, nil
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
		fieldMap["ProductId"],
		fieldMap["TagTypeId"],
		fieldMap["TagSelectId"],
		fieldMap["Value"],
		fieldMap["Active"],
		fieldMap["SortOrder"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).Values(
		item.ID,
		item.ProductId,
		item.TagTypeId,
		item.TagSelectId,
		item.Value,
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
		Set(fieldMap["ProductId"], item.ProductId).
		Set(fieldMap["TagTypeId"], item.TagTypeId).
		Set(fieldMap["TagSelectId"], item.TagSelectId).
		Set(fieldMap["Value"], item.Value).
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
