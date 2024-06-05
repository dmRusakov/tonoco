package model

import (
	"context"
	"fmt"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

func (m *Model) Get(ctx context.Context, id *string, url *string) (*Item, error) {
	rows, err := psql.Get(ctx, m.client, m.makeGetStatement(id, url))
	if err != nil {
		return nil, err
	}

	// return the Item
	return m.scanOneRow(ctx, rows)
}

func (m *Model) List(ctx context.Context, filter *Filter) (*map[string]Item, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	// iterate over the result set
	items := make(map[string]Item)
	var IDs, ProductIDs, TagTypeIDs, TagSelectIDs []string
	idsMap := make(map[string]bool)
	productIDsMap := make(map[string]bool)
	tagTypeIDsMap := make(map[string]bool)
	tagSelectIDsMap := make(map[string]bool)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, err
		}
		items[item.ID] = *item
		IDs = append(IDs, item.ID)
		idsMap[item.ID] = true

		if item.ProductId != "" && !productIDsMap[item.ProductId] {
			ProductIDs = append(ProductIDs, item.ProductId)
			productIDsMap[item.ProductId] = true
		}
		if item.TagTypeId != "" && !tagTypeIDsMap[item.TagTypeId] {
			TagTypeIDs = append(TagTypeIDs, item.TagTypeId)
			tagTypeIDsMap[item.TagTypeId] = true
		}
		if item.TagSelectId != "" && !tagSelectIDsMap[item.TagSelectId] {
			TagSelectIDs = append(TagSelectIDs, item.TagSelectId)
			tagSelectIDsMap[item.TagSelectId] = true
		}
	}

	// update the filter
	filter.IDs = &IDs
	filter.ProductIDs = &ProductIDs
	filter.TagTypeIDs = &TagTypeIDs
	filter.TagSelectIDs = &TagSelectIDs

	// done
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
		m.makeUpdateStatement(ctx, item).Where(fmt.Sprintf("%s = ?", m.fieldMap("ID")), item.ID),
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
		m.qb.Delete(m.table).Where(fmt.Sprintf("%s = ?", m.fieldMap("ID")), id),
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
