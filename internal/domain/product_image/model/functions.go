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

func (m *Model) Get(ctx context.Context, filter *ProductImageFilter) (*ProductImage, error) {
	row, err := psql.Get(ctx, m.client, m.makeGetStatement(filter))
	if err != nil {
		return nil, errors.AddCode(err, "467009")
	}

	return m.scanOneRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *ProductImageFilter, isUpdateFilter bool) (*map[uuid.UUID]ProductImage, *uint64, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		return nil, nil, errors.AddCode(err, "953590")
	}
	defer rows.Close()

	items := make(map[uuid.UUID]ProductImage)
	ids := make([]uuid.UUID, 0)
	productIDs := make([]uuid.UUID, 0)
	imageIDs := make([]uuid.UUID, 0)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, nil, errors.AddCode(err, "389234")
		}

		items[item.Id] = *item

		if isUpdateFilter {
			ids = append(ids, item.Id)
			productIDs = append(productIDs, item.ProductId)
			imageIDs = append(imageIDs, item.ImageId)
		}
	}

	if isUpdateFilter {
		// remove duplicates
		productIDs = entity.RemoveDuplicates(productIDs, true)
		imageIDs = entity.RemoveDuplicates(imageIDs, true)

		filter.Ids = &ids
		filter.ProductIds = &productIDs
		filter.ImageIds = &imageIDs
	}

	return &items, nil, nil
}

func (m *Model) Create(ctx context.Context, item *ProductImage) (*uuid.UUID, error) {
	statement, id := m.makeInsertStatement(ctx, item)
	err := psql.Create(ctx, m.client, statement)
	if err != nil {
		return nil, errors.AddCode(err, "235131")
	}

	return id, nil
}

func (m *Model) Update(ctx context.Context, item *ProductImage) error {
	err := psql.Update(
		ctx,
		m.client,
		m.makeUpdateStatement(ctx, item).Where(fmt.Sprintf("%s = ?", m.fieldMap("Id")), item.Id),
	)

	if err != nil {
		return errors.AddCode(err, "115681")
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
		return errors.AddCode(err, "976385")
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
		return errors.AddCode(err, "129688")
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
		return nil, errors.AddCode(err, "713559")
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
		return nil, errors.AddCode(err, "891898")
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
		return nil, errors.AddCode(err, "605995")
	}

	return order, nil
}
