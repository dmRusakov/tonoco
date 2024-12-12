package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
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
		m.mapFieldToDBColumn("Name"),
		m.mapFieldToDBColumn("Url"),
		m.mapFieldToDBColumn("ShortDescription"),
		m.mapFieldToDBColumn("Description"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Type"),
		m.mapFieldToDBColumn("Prefix"),
		m.mapFieldToDBColumn("Suffix"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).From(m.table + " p")
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

func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = pointer.StringToPtr("SortOrder")
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = pointer.StringToPtr("ASC")
	}

	// PerPage
	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = pointer.UintTo64Ptr(999999999999999999)
		} else {
			filter.PerPage = pointer.UintTo64Ptr(10)
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = pointer.UintTo64Ptr(1)
	}

	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.Ids != nil {
		countIds := len(*filter.Ids)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countIds) {
			*filter.PerPage = uint64(countIds)
		}
	}

	// Urls
	if filter.Urls != nil {
		countUrls := len(*filter.Urls)

		if countUrls > 0 {
			statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Url"): *filter.Urls})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countUrls) {
			*filter.PerPage = uint64(countUrls)
		}
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Active"): *filter.Active})
	}

	// Type
	if filter.Type != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Type"): *filter.Type})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("ShortDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Description")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.mapFieldToDBColumn(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	// Build query
	statement := m.qb.Select("COUNT(*)").From(m.table + " p")

	// Ids
	if filter.Ids != nil {
		countIds := len(*filter.Ids)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Id"): *filter.Ids})
		}
	}

	// Urls
	if filter.Urls != nil {
		countUrls := len(*filter.Urls)

		if countUrls > 0 {
			statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Url"): *filter.Urls})
		}
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Active"): *filter.Active})
	}

	// Type
	if filter.Type != nil {
		statement = statement.Where(sq.Eq{m.mapFieldToDBColumn("Type"): *filter.Type})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("ShortDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Description")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

func (m *Model) scanRow(ctx context.Context, row sq.RowScanner) (*Item, error) {
	var id, name, url, shortDescription, description, typeField, prefix, suffix, createdBy, updatedBy sql.NullString
	var active sql.NullBool
	var sortOrder sql.NullInt64
	var createdAt, updatedAt sql.NullTime

	err := row.Scan(
		&id,
		&name,
		&url,
		&shortDescription,
		&description,
		&active,
		&sortOrder,
		&typeField,
		&prefix,
		&suffix,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "167664")
	}

	var item = Item{}

	if id.Valid {
		item.Id = uuid.MustParse(id.String)
	}
	if name.Valid {
		item.Name = name.String
	}
	if url.Valid {
		item.Url = url.String
	}
	if shortDescription.Valid {
		item.ShortDescription = shortDescription.String
	}
	if description.Valid {
		item.Description = description.String
	}
	if active.Valid {
		item.Active = active.Bool
	}
	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
	}
	if typeField.Valid {
		item.Type = typeField.String
	}
	if prefix.Valid {
		item.Prefix = prefix.String
	}
	if suffix.Valid {
		item.Suffix = suffix.String
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

func (m *Model) scanCountRow(ctx context.Context, row sq.RowScanner) (*uint64, error) {
	var count uint64

	err := row.Scan(&count)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
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

	insertItem := m.qb.Insert(m.table).Columns(
		m.mapFieldToDBColumn("Id"),
		m.mapFieldToDBColumn("Name"),
		m.mapFieldToDBColumn("Url"),
		m.mapFieldToDBColumn("ShortDescription"),
		m.mapFieldToDBColumn("Description"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Type"),
		m.mapFieldToDBColumn("Prefix"),
		m.mapFieldToDBColumn("Suffix"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.Name,
		item.Url,
		item.ShortDescription,
		item.Description,
		item.Active,
		item.SortOrder,
		item.Type,
		item.Prefix,
		item.Suffix,
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
		Set(m.mapFieldToDBColumn("Name"), item.Name).
		Set(m.mapFieldToDBColumn("Url"), item.Url).
		Set(m.mapFieldToDBColumn("ShortDescription"), item.ShortDescription).
		Set(m.mapFieldToDBColumn("Description"), item.Description).
		Set(m.mapFieldToDBColumn("Active"), item.Active).
		Set(m.mapFieldToDBColumn("SortOrder"), item.SortOrder).
		Set(m.mapFieldToDBColumn("Type"), item.Type).
		Set(m.mapFieldToDBColumn("Prefix"), item.Prefix).
		Set(m.mapFieldToDBColumn("Suffix"), item.Suffix).
		Set(m.mapFieldToDBColumn("UpdatedAt"), "NOW()").
		Set(m.mapFieldToDBColumn("UpdatedBy"), by)
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
