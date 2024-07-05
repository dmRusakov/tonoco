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
		m.fieldMap("ProductIds"),
		m.fieldMap("PriceTypeIds"),
		m.fieldMap("CurrencyIds"),
		m.fieldMap("WarehouseIds"),
		m.fieldMap("StoreIds"),
		m.fieldMap("Price"),
		m.fieldMap("SortOrder"),
		m.fieldMap("Active"),
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

	// productID
	if filter.ProductIds != nil {
		statement = statement.Where(m.fieldMap("ProductIds")+" = ?", (*filter.ProductIds)[0])
	}

	// priceTypeID
	if filter.PriceTypeIds != nil {
		statement = statement.Where(m.fieldMap("PriceTypeIds")+" = ?", (*filter.PriceTypeIds)[0])
	}

	// CurrencyIds
	if filter.CurrencyIds != nil {
		statement = statement.Where(m.fieldMap("CurrencyIds")+" = ?", (*filter.CurrencyIds)[0])
	}

	// WarehouseIds
	if filter.WarehouseIds != nil {
		statement = statement.Where(m.fieldMap("WarehouseIds")+" = ?", (*filter.WarehouseIds)[0])
	}

	// StoreIds
	if filter.StoreIds != nil {
		statement = statement.Where(m.fieldMap("StoreIds")+" = ?", (*filter.StoreIds)[0])
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
	if filter.Ids != nil && len(*filter.Ids) > 0 {
		statement = statement.Where(sq.Eq{m.fieldMap("Id"): *filter.Ids})
	}

	// ProductIds
	if filter.ProductIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("ProductIds"): *filter.ProductIds})
	}

	// CurrencyIds
	if filter.CurrencyIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("CurrencyIds"): *filter.CurrencyIds})
	}

	// WarehouseIds
	if filter.WarehouseIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("WarehouseIds"): *filter.WarehouseIds})
	}

	// StoreIds
	if filter.StoreIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("StoreIds"): *filter.StoreIds}).Where("warehouse_id IS NULL").OrderBy("warehouse_id")
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.fieldMap("Price")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.fieldMap(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, productID, priceTypeID, currencyID, warehouseID, storeId, createdBy, updatedBy sql.NullString
	var active sql.NullBool
	var price sql.NullInt64
	var createdAt, updatedAt sql.NullString

	err := rows.Scan(
		&id,
		&productID,
		&priceTypeID,
		&currencyID,
		&warehouseID,
		&storeId,
		&price,
		&active,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	var item = Item{}

	if id.Valid {
		item.Id = id.String
	}

	if productID.Valid {
		item.ProductID = productID.String
	}

	if priceTypeID.Valid {
		item.PriceTypeID = priceTypeID.String
	}

	if currencyID.Valid {
		item.CurrencyID = currencyID.String
	}

	if warehouseID.Valid {
		item.WarehouseID = warehouseID.String
	}

	if storeId.Valid {
		item.StoreId = storeId.String
	}

	if price.Valid {
		item.Price = uint64(price.Int64)
	}

	if active.Valid {
		item.Active = active.Bool
	}

	if createdAt.Valid {
		item.CreatedAt = createdAt.String
	}

	if createdBy.Valid {
		item.CreatedBy = createdBy.String
	}

	if updatedAt.Valid {
		item.UpdatedAt = updatedAt.String
	}

	if updatedBy.Valid {
		item.UpdatedBy = updatedBy.String
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
		m.fieldMap("ProductIds"),
		m.fieldMap("PriceTypeIds"),
		m.fieldMap("CurrencyIds"),
		m.fieldMap("WarehouseIds"),
		m.fieldMap("StoreIds"),
		m.fieldMap("Price"),
		m.fieldMap("Active"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
	).Values(
		item.Id,
		item.ProductID,
		item.PriceTypeID,
		item.CurrencyID,
		item.WarehouseID,
		item.StoreId,
		item.Price,
		item.Active,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	// get itemId from context
	return &insertItem, &item.Id
}

// makeUpdateStatement
func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.fieldMap("ProductIds"), item.ProductID).
		Set(m.fieldMap("PriceTypeIds"), item.PriceTypeID).
		Set(m.fieldMap("CurrencyIds"), item.CurrencyID).
		Set(m.fieldMap("WarehouseIds"), item.WarehouseID).
		Set(m.fieldMap("StoreIds"), item.StoreId).
		Set(m.fieldMap("Price"), item.Price).
		Set(m.fieldMap("Active"), item.Active).
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
