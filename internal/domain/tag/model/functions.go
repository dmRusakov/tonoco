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
		return nil, errors.AddCode(err, "59867")
	}

	// return the Item
	return m.scanOneRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		return nil, errors.AddCode(err, "317886")
	}
	defer rows.Close()

	// iterate over the result set
	items := make(map[uuid.UUID]Item)
	ids := make([]uuid.UUID, 0)
	productIds := make([]uuid.UUID, 0)
	tagTypeIds := make([]uuid.UUID, 0)
	tagSelectIds := make([]uuid.UUID, 0)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, errors.AddCode(err, "890721")
		}
		items[item.Id] = *item

		// update filters if needed
		if filter.IsUpdateFilter != nil && *filter.IsUpdateFilter {
			ids = append(ids, item.Id)
			productIds = append(productIds, item.ProductId)
			tagTypeIds = append(tagTypeIds, item.TagTypeId)
			tagSelectIds = append(tagSelectIds, item.TagSelectId)
		}
	}

	// update filters if needed
	if filter.IsUpdateFilter != nil && *filter.IsUpdateFilter {
		// remove duplicates form productIds
		productIds = entity.RemoveDuplicates(productIds, true)
		tagTypeIds = entity.RemoveDuplicates(tagTypeIds, true)
		tagSelectIds = entity.RemoveDuplicates(tagSelectIds, true)

		// update filter
		filter.Ids = &ids
		filter.ProductIds = &productIds
		filter.TagTypeIds = &tagTypeIds
		filter.TagSelectIds = &tagSelectIds
	}

	// return the items
	return &items, nil
}

func (m *Model) Create(ctx context.Context, item *Item) (*uuid.UUID, error) {
	statement, id := m.makeInsertStatement(ctx, item)
	err := psql.Create(ctx, m.client, statement)
	if err != nil {
		return nil, errors.AddCode(err, "250497")
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

func (m *Model) Patch(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) error {
	err := psql.Update(
		ctx,
		m.client,
		m.makePatchStatement(ctx, id, fields),
	)

	if err != nil {
		return errors.AddCode(err, "897914")
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
		return errors.AddCode(err, "408992")
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
		return nil, errors.AddCode(err, "430148")
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
		return nil, errors.AddCode(err, "811768")
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
		return nil, errors.AddCode(err, "175851")
	}

	return order, nil
}
