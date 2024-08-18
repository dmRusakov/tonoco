package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"reflect"
)

// fieldMap
func (m *Model) fieldMap(field string) string {
	// check if field is in the cash
	if dbField, ok := m.dbFieldCash[field]; ok {
		return dbField
	}

	// get field from struct
	typeOf := reflect.TypeOf(Item{})
	byName, _ := typeOf.FieldByName(field)
	dbField := byName.Tag.Get("db")

	// set field to cash
	m.dbFieldCash[field] = dbField

	// done
	return dbField
}

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.fieldMap("Id"),
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
func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if filter.Ids != nil {
		statement = statement.Where(m.fieldMap("Id")+" = ?", (*filter.Ids)[0])
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
	if filter.Ids != nil {
		countIds := len(*filter.Ids)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("Id"): *filter.Ids})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countIds) {
			*filter.PerPage = uint64(countIds)
		}
	}

	// ProductIds
	if filter.ProductIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("ProductId"): *filter.ProductIds})
	}

	// TagTypeId
	if filter.TagTypeIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("TagTypeId"): *filter.TagTypeIds})
	}

	// TagSelectIds
	if filter.TagSelectIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("TagSelectId"): *filter.TagSelectIds})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.fieldMap("Value")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.fieldMap(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

// makeCountStatementByFilter - make count statement by filter for pagination
func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	// Build query
	statement := m.qb.Select("COUNT(*)").From(m.table + " p")

	// Ids
	if filter.Ids != nil {
		countIds := len(*filter.Ids)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("Id"): *filter.Ids})
		}
	}

	// ProductIds
	if filter.ProductIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("ProductId"): *filter.ProductIds})
	}

	// TagTypeId
	if filter.TagTypeIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("TagTypeId"): *filter.TagTypeIds})
	}

	// TagSelectIds
	if filter.TagSelectIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("TagSelectId"): *filter.TagSelectIds})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.fieldMap("Value")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var Id, ProductId, TagTypeId, TagSelectId, Value, CreatedBy, UpdatedBy sql.NullString
	var Active sql.NullBool
	var SortOrder sql.NullInt64
	var CreatedAt, UpdatedAt sql.NullTime
	err := rows.Scan(
		&Id,
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
		return nil, errors.AddCode(err, "304989")
	}

	var item = Item{}

	// Id if null, set it to empty string
	if Id.Valid {
		item.Id = Id.String
	}

	// ProductId if null, set it to empty string
	if ProductId.Valid {
		item.ProductId = ProductId.String
	}

	// TagTypeId if null, set it to empty string
	if TagTypeId.Valid {
		item.TagTypeId = TagTypeId.String
	}

	// TagSelectId if null, set it to empty string
	if TagSelectId.Valid {
		item.TagSelectId = TagSelectId.String
	}

	// Value if null, set it to empty string
	if Value.Valid {
		item.Value = Value.String
	}

	// Active if null, set it to false
	if Active.Valid {
		item.Active = Active.Bool
	}

	// SortOrder if null, set it to 0
	if SortOrder.Valid {
		item.SortOrder = uint64(SortOrder.Int64)
	}

	// CreatedAt if null, set it to time.Time{}
	if CreatedAt.Valid {
		item.CreatedAt = CreatedAt.Time
	}

	// CreatedBy if null, set it to empty string
	if CreatedBy.Valid {
		item.CreatedBy = CreatedBy.String
	}

	// UpdatedAt if null, set it to time.Time{}
	if UpdatedAt.Valid {
		item.UpdatedAt = UpdatedAt.Time
	}

	// UpdatedBy if null, set it to empty string
	if UpdatedBy.Valid {
		item.UpdatedBy = UpdatedBy.String
	}

	return &item, nil
}

// makeInsertStatement
func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *string) {

	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if Id is not set, generate a new UUID
	if item.Id == "" {
		item.Id = uuid.New().String()
	}

	// set Id to context
	ctx = context.WithValue(ctx, "itemId", item.Id)

	insertItem := m.qb.Insert(m.table).Columns(
		m.fieldMap("Id"),
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
		item.Id,
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

	return &insertItem, &item.Id
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
