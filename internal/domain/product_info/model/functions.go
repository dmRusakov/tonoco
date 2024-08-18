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
		return nil, err
	}

	// return the Item
	return m.scanOneRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter, isUpdateFilter bool) (*map[string]Item, *uint64, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		err = errors.AddCode(err, "756592")
		return nil, nil, err
	}
	defer rows.Close()

	// iterate over the result set
	items := make(map[string]Item)
	idsMap := make(map[string]bool)
	urlsMap := make(map[string]bool)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			err = errors.AddCode(err, "646891")
			return nil, nil, err
		}

		items[item.Id] = *item

		// update filters if needed
		if isUpdateFilter {
			idsMap[item.Id] = true
			urlsMap[item.Url] = true
		}
	}

	// count the number of rows
	count := new(uint64)
	if filter.IsCount != nil && *filter.IsCount == true {
		rows, err = psql.List(ctx, m.client, m.makeCountStatementByFilter(filter))
		if err != nil {
			return nil, nil, err
		}

		defer rows.Close()
		for rows.Next() {
			count, err = m.scanCountRow(ctx, rows)
			if err != nil {
				err = errors.AddCode(err, "24696")
				return nil, nil, err
			}
		}
	}

	// update filters if needed
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
	return &items, count, nil
}

func (m *Model) Create(ctx context.Context, item *Item) (*string, error) {
	statement, id := m.makeInsertStatement(ctx, item)
	err := psql.Create(ctx, m.client, statement)
	if err != nil {
		err = errors.AddCode(err, "688828")
		return nil, err
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
		err = errors.AddCode(err, "229830")
		return err
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
		err = errors.AddCode(err, "979305")
		return err
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
		err = errors.AddCode(err, "58098")
		return err
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
		err = errors.AddCode(err, "665945")
		return nil, err
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
		err = errors.AddCode(err, "130004")
		return nil, err
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
		err = errors.AddCode(err, "656820")
		return nil, err
	}

	return order, nil
}
