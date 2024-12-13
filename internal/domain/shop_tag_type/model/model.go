package model

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/google/uuid"
	"time"
)

type Item = db.ShopTagType
type Filter = db.ShopTagTypeFilter

type Storage interface {
	Get(context.Context, *Filter) (*Item, error)
	List(context.Context, *Filter) (*map[uuid.UUID]Item, error)
	Ids(context.Context, *Filter) (*[]uuid.UUID, error)
	Create(context.Context, *Item) (*uuid.UUID, error)
	Update(context.Context, *Item) error
	Patch(context.Context, *uuid.UUID, *map[string]interface{}) error
	UpdatedAt(context.Context, *uuid.UUID) (*time.Time, error)
	MaxSortOrder(context.Context) (*uint64, error)
	TableIndexCount(context.Context) (*uint64, error)
	Delete(context.Context, *uuid.UUID) error

	makeStatement() sq.SelectBuilder
	makeGetStatement(*Filter) sq.SelectBuilder
	filterDTO(*Filter)
	makeStatementByFilter(sq.SelectBuilder, *Filter) sq.SelectBuilder
	makeCountStatementByFilter(*Filter) sq.SelectBuilder
	scanRow(context.Context, sq.RowScanner) (*Item, error)
	scanIdRow(context.Context, sq.RowScanner) (*uuid.UUID, error)
	scanCountRow(context.Context, sq.RowScanner) (*uint64, error)
	makeInsertStatement(context.Context, *Item) (*sq.InsertBuilder, *uuid.UUID)
	makeUpdateStatement(context.Context, *Item) sq.UpdateBuilder
	makePatchStatement(context.Context, *uuid.UUID, *map[string]interface{}) sq.UpdateBuilder
}

type Model struct {
	table   string
	qb      sq.StatementBuilderType
	client  psql.Client
	dbField map[string]string
}

func NewStorage(client psql.Client) *Model {
	return &Model{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
		table:  "shop_tag_type",
		dbField: map[string]string{
			"Id":        "id",
			"ShopId":    "shop_id",
			"TagTypeId": "tag_type_id",
			"Type":      "type",
			"Source":    "source",
			"SortOrder": "sort_order",
			"Active":    "active",
			"CreatedAt": "created_at",
			"CreatedBy": "created_by",
			"UpdatedAt": "updated_at",
			"UpdatedBy": "updated_by",
		},
	}
}

func (m *Model) Get(ctx context.Context, filter *Filter) (*Item, error) {
	// check filter
	m.filterDTO(filter)

	// get the row
	row, err := psql.Get(ctx, m.client, m.makeGetStatement(filter))
	if err != nil {
		return nil, err
	}

	// return the Item
	return m.scanRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	// check filter
	m.filterDTO(filter)

	// get the rows
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(m.makeStatement(), filter))
	if err != nil {
		return nil, errors.AddCode(err, "331081")
	}
	defer rows.Close()

	// iterate over the result set
	items := make(map[uuid.UUID]Item)
	ids := make([]uuid.UUID, 0)
	shopIds := make([]uuid.UUID, 0)
	tagTypeIds := make([]uuid.UUID, 0)

	for rows.Next() {
		item, err := m.scanRow(ctx, rows)
		if err != nil {
			return nil, err
		}

		if filter.DataConfig.IsIdsOnly == nil || !*filter.DataConfig.IsIdsOnly {
			items[item.Id] = *item
		}

		// update filters if needed
		if filter.DataConfig.IsUpdateFilter != nil && *filter.DataConfig.IsUpdateFilter {
			ids = append(ids, item.Id)
			shopIds = append(shopIds, item.ShopId)
			tagTypeIds = append(tagTypeIds, item.TagTypeId)
		}
	}

	// count the number of rows
	if filter.DataConfig.IsCount != nil && *filter.DataConfig.IsCount == true {
		rows, err = psql.List(ctx, m.client, m.makeCountStatementByFilter(filter))
		if err != nil {
			return nil, errors.AddCode(err, "419930")
		}

		defer rows.Close()
		for rows.Next() {
			filter.DataConfig.Count, err = m.scanCountRow(ctx, rows)
			if err != nil {
				return nil, err
			}
		}
	}

	// update filters if needed
	if filter.DataConfig.IsUpdateFilter != nil && *filter.DataConfig.IsUpdateFilter {
		filter.Ids = &ids
		filter.ShopIds = &shopIds
		filter.TagTypeIds = &tagTypeIds
	}

	// return the Items
	return &items, nil
}

func (m *Model) Ids(ctx context.Context, filter *Filter) (*[]uuid.UUID, error) {
	// check filter
	m.filterDTO(filter)

	// get the rows
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(m.makeStatement(), filter))
	if err != nil {
		return nil, errors.AddCode(err, "331080")
	}
	defer rows.Close()

	// iterate over the result set
	ids := make([]uuid.UUID, 0)
	for rows.Next() {
		id, err := m.scanIdRow(ctx, rows)
		if err != nil {
			return nil, err
		}

		ids = append(ids, *id)
	}

	// count the number of rows
	if filter.DataConfig.IsCount != nil && *filter.DataConfig.IsCount == true {
		rows, err = psql.List(ctx, m.client, m.makeCountStatementByFilter(filter))
		if err != nil {
			return nil, errors.AddCode(err, "283721")
		}

		defer rows.Close()
		for rows.Next() {
			filter.DataConfig.Count, err = m.scanCountRow(ctx, rows)
			if err != nil {
				return nil, err
			}
		}
	}

	return &ids, nil
}

func (m *Model) Create(ctx context.Context, item *Item) (*uuid.UUID, error) {
	statement, id := m.makeInsertStatement(ctx, item)
	err := psql.Create(ctx, m.client, statement)
	if err != nil {
		return nil, errors.AddCode(err, "572729")
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
		return errors.AddCode(err, "502551")
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
		return errors.AddCode(err, "486236")
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
		return errors.AddCode(err, "392784")
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
		return nil, errors.AddCode(err, "965059")
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
		return nil, errors.AddCode(err, "329909")
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
		return nil, errors.AddCode(err, "491286")
	}

	return order, nil
}
