package model

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/google/uuid"
	"time"
)

type Item = entity.ProductImage
type Filter = entity.ProductImageFilter

type Storage interface {
	Get(context.Context, *Filter) (*Item, error)
	List(context.Context, *Filter) (*map[uuid.UUID]Item, error)
	Create(context.Context, *Item) (*uuid.UUID, error)
	Update(context.Context, *Item) error
	Patch(context.Context, *uuid.UUID, *map[string]interface{}) error
	UpdatedAt(context.Context, *uuid.UUID) (*time.Time, error)
	MaxSortOrder(context.Context) (*uint64, error)
	TableIndexCount(context.Context) (*uint64, error)
	Delete(context.Context, *uuid.UUID) error
}

type Model struct {
	table       string
	qb          sq.StatementBuilderType
	client      psql.Client
	dbFieldCash map[string]string
}

func NewStorage(client psql.Client) *Model {
	return &Model{
		qb:          sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:      client,
		table:       "product_image",
		dbFieldCash: map[string]string{},
	}
}

func (m *Model) Get(ctx context.Context, filter *Filter) (*Item, error) {
	row, err := psql.Get(ctx, m.client, m.makeGetStatement(filter))
	if err != nil {
		return nil, errors.AddCode(err, "467009")
	}

	return m.scanOneRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(filter))
	if err != nil {
		return nil, errors.AddCode(err, "953590")
	}
	defer rows.Close()

	items := make(map[uuid.UUID]Item)
	ids := make([]uuid.UUID, 0)
	productIDs := make([]uuid.UUID, 0)
	imageIDs := make([]uuid.UUID, 0)
	for rows.Next() {
		item, err := m.scanOneRow(ctx, rows)
		if err != nil {
			return nil, err
		}

		if filter.IsIdsOnly == nil || !*filter.IsIdsOnly {
			items[item.Id] = *item
		}

		if filter.IsUpdateFilter != nil && *filter.IsUpdateFilter {
			ids = append(ids, item.Id)
			productIDs = append(productIDs, item.ProductId)
			imageIDs = append(imageIDs, item.ImageId)
		}
	}

	if filter.IsUpdateFilter != nil && *filter.IsUpdateFilter {
		// remove duplicates
		productIDs = entity.RemoveDuplicates(productIDs, filter.IsKeepIdsOrder != nil && *filter.IsKeepIdsOrder)
		imageIDs = entity.RemoveDuplicates(imageIDs, filter.IsKeepIdsOrder != nil && *filter.IsKeepIdsOrder)

		filter.Ids = &ids
		filter.ProductIds = &productIDs
		filter.ImageIds = &imageIDs
	}

	return &items, nil
}

func (m *Model) Create(ctx context.Context, item *Item) (*uuid.UUID, error) {
	statement, id := m.makeInsertStatement(ctx, item)
	err := psql.Create(ctx, m.client, statement)
	if err != nil {
		return nil, errors.AddCode(err, "235131")
	}

	return id, nil
}

func (m *Model) Update(ctx context.Context, item *Item) error {
	err := psql.Update(
		ctx,
		m.client,
		m.makeUpdateStatement(ctx, item).Where(fmt.Sprintf("%s = ?", m.mapFieldToDBColumn("Id")), item.Id),
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
		m.qb.Delete(m.table).Where(fmt.Sprintf("%s = ?", m.mapFieldToDBColumn("Id")), id),
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
		m.qb.Select(m.mapFieldToDBColumn("UpdatedAt")).From(m.table).Where("id = ?", id),
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
