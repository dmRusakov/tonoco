package model

import (
	"context"
	"fmt"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

func (m *Model) Get(
	ctx context.Context,
	id *string,
	productID *string,
	priceTypeID *string,
	currencyID *string,
	warehouseID *string,
	storeId *string,
) (*Item, error) {
	rows, err := psql.Get(ctx, m.client, m.makeGetStatement(
		id,
		productID,
		priceTypeID,
		currencyID,
		warehouseID,
		storeId,
	))
	if err != nil {
		return nil, err
	}

	// return the Item
	return m.scanOneRow(ctx, rows)
}

func (m *Model) List(ctx context.Context, filter *Filter) ([]*Item, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the result set
	var items []*Item
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (m *Model) Create(ctx context.Context, item *Item) (*string, error) {
	statement, id := m.makeInsertStatement(ctx, item)
	return id, psql.Create(
		ctx,
		m.client,
		statement,
	)
}

func (m *Model) Update(ctx context.Context, item *Item) (err error) {
	return psql.Update(
		ctx,
		m.client,
		m.makeUpdateStatement(ctx, item).Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), item.ID),
	)
}

func (m *Model) Patch(ctx context.Context, id *string, fields *map[string]interface{}) error {
	return psql.Update(
		ctx,
		m.client,
		m.makePatchStatement(ctx, id, fields),
	)
}

func (m *Model) Delete(ctx context.Context, id *string) error {
	return psql.Delete(
		ctx,
		m.client,
		m.qb.Delete(m.table).Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), id),
	)
}

func (m *Model) UpdatedAt(ctx context.Context, id *string) (*time.Time, error) {
	return psql.UpdatedAt(
		ctx,
		m.client,
		m.qb.Select(fieldMap["UpdatedAt"]).From(m.table).Where("id = ?", id),
	)
}

func (m *Model) TableIndexCount(ctx context.Context) (*uint64, error) {
	return psql.TableIndexCount(
		ctx,
		m.client,
		m.qb.Select("n_tup_upd").From("pg_stat_user_tables").Where("relname = ?", m.table),
	)
}

func (m *Model) MaxSortOrder(ctx context.Context) (*uint64, error) {
	return psql.MaxSortOrder(
		ctx,
		m.client,
		m.qb,
		&m.table,
	)
}
