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
)

func (m *Model) mapFieldToDBColumn(field string) string {
	return m.dbField[field]
}

func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("TagTypeId"),
		m.mapFieldToDBColumn("Name"),
		m.mapFieldToDBColumn("Url"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table + " p")
}

func (m *Model) filterDTO(filter *Filter) {
	if filter == nil {
		filter = &Filter{}
	}

	// check DataConfig
	if filter.DataConfig == nil {
		filter.DataConfig = &entity.DataConfig{}
	}

	// check DataPagination
	if filter.DataPagination == nil {
		filter.DataPagination = &entity.DataPagination{}
	}

	entity.CheckDataPagination(filter.DataPagination)
}

func (m *Model) filterToStatement(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
	// Ids
	if filter.Ids != nil && len(*filter.Ids) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
	}

	// Urls
	if filter.Urls != nil && len(*filter.Urls) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Url"): *filter.Urls})
	}

	// TagTypeIds
	if filter.TagTypeIds != nil && len(*filter.TagTypeIds) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("TagTypeId"): *filter.TagTypeIds})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if filter.Ids != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Id")+" = ?", (*filter.Ids)[0])
	}

	// url
	if filter.Urls != nil {
		statement = statement.Where(m.mapFieldToDBColumn("Url")+" = ?", (*filter.Urls)[0])
	}

	return statement
}

func (m *Model) makeStatementByFilter(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
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
		item                                           = Item{}
		id, tagTypeId, name, url, createdBy, updatedBy sql.NullString
		active                                         sql.NullBool
		sortOrder                                      sql.NullInt64
		createdAt, updatedAt                           sql.NullTime
	)

	err := row.Scan(
		&id,
		&tagTypeId,
		&name,
		&url,
		&active,
		&sortOrder,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "391598")
	}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if tagTypeId.Valid {
		item.TagTypeId = uuid.MustParse(tagTypeId.String)
	}

	if name.Valid {
		item.Name = name.String
	}

	if url.Valid {
		item.Url = url.String
	}

	if active.Valid {
		item.Active = active.Bool
	}

	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
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
		return nil, errors.AddCode(err, "222149")
	}

	return pointer.StringToUUID(id.String), nil
}

func (m *Model) scanCountRow(ctx context.Context, row sq.RowScanner) (*uint64, error) {
	var count uint64

	err := row.Scan(&count)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "491602")
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

	// build query
	insertItem := m.qb.Insert(m.table).Columns(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("TagTypeId"),
		m.mapFieldToDBColumn("Name"),
		m.mapFieldToDBColumn("Url"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.TagTypeId,
		item.Name,
		item.Url,
		item.Active,
		item.SortOrder,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	return &insertItem, &item.Id
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.mapFieldToDBColumn("TagTypeId"), item.TagTypeId).
		Set(m.mapFieldToDBColumn("Name"), item.Name).
		Set(m.mapFieldToDBColumn("Url"), item.Url).
		Set(m.mapFieldToDBColumn("Active"), item.Active).
		Set(m.mapFieldToDBColumn("SortOrder"), item.SortOrder).
		Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").
		Set(m.mapFieldToDBColumn("UpdatedBy"), by).
		Where("id = ?", item.Id)
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
