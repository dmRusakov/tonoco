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
		m.fieldMap("ProductID"),
		m.fieldMap("PriceTypeID"),
		m.fieldMap("CurrencyID"),
		m.fieldMap("WarehouseID"),
		m.fieldMap("StoreID"),
		m.fieldMap("Price"),
		m.fieldMap("SortOrder"),
		m.fieldMap("Active"),
		m.fieldMap("StartDate"),
		m.fieldMap("EndDate"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
	).From(m.table + " p")
}

func (m *Model) makeStatementWithFilter(filter *Filter) sq.SelectBuilder {
	statement := m.makeStatement()

	// Ids
	if filter.Ids != nil && len(*filter.Ids) > 0 {
		statement = statement.Where(sq.Eq{m.fieldMap("Id"): *filter.Ids})
	}

	// ProductIds
	if filter.ProductIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("ProductID"): *filter.ProductIds})
	}

	// PriceTypeIds
	if filter.PriceTypeIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("PriceTypeID"): *filter.PriceTypeIds})
	}

	// CurrencyIds
	if filter.CurrencyIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("CurrencyID"): *filter.CurrencyIds})
	}

	// WarehouseIds
	if filter.WarehouseIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("WarehouseID"): *filter.WarehouseIds})
	}

	// StoreIds
	if filter.StoreIds != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("StoreID"): *filter.StoreIds})
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

	return statement
}

// make Get statement
func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	return m.makeStatementWithFilter(filter)
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
	statement := m.makeStatementWithFilter(filter)

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.fieldMap(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

// makeCountStatementByFilter - make count statement by filter for pagination
func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	// Build query
	statement := m.qb.Select("COUNT(*)").From(m.table)

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

	return statement
}

// scanOneRow
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
		err = errors.AddCode(err, "cdczt0")
		return nil, err
	}

	var item = Item{}

	if id.Valid {
		item.Id = id.String
	}

	if productId.Valid {
		item.ProductID = productId.String
	}

	if priceTypeId.Valid {
		item.PriceTypeID = priceTypeId.String
	}

	if currencyId.Valid {
		item.CurrencyID = currencyId.String
	}

	if warehouseId.Valid {
		item.WarehouseID = warehouseId.String
	}

	if storeId.Valid {
		item.StoreID = storeId.String
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
		item.CreatedBy = createdBy.String
	}

	if updatedAt.Valid {
		item.UpdatedAt = updatedAt.Time
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
		m.fieldMap("StartDate"),
		m.fieldMap("EndDate"),
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

// makeUpdateStatement
func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.fieldMap("ProductIds"), item.ProductID).
		Set(m.fieldMap("PriceTypeIds"), item.PriceTypeID).
		Set(m.fieldMap("CurrencyIds"), item.CurrencyID).
		Set(m.fieldMap("WarehouseIds"), item.WarehouseID).
		Set(m.fieldMap("StoreIds"), item.StoreID).
		Set(m.fieldMap("Price"), item.Price).
		Set(m.fieldMap("Active"), item.Active).
		Set(m.fieldMap("StartDate"), item.StartDate).
		Set(m.fieldMap("EndDate"), item.EndDate).
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

func convertToUUIDSlice(strSlice []string) ([]uuid.UUID, error) {
	uuidSlice := make([]uuid.UUID, len(strSlice))
	for i, str := range strSlice {
		u, err := uuid.Parse(str)
		if err != nil {
			return nil, err
		}
		uuidSlice[i] = u
	}
	return uuidSlice, nil
}
