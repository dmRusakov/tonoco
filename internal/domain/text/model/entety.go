package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/google/uuid"
	"time"
)

func (m *Model) mapFieldToDBColumn(field string) string {
	return m.dbField[field]
}

func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("Language"),
		m.mapFieldToDBColumn("Source"),
		m.mapFieldToDBColumn("SourceId"),
		m.mapFieldToDBColumn("Text"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table)
}

func (m *Model) makeIdsStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.mapFieldToDBColumn("Id"),
	).From(m.table)
}

func (m *Model) filterToStatement(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
	// id
	if filter.Ids != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Id")+" = ?", (*filter.Ids)[0])
	}

	// language
	if filter.Language != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Language")+" = ?", *filter.Language)
	}

	// Source
	if filter.Source != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Source")+" IN (?)", *filter.Source)
	}

	// SourceId
	if filter.SourceId != nil {
		statement = statement.Where(m.mapFieldToDBColumn("SourceId")+" IN (?)", *filter.SourceId)
	}

	// active
	if filter.Active != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Active")+" = ?", *filter.Active)
	}

	// search
	if filter.Search != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Text")+" ILIKE ?", "%"+*filter.Search+"%")
	}

	// Done
	return statement
}

func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if filter.Ids != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Id")+" = ?", (*filter.Ids)[0])
	}

	// language
	if filter.Language != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Language")+" = ?", *filter.Language)
	}

	// Source
	if filter.Source != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Source")+" IN (?)", *filter.Source)
	}

	// SourceId
	if filter.SourceId != nil {
		statement = statement.Where(m.mapFieldToDBColumn("SourceId")+" IN (?)", *filter.SourceId)
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Active")+" = ?", *filter.Active)
	}

	return statement
}

func (m *Model) makeStatementByFilter(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
	entity.CheckDataPagination(filter.DataPagination)
	statement = m.filterToStatement(statement, filter)
	return statement.OrderBy(m.mapFieldToDBColumn(*filter.DataPagination.OrderBy) + " " + *filter.DataPagination.OrderDir).
		Offset((*filter.DataPagination.Page - 1) * *filter.DataPagination.PerPage).Limit(*filter.DataPagination.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	statement := m.qb.Select("COUNT(*) as count").From(m.table + " p")
	return m.filterToStatement(statement, filter)
}

func (m *Model) scanRow(ctx context.Context, row sq.RowScanner) (*Item, error) {
	var (
		item                                                       = Item{}
		id, language, source, sourceId, text, createdBy, updatedBy sql.NullString
		active                                                     sql.NullBool
		createdAt, updatedAt                                       sql.NullTime
	)

	err := row.Scan(&id, &language, &source, &sourceId, &text, &active, &createdAt, &createdBy, &updatedAt, &updatedBy)
	if err != nil {
		tracing.Error(ctx, psql.ErrScan(psql.ParsePgError(err)))
		return nil, errors.AddCode(err, "854475")
	}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if language.Valid {
		item.Language = language.String
	}

	if source.Valid {
		item.Source = source.String
	}

	if sourceId.Valid {
		item.SourceId = uuid.MustParse(sourceId.String)
	}

	if text.Valid {
		item.Text = text.String
	}

	if active.Valid {
		item.Active = active.Bool
	}

	if createdAt.Valid {
		item.CreatedAt = createdAt.Time
	}

	if createdBy.Valid {
		item.CreatedBy = uuid.MustParse(createdBy.String)
	}

	if updatedAt.Valid {
		item.UpdatedAt = updatedAt.Time
	}

	if updatedBy.Valid {
		item.UpdatedBy = uuid.MustParse(updatedBy.String)
	}

	return &item, nil
}

func (m *Model) scanIdRow(ctx context.Context, row sq.RowScanner) (*uuid.UUID, error) {
	var id sql.NullString

	err := row.Scan(&id)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "608735")
	}

	return pointer.StringToUUID(id.String), nil
}

func (m *Model) scanCountRow(ctx context.Context, row sq.RowScanner) (*uint64, error) {
	var count uint64

	err := row.Scan(&count)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "547008")
	}

	return &count, nil
}

func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *uuid.UUID) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if Id is not set, generate a new UUID
	if item.Id == uuid.Nil {
		item.Id = uuid.New()
	}

	// set Id to context
	ctx = context.WithValue(ctx, "itemId", item.Id)

	insertItem := m.qb.Insert(m.table).
		Columns(
			m.mapFieldToDBColumn("Id"),
			m.mapFieldToDBColumn("Language"),
			m.mapFieldToDBColumn("Source"),
			m.mapFieldToDBColumn("SourceId"),
			m.mapFieldToDBColumn("Text"),
			m.mapFieldToDBColumn("Active"),
			m.mapFieldToDBColumn("CreatedAt"),
			m.mapFieldToDBColumn("CreatedBy"),
			m.mapFieldToDBColumn("UpdatedAt"),
			m.mapFieldToDBColumn("UpdatedBy"),
		).
		Values(
			item.Id,
			item.Language,
			item.Source,
			item.SourceId,
			item.Text,
			item.Active,
			time.Now(),
			uuid.MustParse(by),
			time.Now(),
			uuid.MustParse(by),
		)

	return &insertItem, &item.Id
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		SetMap(map[string]interface{}{
			m.mapFieldToDBColumn("Language"):  item.Language,
			m.mapFieldToDBColumn("Source"):    item.Source,
			m.mapFieldToDBColumn("SourceId"):  item.SourceId,
			m.mapFieldToDBColumn("Text"):      item.Text,
			m.mapFieldToDBColumn("Active"):    item.Active,
			m.mapFieldToDBColumn("UpdatedAt"): time.Now(),
			m.mapFieldToDBColumn("UpdatedBy"): uuid.MustParse(by),
		}).Where(m.mapFieldToDBColumn("Id")+" = ?", item.Id)

}

func (m *Model) makePatchStatement(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		statement = statement.Set(m.mapFieldToDBColumn(field), value)
	}

	return statement.Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}
