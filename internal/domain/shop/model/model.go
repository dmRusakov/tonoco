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

type Item = db.Shop
type Filter = db.ShopFilter

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
	makeStatementByFilter(sq.SelectBuilder, *Filter) sq.SelectBuilder
	makeCountStatementByFilter(*Filter) sq.SelectBuilder
	scanRow(context.Context, sq.RowScanner) (*Item, error)
	scanIdRow(context.Context, sq.RowScanner) (*uuid.UUID, error)
	scanCountRow(context.Context, sq.RowScanner) (*uint64, error)
	makeInsertStatement(context.Context, *Item) (*sq.InsertBuilder, *uuid.UUID)
	makeUpdateStatement(context.Context, *Item) sq.UpdateBuilder
	makePatchStatement(context.Context, *uuid.UUID, *map[string]interface{}) sq.UpdateBuilder
}

// Model is a struct that contains the SQL statement builder and the PostgreSQL client.
type Model struct {
	table   string
	qb      sq.StatementBuilderType
	client  psql.Client
	dbField map[string]string
}

// NewStorage is a constructor function that returns a new instance of the Model.
func NewStorage(client psql.Client) *Model {
	return &Model{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
		table:  "shop",
		dbField: map[string]string{
			"Id":               "id",
			"Name":             "name",
			"SeoTitle":         "seo_title",
			"ShortDescription": "short_description",
			"Description":      "description",
			"Url":              "url",
			"ImageId":          "image_id",
			"HoverImageId":     "hover_image_id",
			"Page":             "page",
			"PerPage":          "per_page",
			"SortOrder":        "sort_order",
			"Active":           "active",
			"Prime":            "prime",
			"CreatedAt":        "created_at",
			"CreatedBy":        "created_by",
			"UpdatedAt":        "updated_at",
			"UpdatedBy":        "updated_by",
		},
	}
}

func (m *Model) Get(ctx context.Context, filter *Filter) (*Item, error) {
	row, err := psql.Get(ctx, m.client, m.makeGetStatement(filter))
	if err != nil {
		return nil, err
	}
	return m.scanRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(m.makeStatement(), filter))
	if err != nil {
		return nil, errors.AddCode(err, "523414")
	}
	defer rows.Close()

	// iterate over the result set
	items := make(map[uuid.UUID]Item)
	ids := make([]uuid.UUID, 0)
	urls := make([]string, 0)
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
			urls = append(urls, item.Url)
		}
	}

	// count the number of rows
	if filter.DataConfig.IsCount != nil && *filter.DataConfig.IsCount == true {
		rows, err = psql.List(ctx, m.client, m.makeCountStatementByFilter(filter))
		if err != nil {
			return nil, errors.AddCode(err, "420456")
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
		filter.Urls = &urls
	}

	// return the Items
	return &items, nil
}

func (m *Model) Ids(ctx context.Context, filter *Filter) (*[]uuid.UUID, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(m.makeIdsStatement(), filter))
	if err != nil {
		return nil, errors.AddCode(err, "415826")
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
			return nil, errors.AddCode(err, "22284")
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
		return nil, errors.AddCode(err, "459081")
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
		return errors.AddCode(err, "213982")
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
		return errors.AddCode(err, "706373")
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
		return errors.AddCode(err, "566352")
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
		return nil, errors.AddCode(err, "351796")
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
		return nil, errors.AddCode(err, "212061")
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
		return nil, errors.AddCode(err, "98560")
	}

	return order, nil
}