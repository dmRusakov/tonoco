package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Item = entity.Tag
type Filter = entity.TagFilter

func (m *Model) fieldMap(field string) string {
	typeOf := reflect.TypeOf(Item{})
	byName, _ := typeOf.FieldByName(field)
	return byName.Tag.Get("db")
}

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.fieldMap("ID"),
		m.fieldMap("ProductId"),
		m.fieldMap("TagTypeId"),
		m.fieldMap("TagSelectId"),
		m.fieldMap("Value"),
		m.fieldMap("Active"),
		m.fieldMap("SortOrder"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
	).From(m.table + " p")
}

// make Get statement
func (m *Model) makeGetStatement(id *string, url *string) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if id != nil {
		statement = statement.Where(m.fieldMap("ID")+" = ?", *id)
	}

	// url
	if url != nil {
		statement = statement.Where(m.fieldMap("Url")+" = ?", *url)
	}

	return statement
}

// makeStatementByFilter
func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = entity.StringPtr("SortOrder")
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = entity.StringPtr("ASC")
	}

	// PerPage
	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = entity.Uint64Ptr(999999999999999999)
		} else {
			filter.PerPage = entity.Uint64Ptr(10)
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = entity.Uint64Ptr(1)
	}

	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.IDs != nil {
		countIds := len(*filter.IDs)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("ID"): *filter.IDs})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countIds) {
			*filter.PerPage = uint64(countIds)
		}
	}

	// ProductIDs
	if filter.ProductIDs != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("ProductId"): *filter.ProductIDs})
	}

	// TagTypeId
	if filter.TagTypeIDs != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("TagTypeId"): *filter.TagTypeIDs})
	}

	// TagSelectIDs
	if filter.TagSelectIDs != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("TagSelectId"): *filter.TagSelectIDs})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Active"): *filter.Active})
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.fieldMap(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var ID, ProductId, TagTypeId, TagSelectId, Value, CreatedBy, UpdatedBy sql.NullString
	var Active sql.NullBool
	var SortOrder sql.NullInt64
	var CreatedAt, UpdatedAt sql.NullTime
	err := rows.Scan(
		&ID,
		&ProductId,
		&TagTypeId,
		&TagSelectId,
		&Value,
		&Active,
		&SortOrder,
		&CreatedAt,
		&CreatedBy,
		&UpdatedAt,
		&UpdatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	var item = Item{}

	// ID if null, set it to empty string
	if ID.Valid {
		item.ID = ID.String
	} else {
		item.ID = ""
	}

	// ProductId if null, set it to empty string
	if ProductId.Valid {
		item.ProductId = ProductId.String
	} else {
		item.ProductId = ""
	}

	// TagTypeId if null, set it to empty string
	if TagTypeId.Valid {
		item.TagTypeId = TagTypeId.String
	} else {
		item.TagTypeId = ""
	}

	// TagSelectId if null, set it to empty string
	if TagSelectId.Valid {
		item.TagSelectId = TagSelectId.String
	} else {
		item.TagSelectId = ""
	}

	// Value if null, set it to empty string
	if Value.Valid {
		item.Value = Value.String
	} else {
		item.Value = ""
	}

	// Active if null, set it to false
	if Active.Valid {
		item.Active = Active.Bool
	} else {
		item.Active = false
	}

	// SortOrder if null, set it to 0
	if SortOrder.Valid {
		item.SortOrder = uint64(SortOrder.Int64)
	} else {
		item.SortOrder = 0
	}

	// CreatedAt if null, set it to time.Time{}
	if CreatedAt.Valid {
		item.CreatedAt = CreatedAt.Time
	} else {
		item.CreatedAt = time.Time{}
	}

	// CreatedBy if null, set it to empty string
	if CreatedBy.Valid {
		item.CreatedBy = CreatedBy.String
	} else {
		item.CreatedBy = ""
	}

	// UpdatedAt if null, set it to time.Time{}
	if UpdatedAt.Valid {
		item.UpdatedAt = UpdatedAt.Time
	} else {
		item.UpdatedAt = time.Time{}
	}

	// UpdatedBy if null, set it to empty string
	if UpdatedBy.Valid {
		item.UpdatedBy = UpdatedBy.String
	} else {
		item.UpdatedBy = ""
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
		m.fieldMap("ID"),
		m.fieldMap("ProductId"),
		m.fieldMap("TagTypeId"),
		m.fieldMap("TagSelectId"),
		m.fieldMap("Value"),
		m.fieldMap("Active"),
		m.fieldMap("SortOrder"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
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
		Set(m.fieldMap("ProductId"), item.ProductId).
		Set(m.fieldMap("TagTypeId"), item.TagTypeId).
		Set(m.fieldMap("TagSelectId"), item.TagSelectId).
		Set(m.fieldMap("Value"), item.Value).
		Set(m.fieldMap("Active"), item.Active).
		Set(m.fieldMap("SortOrder"), item.SortOrder).
		Set(m.fieldMap("UpdatedAt"), "NOW()").
		Set(m.fieldMap("UpdatedBy"), by)
}

// makePatchStatement
func (m *Model) makePatchStatement(ctx context.Context, id *string, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		field = m.fieldMap(field)
		statement = statement.Set(field, value)
	}

	return statement.Set(m.fieldMap("UpdatedAt"), "NOW()").Set(m.fieldMap("UpdatedBy"), by)
}
