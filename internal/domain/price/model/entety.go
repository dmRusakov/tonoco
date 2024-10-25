package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
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
		m.mapFieldToDBColumn("ProductID"),
		m.mapFieldToDBColumn("PriceTypeID"),
		m.mapFieldToDBColumn("CurrencyID"),
		m.mapFieldToDBColumn("WarehouseID"),
		m.mapFieldToDBColumn("StoreID"),
		m.mapFieldToDBColumn("Price"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("StartDate"),
		m.mapFieldToDBColumn("EndDate"),
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
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("ProductID"): *filter.ProductIds})
	}

	// PriceTypeIds
	if filter.PriceTypeIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("PriceTypeID"): *filter.PriceTypeIds})
	}

	// CurrencyIds
	if filter.CurrencyIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("CurrencyID"): *filter.CurrencyIds})
	}

	// WarehouseIds
	if filter.WarehouseIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("WarehouseID"): *filter.WarehouseIds})
	}

	// StoreIds
	if filter.StoreIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("StoreID"): *filter.StoreIds})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Price")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
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
		filter.OrderBy = entity.StringPtr("SortOrder")
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = entity.StringPtr("ASC")
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

func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, productId, priceTypeId, currencyId, warehouseId, storeId, createdBy, updatedBy sql.NullString
	var active sql.NullBool
	var price sql.NullFloat64
	var sortOrder sql.NullInt64
	var startDate, endData, createdAt, updatedAt sql.NullTime

	err := rows.Scan(
		&id,
		&productId,
		&priceTypeId,
		&currencyId,
		&warehouseId,
		&storeId,
		&price,
		&sortOrder,
		&active,
		&startDate,
		&endData,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "802311")
	}

	var item = Item{}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if productId.Valid {
		item.ProductID = uuid.MustParse(productId.String)
	}

	if priceTypeId.Valid {
		item.PriceTypeID = uuid.MustParse(priceTypeId.String)
	}

	if currencyId.Valid {
		item.CurrencyID = uuid.MustParse(currencyId.String)
	}

	if warehouseId.Valid {
		item.WarehouseID = uuid.MustParse(warehouseId.String)
	}

	if storeId.Valid {
		item.StoreID = uuid.MustParse(storeId.String)
	}

	if price.Valid {
		item.Price = price.Float64
	}

	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
	}

	if active.Valid {
		item.Active = active.Bool
	}

	if startDate.Valid {
		item.StartDate = startDate.Time
	}

	if endData.Valid {
		item.EndDate = endData.Time
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

	insertItem := m.qb.Insert(m.table).Columns(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("ProductIds"),
		m.mapFieldToDBColumn("PriceTypeIds"),
		m.mapFieldToDBColumn("CurrencyIds"),
		m.mapFieldToDBColumn("WarehouseIds"),
		m.mapFieldToDBColumn("StoreIds"),
		m.mapFieldToDBColumn("Price"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("StartDate"),
		m.mapFieldToDBColumn("EndDate"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.ProductID,
		item.PriceTypeID,
		item.CurrencyID,
		item.WarehouseID,
		item.StoreID,
		item.Price,
		item.Active,
		item.StartDate,
		item.EndDate,
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
		Set(m.mapFieldToDBColumn("ProductIds"), item.ProductID).
		Set(m.mapFieldToDBColumn("PriceTypeIds"), item.PriceTypeID).
		Set(m.mapFieldToDBColumn("CurrencyIds"), item.CurrencyID).
		Set(m.mapFieldToDBColumn("WarehouseIds"), item.WarehouseID).
		Set(m.mapFieldToDBColumn("StoreIds"), item.StoreID).
		Set(m.mapFieldToDBColumn("Price"), item.Price).
		Set(m.mapFieldToDBColumn("Active"), item.Active).
		Set(m.mapFieldToDBColumn("StartDate"), item.StartDate).
		Set(m.mapFieldToDBColumn("EndDate"), item.EndDate).
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
