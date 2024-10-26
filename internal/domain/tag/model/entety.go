package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/google/uuid"
	"reflect"
)

func (m *Model) mapFieldToDBColumn(field string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

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

func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("ProductId"),
		m.mapFieldToDBColumn("TagTypeId"),
		m.mapFieldToDBColumn("TagSelectId"),
		m.mapFieldToDBColumn("Value"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table + " p")
}

func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if filter.Ids != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Id")+" = ?", (*filter.Ids)[0])
	}

	return statement
}

func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = pointer.StringPtr("SortOrder")
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = pointer.StringPtr("ASC")
	}

	// PerPage
	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = pointer.Uint64Ptr(999999999999999999)
		} else {
			filter.PerPage = pointer.Uint64Ptr(10)
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = pointer.Uint64Ptr(1)
	}

	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.Ids != nil {
		countIds := len(*filter.Ids)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countIds) {
			*filter.PerPage = uint64(countIds)
		}
	}

	// ProductIds
	if filter.ProductIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("ProductId"): *filter.ProductIds})
	}

	// TagTypeId
	if filter.TagTypeIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("TagTypeId"): *filter.TagTypeIds})
	}

	// TagSelectIds
	if filter.TagSelectIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("TagSelectId"): *filter.TagSelectIds})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Value")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.mapFieldToDBColumn(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	// Build query
	statement := m.qb.Select("COUNT(*)").From(m.table + " p")

	// Ids
	if filter.Ids != nil {
		countIds := len(*filter.Ids)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
		}
	}

	// ProductIds
	if filter.ProductIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("ProductId"): *filter.ProductIds})
	}

	// TagTypeId
	if filter.TagTypeIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("TagTypeId"): *filter.TagTypeIds})
	}

	// TagSelectIds
	if filter.TagSelectIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("TagSelectId"): *filter.TagSelectIds})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Value")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, productId, tagTypeId, tagSelectId, value, createdBy, updatedBy sql.NullString
	var active sql.NullBool
	var sortOrder sql.NullInt64
	var createdAt, updatedAt sql.NullTime
	err := rows.Scan(
		&id,
		&productId,
		&tagTypeId,
		&tagSelectId,
		&value,
		&active,
		&sortOrder,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "304989")
	}

	var item = Item{}

	// id if null, set it to empty string
	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	// productId if null, set it to empty string
	if productId.Valid {
		item.ProductId = uuid.MustParse(productId.String)
	}

	// tagTypeId if null, set it to empty string
	if tagTypeId.Valid {
		item.TagTypeId = uuid.MustParse(tagTypeId.String)
	}

	// tagSelectId if null, set it to empty string
	if tagSelectId.Valid {
		item.TagSelectId = uuid.MustParse(tagSelectId.String)
	}

	// value if null, set it to empty string
	if value.Valid {
		item.Value = value.String
	}

	// active if null, set it to false
	if active.Valid {
		item.Active = active.Bool
	}

	// sortOrder if null, set it to 0
	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
	}

	// createdAt if null, set it to time.Time{}
	if createdAt.Valid {
		item.CreatedAt = createdAt.Time
	}

	// createdBy if null, set it to empty string
	if createdBy.Valid {
		item.CreatedBy = uuid.MustParse(createdBy.String)
	}

	// updatedAt if null, set it to time.Time{}
	if updatedAt.Valid {
		item.UpdatedAt = updatedAt.Time
	}

	// updatedBy if null, set it to empty string
	if updatedBy.Valid {
		item.UpdatedBy = uuid.MustParse(updatedBy.String)
	}

	return &item, nil
}

func (m *Model) scanCountRow(ctx context.Context, rows sq.RowScanner) (*uint64, error) {
	var count uint64

	err := rows.Scan(&count)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	return &count, nil
}

func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *uuid.UUID) {

	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if Id is not set, generate a new UUID
	if item.Id == uuid.Nil {
		item.Id = uuid.New()
	}

	// set Id to context
	ctx = context.WithValue(ctx, "itemId", item.Id)

	insertItem := m.qb.Insert(m.table).Columns(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("ProductId"),
		m.mapFieldToDBColumn("TagTypeId"),
		m.mapFieldToDBColumn("TagSelectId"),
		m.mapFieldToDBColumn("Value"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
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

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.mapFieldToDBColumn("ProductId"), item.ProductId).
		Set(m.mapFieldToDBColumn("TagTypeId"), item.TagTypeId).
		Set(m.mapFieldToDBColumn("TagSelectId"), item.TagSelectId).
		Set(m.mapFieldToDBColumn("Value"), item.Value).
		Set(m.mapFieldToDBColumn("Active"), item.Active).
		Set(m.mapFieldToDBColumn("SortOrder"), item.SortOrder).
		Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").
		Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}

func (m *Model) makePatchStatement(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		statement = statement.Set(m.mapFieldToDBColumn(field), value)
	}

	return statement.Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}
