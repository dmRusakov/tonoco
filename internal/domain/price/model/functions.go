package model

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

func (m *Model) Get(ctx context.Context, filter *Filter) (*Item, error) {
	row, err := psql.Get(ctx, m.client, m.makeGetStatement(filter))
	if err != nil {
		return nil, errors.AddCode(err, "372351")
	}

	// return the Item
	return m.scanOneRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter, isUpdateFilter bool) (*map[string]Item, *uint64, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		return nil, nil, errors.AddCode(err, "886656")
	}
	defer rows.Close()

	// iterate over the result set
	items := make(map[string]Item)
	idsMap := make(map[string]bool)
	productIdsMap := make(map[string]bool)
	currencyIdsMap := make(map[string]bool)
	warehouseIdsMap := make(map[string]bool)
	storeIdsMap := make(map[string]bool)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, nil, errors.AddCode(err, "295128")
		}

		items[item.Id] = *item

		// update filters if needed
		if isUpdateFilter {
			idsMap[item.Id] = true
			productIdsMap[item.ProductID] = true
			currencyIdsMap[item.CurrencyID] = true
			warehouseIdsMap[item.WarehouseID] = true
			storeIdsMap[item.StoreID] = true
		}
	}

	// update filters if needed
	if isUpdateFilter {
		ids := make([]string, 0, len(idsMap))
		for id := range idsMap {
			ids = append(ids, id)
		}
		productIds := make([]string, 0, len(productIdsMap))
		for productId := range productIdsMap {
			productIds = append(productIds, productId)
		}
		currencyIds := make([]string, 0, len(currencyIdsMap))
		for currencyId := range currencyIdsMap {
			currencyIds = append(currencyIds, currencyId)
		}
		warehouseIds := make([]string, 0, len(warehouseIdsMap))
		for warehouseId := range warehouseIdsMap {
			warehouseIds = append(warehouseIds, warehouseId)
		}
		storeIds := make([]string, 0, len(storeIdsMap))
		for storeId := range storeIdsMap {
			storeIds = append(storeIds, storeId)
		}

		// update filter
		filter.Ids = &ids
		filter.ProductIds = &productIds
		filter.CurrencyIds = &currencyIds
		filter.WarehouseIds = &warehouseIds
		filter.StoreIds = &storeIds
	}

	return &items, nil, nil
}

func (m *Model) Create(ctx context.Context, item *Item) (*string, error) {
	statement, id := m.makeInsertStatement(ctx, item)
	err := psql.Create(ctx, m.client, statement)
	if err != nil {
		return nil, errors.AddCode(err, "239908")
	}

	return id, nil
}

func (m *Model) Update(ctx context.Context, item *Item) error {
	err := psql.Update(
		ctx,
		m.client,
		m.makeUpdateStatement(ctx, item).Where(fmt.Sprintf("%s = ?", m.fieldMap("Id")), item.Id),
	)

	if err != nil {
		return errors.AddCode(err, "940857")
	}

	return nil
}

func (m *Model) Patch(ctx context.Context, id *string, fields *map[string]interface{}) error {
	err := psql.Update(
		ctx,
		m.client,
		m.makePatchStatement(ctx, id, fields),
	)

	if err != nil {
		return errors.AddCode(err, "602218")
	}

	return nil
}

func (m *Model) Delete(ctx context.Context, id *string) error {
	err := psql.Delete(
		ctx,
		m.client,
		m.qb.Delete(m.table).Where(fmt.Sprintf("%s = ?", m.fieldMap("Id")), id),
	)

	if err != nil {
		return errors.AddCode(err, "334950")
	}

	return nil
}

func (m *Model) UpdatedAt(ctx context.Context, id *string) (*time.Time, error) {
	at, err := psql.UpdatedAt(
		ctx,
		m.client,
		m.qb.Select(m.fieldMap("UpdatedAt")).From(m.table).Where("id = ?", id),
	)

	if err != nil {
		return nil, errors.AddCode(err, "865717")
	}

	return at, nil
}

func (m *Model) TableIndexCount(ctx context.Context) (*uint64, error) {
	count, err := psql.TableIndexCount(
		ctx,
		m.client,
		m.qb.Select("n_tup_upd").From("pg_stat_user_tables").Where("relname = ?", m.table),
	)

	if err != nil {
		return nil, errors.AddCode(err, "474075")
	}

	return count, nil
}

func (m *Model) MaxSortOrder(ctx context.Context) (*uint64, error) {
	order, err := psql.MaxSortOrder(
		ctx,
		m.client,
		m.qb,
		&m.table,
	)

	if err != nil {
		return nil, errors.AddCode(err, "268983")
	}

	return order, nil
}
