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
	"reflect"
)

func (m *Model) mapFieldToDBColumn(field string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	// check if field is in the cash
	if dbField, ok := m.dbFieldCash[field]; ok {
		return dbField
	}

	// get field from struct
	typeOf := reflect.TypeOf(Item{})
	byName, _ := typeOf.FieldByName(field)
	dbField := byName.Tag.Get("db")

	// set field to cash
	m.dbFieldCash[field] = dbField

	// done
	return dbField
}

func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("ProductId"),
		m.mapFieldToDBColumn("ImageId"),
		m.mapFieldToDBColumn("Type"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table + " p")
}

func (m *Model) fillInFilter(statement sq.SelectBuilder, filter *Filter) sq.SelectBuilder {
	if filter.Ids != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
	}

	if filter.ProductIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("ProductId"): *filter.ProductIds})
	}

	if filter.ImageIds != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("ImageId"): *filter.ImageIds})
	}

	if filter.Type != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Type"): *filter.Type})
	}

	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Type")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	return m.fillInFilter(m.makeStatement(), filter)
}

func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	if filter.OrderBy == nil {
		filter.OrderBy = entity.StringPtr("SortOrder")
	}

	if filter.OrderDir == nil {
		filter.OrderDir = entity.StringPtr("ASC")
	}

	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = pointer.Uint64Ptr(999999999999999999)
		} else {
			filter.PerPage = pointer.Uint64Ptr(10)
		}
	}

	if filter.Page == nil {
		filter.Page = pointer.Uint64Ptr(1)
	}

	statement := m.makeGetStatement(filter)

	return statement.OrderBy(m.mapFieldToDBColumn(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	return m.fillInFilter(m.qb.Select("COUNT(*)").From(m.table), filter)
}

func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, productId, imageId, typeField, createdBy, updatedBy sql.NullString
	var sortOrder sql.NullInt64
	var createdAt, updatedAt sql.NullTime

	err := rows.Scan(
		&id,
		&productId,
		&imageId,
		&typeField,
		&sortOrder,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "853456")
	}

	var item = Item{}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}

	if productId.Valid {
		item.ProductId = uuid.MustParse(productId.String)
	}

	if imageId.Valid {
		item.ImageId = uuid.MustParse(imageId.String)
	}

	if typeField.Valid {
		item.Type = typeField.String
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

func (m *Model) scanCountRow(ctx context.Context, rows sq.RowScanner) (*uint64, error) {
	var count uint64

	err := rows.Scan(&count)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	return &count, nil
}

func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *uuid.UUID) {
	by := ctx.Value("user_id").(string)

	if item.Id == uuid.Nil {
		item.Id = uuid.New()
	}

	ctx = context.WithValue(ctx, "itemId", item.Id)

	insertItem := m.qb.Insert(m.table).Columns(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("ProductId"),
		m.mapFieldToDBColumn("ImageId"),
		m.mapFieldToDBColumn("Type"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.ProductId,
		item.ImageId,
		item.Type,
		item.SortOrder,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	return &insertItem, ctx.Value("itemId").(*uuid.UUID)
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.mapFieldToDBColumn("ProductId"), item.ProductId).
		Set(m.mapFieldToDBColumn("ImageId"), item.ImageId).
		Set(m.mapFieldToDBColumn("Type"), item.Type).
		Set(m.mapFieldToDBColumn("SortOrder"), item.SortOrder).
		Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").
		Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}

func (m *Model) makePatchStatement(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) sq.UpdateBuilder {
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		field = m.mapFieldToDBColumn(field)
		statement = statement.Set(field, value)
	}

	return statement.Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").Set(m.mapFieldToDBColumn("UpdatedBy"), by)
}
