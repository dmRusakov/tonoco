package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
)

type Item = entity.Price
type Filter = entity.PriceFilter

// fieldMap
var fieldMap = map[string]string{
	"ID":          "id",
	"ProductID":   "product_id",
	"PriceTypeID": "price_type_id",
	"CurrencyID":  "currency_id",
	"WarehouseID": "warehouse_id",
	"StoreId":     "store_id",
	"Price":       "price",
	"SortOrder":   "sort_order",
	"Active":      "active",
	"CreatedAt":   "created_at",
	"CreatedBy":   "created_by",
	"UpdatedAt":   "updated_at",
	"UpdatedBy":   "updated_by",
}

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		fieldMap["ID"],
		fieldMap["ProductID"],
		fieldMap["PriceTypeID"],
		fieldMap["CurrencyID"],
		fieldMap["WarehouseID"],
		fieldMap["StoreId"],
		fieldMap["Price"],
		fieldMap["SortOrder"],
		fieldMap["Active"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).From(m.table + " p")
}

// make Get statement
func (m *Model) makeGetStatement(
	id *string,
	productID *string,
	priceTypeID *string,
	currencyID *string,
	warehouseID *string,
	storeId *string,
) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if id != nil {
		statement = statement.Where("id = ?", *id)
	}

	// productID
	if productID != nil {
		statement = statement.Where("product_id = ?", *productID)
	}

	// priceTypeID
	if priceTypeID != nil {
		statement = statement.Where("price_type_id = ?", *priceTypeID)
	}

	// currencyID
	if currencyID != nil {
		statement = statement.Where("currency_id = ?", *currencyID)
	}

	// warehouseID
	if warehouseID != nil {
		statement = statement.
			Where("warehouse_id = ?", *warehouseID).
			Where("store_id IS NULL").
			OrderBy("store_id")
	}

	// storeId
	if storeId != nil {
		statement = statement.
			Where("store_id = ?", *storeId).
			Where("warehouse_id IS NULL").
			OrderBy("warehouse_id")
	}

	return statement
}

// makeStatementByFilter
func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.IDs != nil && len(*filter.IDs) > 0 {
		statement = statement.Where(sq.Eq{fieldMap["ID"]: *filter.IDs})
	}

	// ProductID
	if filter.ProductID != nil {
		statement = statement.Where(sq.Eq{fieldMap["ProductID"]: *filter.ProductID})
	}

	// CurrencyID
	if filter.CurrencyID != nil {
		statement = statement.Where(sq.Eq{fieldMap["CurrencyID"]: *filter.CurrencyID})
	}

	// WarehouseID
	if filter.WarehouseID != nil {
		statement = statement.Where(sq.Eq{fieldMap["WarehouseID"]: *filter.WarehouseID})
	}

	// StoreId
	if filter.StoreId != nil {
		statement = statement.Where(sq.Eq{fieldMap["StoreId"]: *filter.StoreId}).Where("warehouse_id IS NULL").OrderBy("warehouse_id")
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{fieldMap["Active"]: *filter.Active})
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = &Item{}
	err := rows.Scan(
		&item.ID,
		&item.ProductID,
		&item.PriceTypeID,
		&item.CurrencyID,
		&item.WarehouseID,
		&item.StoreId,
		&item.Price,
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
		fieldMap["ProductID"],
		fieldMap["PriceTypeID"],
		fieldMap["CurrencyID"],
		fieldMap["WarehouseID"],
		fieldMap["StoreId"],
		fieldMap["Price"],
		fieldMap["Active"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).Values(
		item.ID,
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
	return &insertItem, &item.ID
}

// makeUpdateStatement
func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(fieldMap["ProductID"], item.ProductID).
		Set(fieldMap["PriceTypeID"], item.PriceTypeID).
		Set(fieldMap["CurrencyID"], item.CurrencyID).
		Set(fieldMap["WarehouseID"], item.WarehouseID).
		Set(fieldMap["StoreId"], item.StoreId).
		Set(fieldMap["Price"], item.Price).
		Set(fieldMap["Active"], item.Active).
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
