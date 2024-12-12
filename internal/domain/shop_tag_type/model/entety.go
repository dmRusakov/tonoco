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
		m.mapFieldToDBColumn("ShopId"),
		m.mapFieldToDBColumn("TagTypeId"),
		m.mapFieldToDBColumn("Source"),
		m.mapFieldToDBColumn("SortOrder"),
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

	// ShopIds
	if filter.ShopIds != nil && len(*filter.ShopIds) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("ShopId"): *filter.ShopIds})
	}

	// TagTypeIds
	if filter.TagTypeIds != nil && len(*filter.TagTypeIds) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("TagTypeId"): *filter.TagTypeIds})
	}

	// Sources
	if filter.Sources != nil && len(*filter.Sources) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Source"): *filter.Sources})
	}

	// SortOrders
	if filter.SortOrders != nil && len(*filter.SortOrders) > 0 {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("SortOrder"): *filter.SortOrders})
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Active"): *filter.Active})
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
		item                                                = Item{}
		id, shopId, tagTypeId, source, createdBy, updatedBy sql.NullString
		sortOrder                                           sql.NullInt64
		active                                              sql.NullBool
		createdAt, updatedAt                                sql.NullTime
	)

	err := row.Scan(&id, &shopId, &tagTypeId, &source, &sortOrder, &active, &createdAt, &createdBy, &updatedAt, &updatedBy)
	if err != nil {
		tracing.Error(ctx, psql.ErrScan(psql.ParsePgError(err)))
		return nil, errors.AddCode(err, "396147")
	}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if shopId.Valid {
		item.ShopId = uuid.MustParse(shopId.String)
	}

	if tagTypeId.Valid {
		item.TagTypeId = uuid.MustParse(tagTypeId.String)
	}

	if source.Valid {
		item.Source = source.String
	}

	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
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

	// set Id to context
	ctx = context.WithValue(ctx, "itemId", item.Id)

	// build query
	insertItem := m.qb.Insert(m.table).
		Columns(
			m.mapFieldToDBColumn("Id"),
			m.mapFieldToDBColumn("ShopId"),
			m.mapFieldToDBColumn("TagTypeId"),
			m.mapFieldToDBColumn("Source"),
			m.mapFieldToDBColumn("SortOrder"),
			m.mapFieldToDBColumn("Active"),
			m.mapFieldToDBColumn("CreatedAt"),
			m.mapFieldToDBColumn("CreatedBy"),
			m.mapFieldToDBColumn("UpdatedAt"),
			m.mapFieldToDBColumn("UpdatedBy"),
		).
		Values(
			item.Id,
			item.ShopId,
			item.TagTypeId,
			item.Source,
			item.SortOrder,
			item.Active,
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
		SetMap(map[string]interface{}{
			m.mapFieldToDBColumn("ShopId"):    item.ShopId,
			m.mapFieldToDBColumn("TagTypeId"): item.TagTypeId,
			m.mapFieldToDBColumn("Source"):    item.Source,
			m.mapFieldToDBColumn("SortOrder"): item.SortOrder,
			m.mapFieldToDBColumn("Active"):    item.Active,
			m.mapFieldToDBColumn("UpdatedAt"): "NOW()",
			m.mapFieldToDBColumn("UpdatedBy"): by,
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
