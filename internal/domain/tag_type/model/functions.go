package model

import (
	"context"
	"fmt"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

func (m *Model) Get(ctx context.Context, filter *Filter) (*Item, error) {
	row, err := psql.Get(ctx, m.client, m.makeGetStatement(filter))
	if err != nil {
		return nil, err
	}

	// return the Item
	return m.scanOneRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter, isUpdateFilter bool) (*map[string]Item, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the result set
	items := make(map[string]Item)
	idsMap := make(map[string]bool)
	urlsMap := make(map[string]bool)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, err
		}

		items[item.Id] = *item

		// update filters if needed
		if isUpdateFilter {
			idsMap[item.Id] = true
			urlsMap[item.Url] = true
		}
	}

	// update filter if needed
	if isUpdateFilter {
		ids := make([]string, 0, len(idsMap))
		for id := range idsMap {
			ids = append(ids, id)
		}
		urls := make([]string, 0, len(urlsMap))
		for url := range urlsMap {
			urls = append(urls, url)
		}

		// update filter
		filter.Ids = &ids
		filter.Urls = &urls
	}

	// return the Items
	return &items, nil
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
		m.makeUpdateStatement(ctx, item).Where(fmt.Sprintf("%s = ?", m.fieldMap("Id")), item.Id),
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
		m.qb.Delete(m.table).Where(fmt.Sprintf("%s = ?", m.fieldMap("Id")), id),
	)
}

func (m *Model) UpdatedAt(ctx context.Context, id *string) (*time.Time, error) {
	return psql.UpdatedAt(
		ctx,
		m.client,
		m.qb.Select(m.fieldMap("UpdatedAt")).From(m.table).Where("id = ?", id),
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
