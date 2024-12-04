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

type Item = db.Text
type Filter = db.TextFilter

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
		table:  "text",
		dbField: map[string]string{
			"Id":        "id",
			"Language":  "language",
			"Source":    "source",
			"SubSource": "sub_source",
			"SourceId":  "source_id",
			"Text":      "text",
			"Active":    "active",
			"CreatedAt": "created_at",
			"CreatedBy": "created_by",
			"UpdatedAt": "updated_at",
			"UpdatedBy": "updated_by",
		},
	}
}

func (m *Model) Get(ctx context.Context, filter *Filter) (*Item, error) {
	row, err := psql.Get(ctx, m.client, m.makeGetStatement(filter))
	if err != nil {
		return nil, err
	}

	// return the Item
	return m.scanRow(ctx, row)
}

func (m *Model) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(m.makeStatement(), filter))
	if err != nil {
		return nil, errors.AddCode(err, "501103")
	}
	defer rows.Close()
	// iterate over the result set
	items := make(map[uuid.UUID]Item)
	ids := make([]uuid.UUID, 0)
	for rows.Next() {
		item, err := m.scanRow(ctx, rows)
		if err != nil {
			return nil, errors.AddCode(err, "501104")
		}
		items[item.Id] = *item
		ids = append(ids, item.Id)
	}

	// count the number of rows
	if filter.DataConfig.IsCount != nil && *filter.DataConfig.IsCount == true {
		rows, err = psql.List(ctx, m.client, m.makeCountStatementByFilter(filter))
		if err != nil {
			return nil, errors.AddCode(err, "501105")
		}

		defer rows.Close()
		for rows.Next() {
			filter.DataConfig.Count, err = m.scanCountRow(ctx, rows)
			if err != nil {
				return nil, errors.AddCode(err, "501106")
			}
		}
	}

	return &items, nil
}

func (m *Model) Ids(ctx context.Context, filter *Filter) (*[]uuid.UUID, error) {
	rows, err := psql.List(ctx, m.client, m.makeStatementByFilter(m.makeIdsStatement(), filter))
	if err != nil {
		return nil, errors.AddCode(err, "401074")
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
			return nil, errors.AddCode(err, "145742")
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
		return nil, errors.AddCode(err, "235885")
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
		return errors.AddCode(err, "873950")
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
		return errors.AddCode(err, "942433")
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
		return errors.AddCode(err, "192133")
	}

	return nil
}

func (m *Model) TableIndexCount(ctx context.Context) (*uint64, error) {
	count, err := psql.TableIndexCount(
		ctx,
		m.client,
		m.qb.Select("n_tup_upd").From("pg_stat_user_tables").Where("relname = ?", m.table),
	)

	if err != nil {
		return nil, errors.AddCode(err, "792552")
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
		return nil, errors.AddCode(err, "936425")
	}

	return order, nil
}

func (m *Model) UpdatedAt(ctx context.Context, u *uuid.UUID) (*time.Time, error) {
	return psql.UpdatedAt(
		ctx,
		m.client,
		m.qb.Select(m.mapFieldToDBColumn("UpdatedAt")).From(m.table).Where(fmt.Sprintf("%s = ?", m.mapFieldToDBColumn("Id")), u),
	)
}
