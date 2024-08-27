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
		return nil, errors.AddCode(err, "947025")
	}

	// return the Item
	return m.scanOneRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter, isUpdateFilter bool) (*map[string]Item, *uint64, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		return nil, nil, errors.AddCode(err, "231402")
	}
	defer rows.Close()

	// iterate over the result set
	items := make(map[string]Item)
	idsMap := make(map[string]bool)
	productIdsMap := make(map[string]bool)
	warehouseIdsMap := make(map[string]bool)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, nil, errors.AddCode(err, "943729")
		}

		items[item.Id] = *item

		// update filters if needed
		if isUpdateFilter {
			idsMap[item.Id] = true
			productIdsMap[item.ProductId] = true
			warehouseIdsMap[item.WarehouseId] = true
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

		warehouseIds := make([]string, 0, len(warehouseIdsMap))
		for warehouseId := range warehouseIdsMap {
			warehouseIds = append(warehouseIds, warehouseId)
		}

		// update filter
		filter.Ids = &ids
		filter.ProductIds = &productIds
		filter.WarehouseIds = &warehouseIds
	}

	return &items, nil, nil
}

func (m *Model) Create(ctx context.Context, item *Item) (*string, error) {
	statement, id := m.makeInsertStatement(ctx, item)
	err := psql.Create(ctx, m.client, statement)
	if err != nil {
		return nil, errors.AddCode(err, "102122")
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
		return errors.AddCode(err, "410295")
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
		return errors.AddCode(err, "508876")
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
		return errors.AddCode(err, "707629")
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
		return nil, errors.AddCode(err, "176261")
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
		return nil, errors.AddCode(err, "899195")
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
		return nil, errors.AddCode(err, "217631")
	}

	return order, nil
}
