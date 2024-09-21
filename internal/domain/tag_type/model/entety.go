package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"reflect"
)

// fieldMap
func (m *Model) fieldMap(field string) string {
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

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		m.fieldMap("ID"),
		m.fieldMap("Name"),
		m.fieldMap("Url"),
		m.fieldMap("ShortDescription"),
		m.fieldMap("Description"),
		m.fieldMap("Required"),
		m.fieldMap("Active"),
		m.fieldMap("Prime"),
		m.fieldMap("ListItem"),
		m.fieldMap("Filter"),
		m.fieldMap("SortOrder"),
		m.fieldMap("Type"),
		m.fieldMap("Prefix"),
		m.fieldMap("Suffix"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
	).From(m.table + " p")
}

// make Get statement
func (m *Model) makeGetStatement(filter *Filter) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if filter.Ids != nil {
		statement = statement.Where(m.fieldMap("ID")+" = ?", (*filter.Ids)[0])
	}

	// url
	if filter.Urls != nil {
		statement = statement.Where(m.fieldMap("Url")+" = ?", (*filter.Urls)[0])
	}

	return statement
}

// makeStatementByFilter
func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = entity.StringPtr("SortOrder")
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = entity.StringPtr("ASC")
	}

	// PerPage
	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = entity.Uint64Ptr(999999999999999999)
		} else {
			filter.PerPage = entity.Uint64Ptr(10)
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = entity.Uint64Ptr(1)
	}

	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.Ids != nil {
		countIds := len(*filter.Ids)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("ID"): *filter.Ids})
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
			statement = statement.Where(sq.Eq{m.fieldMap("Url"): *filter.Urls})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countUrls) {
			*filter.PerPage = uint64(countUrls)
		}
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Active"): *filter.Active})
	}

	// Prime
	if filter.Prime != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Prime"): *filter.Prime})
	}

	// ListItem
	if filter.ListItem != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("ListItem"): *filter.ListItem})
	}

	// Filter
	if filter.Filter != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Filter"): *filter.Filter})
	}

	// Type
	if filter.Type != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Type"): *filter.Type})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.fieldMap("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("ShortDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Description")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.fieldMap(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

// makeCountStatementByFilter - make count statement by filter for pagination
func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
	// Build query
	statement := m.qb.Select("COUNT(*)").From(m.table + " p")

	// Ids
	if filter.Ids != nil {
		countIds := len(*filter.Ids)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("ID"): *filter.Ids})
		}
	}

	// Urls
	if filter.Urls != nil {
		countUrls := len(*filter.Urls)

		if countUrls > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("Url"): *filter.Urls})
		}
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Active"): *filter.Active})
	}

	// Prime
	if filter.Prime != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Prime"): *filter.Prime})
	}

	// ListItem
	if filter.ListItem != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("ListItem"): *filter.ListItem})
	}

	// Filter
	if filter.Filter != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Filter"): *filter.Filter})
	}

	// Type
	if filter.Type != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Type"): *filter.Type})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.fieldMap("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("ShortDescription")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Description")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, name, url, shortDescription, description, typeField, prefix, suffix, createdBy, updatedBy sql.NullString
	var required, active, prime, listItem, filter sql.NullBool
	var sortOrder sql.NullInt64
	var createdAt, updatedAt sql.NullTime

	err := rows.Scan(
		&id,
		&name,
		&url,
		&shortDescription,
		&description,
		&required,
		&active,
		&prime,
		&listItem,
		&filter,
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
		item.ID = uuid.MustParse(id.String)
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
	if required.Valid {
		item.Required = required.Bool
	}
	if active.Valid {
		item.Active = active.Bool
	}
	if prime.Valid {
		item.Prime = prime.Bool
	}
	if listItem.Valid {
		item.ListItem = listItem.Bool
	}
	if filter.Valid {
		item.Filter = filter.Bool
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

// makeInsertStatement
func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *uuid.UUID) {

	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if ID is not set, generate a new UUID
	if item.ID == uuid.Nil {
		item.ID = uuid.New()
	}

	// set ID to context
	ctx = context.WithValue(ctx, "itemId", item.ID)

	insertItem := m.qb.Insert(m.table).Columns(
		m.fieldMap("ID"),
		m.fieldMap("Name"),
		m.fieldMap("Url"),
		m.fieldMap("ShortDescription"),
		m.fieldMap("Description"),
		m.fieldMap("Required"),
		m.fieldMap("Active"),
		m.fieldMap("Prime"),
		m.fieldMap("ListItem"),
		m.fieldMap("Filter"),
		m.fieldMap("SortOrder"),
		m.fieldMap("Type"),
		m.fieldMap("Prefix"),
		m.fieldMap("Suffix"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
	).Values(
		item.ID,
		item.Name,
		item.Url,
		item.ShortDescription,
		item.Description,
		item.Required,
		item.Active,
		item.Prime,
		item.ListItem,
		item.Filter,
		item.SortOrder,
		item.Type,
		item.Prefix,
		item.Suffix,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	return &insertItem, &item.ID
}

// makeUpdateStatement
func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.fieldMap("Name"), item.Name).
		Set(m.fieldMap("Url"), item.Url).
		Set(m.fieldMap("ShortDescription"), item.ShortDescription).
		Set(m.fieldMap("Description"), item.Description).
		Set(m.fieldMap("Required"), item.Required).
		Set(m.fieldMap("Active"), item.Active).
		Set(m.fieldMap("Prime"), item.Prime).
		Set(m.fieldMap("ListItem"), item.ListItem).
		Set(m.fieldMap("Filter"), item.Filter).
		Set(m.fieldMap("SortOrder"), item.SortOrder).
		Set(m.fieldMap("Type"), item.Type).
		Set(m.fieldMap("Prefix"), item.Prefix).
		Set(m.fieldMap("Suffix"), item.Suffix).
		Set(m.fieldMap("UpdatedAt"), "NOW()").
		Set(m.fieldMap("UpdatedBy"), by)
}

// makePatchStatement
func (m *Model) makePatchStatement(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		field = m.fieldMap(field)
		statement = statement.Set(field, value)
	}

	return statement.Set(m.fieldMap("UpdatedAt"), "NOW()").Set(m.fieldMap("UpdatedBy"), by)
}
