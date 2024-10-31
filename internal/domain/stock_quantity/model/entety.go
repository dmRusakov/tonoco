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
		m.mapFieldToDBColumn("Quality"),
		m.mapFieldToDBColumn("WarehouseId"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table + " p")
}

// fillInFilter
func (m *Model) fillInFilter(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
	// Ids
	if filter.Ids != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
	}

	// ProductIds
	if filter.ProductIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("ProductId"): *filter.ProductIds})
	}

	// WarehouseId
	if filter.WarehouseIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("WarehouseId"): *filter.WarehouseIds})
	}

	return statement
}

// makeGetStatement - make get statement by filter
func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	return m.fillInFilter(m.makeStatement(), filter)
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
	statement := m.makeGetStatement(filter)

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.mapFieldToDBColumn(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	return m.fillInFilter(m.qb.Select("COUNT(*)").From(m.table), filter)
}

func (m *Model) scanOneRow(ctx context.Context, row sq.RowScanner) (*Item, error) {
	var id, productId, warehouseId, createdBy, updatedBy sql.NullString
	var quality sql.NullInt64
	var createdAt, updatedAt sql.NullTime

	err := row.Scan(
		&id,
		&productId,
		&quality,
		&warehouseId,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "")
	}

	var item = Item{}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if productId.Valid {
		item.ProductId = uuid.MustParse(productId.String)
	}

	if quality.Valid {
		item.Quality = quality.Int64
	}

	if warehouseId.Valid {
		item.WarehouseId = uuid.MustParse(warehouseId.String)
	}

	if createdAt.Valid {
		item.CreatedAt = createdAt.Time
	}

	if createdBy.Valid {
		item.CreatedBy = uuid.MustParse(createdBy.String)
	}

	if updatedAt.Valid {
		item.UpdatedAt = updatedAt.Time
	}

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

	// insert statement
	insertItem := m.qb.Insert(m.table).Columns(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("ProductId"),
		m.mapFieldToDBColumn("Quality"),
		m.mapFieldToDBColumn("WarehouseId"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.ProductId,
		item.Quality,
		item.WarehouseId,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	// get itemId from context
	return &insertItem, &item.Id
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.mapFieldToDBColumn("ProductId"), item.ProductId).
		Set(m.mapFieldToDBColumn("Quality"), item.Quality).
		Set(m.mapFieldToDBColumn("WarehouseId"), item.WarehouseId).
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
