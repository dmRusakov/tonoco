package model

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/google/uuid"
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

func (m *Model) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		return nil, errors.AddCode(err, "886656")
	}
	defer rows.Close()

	// iterate over the result set
	items := make(map[uuid.UUID]Item)
	ids := make([]uuid.UUID, 0)
	productIDs := make([]uuid.UUID, 0)
	currencyIDs := make([]uuid.UUID, 0)
	warehouseIDs := make([]uuid.UUID, 0)
	storeIDs := make([]uuid.UUID, 0)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, errors.AddCode(err, "295128")
		}

		items[item.Id] = *item

		// update filters if needed
		if filter.IsUpdateFilter != nil && *filter.IsUpdateFilter {
			ids = append(ids, item.Id)
			productIDs = append(productIDs, item.ProductID)
			currencyIDs = append(currencyIDs, item.CurrencyID)
			warehouseIDs = append(warehouseIDs, item.WarehouseID)
			storeIDs = append(storeIDs, item.StoreID)
		}
	}

	// update filters if needed
	if filter.IsUpdateFilter != nil && *filter.IsUpdateFilter {
		// remove duplicates form productIDs
		productIDs = entity.RemoveDuplicates(productIDs, true)
		currencyIDs = entity.RemoveDuplicates(currencyIDs, true)
		warehouseIDs = entity.RemoveDuplicates(warehouseIDs, true)
		storeIDs = entity.RemoveDuplicates(storeIDs, true)

		// update filter
		filter.Ids = &ids
		filter.ProductIds = &productIDs
		filter.CurrencyIds = &currencyIDs
		filter.WarehouseIds = &warehouseIDs
		filter.StoreIds = &storeIDs
	}

	return &items, nil
}

func (m *Model) Create(ctx context.Context, item *Item) (*uuid.UUID, error) {
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

func (m *Model) Patch(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) error {
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

func (m *Model) Delete(ctx context.Context, id *uuid.UUID) error {
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

func (m *Model) UpdatedAt(ctx context.Context, id *uuid.UUID) (*time.Time, error) {
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
