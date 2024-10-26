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
		m.mapFieldToDBColumn("Abbreviation"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("AddressLine1"),
		m.mapFieldToDBColumn("AddressLine2"),
		m.mapFieldToDBColumn("City"),
		m.mapFieldToDBColumn("State"),
		m.mapFieldToDBColumn("ZipCode"),
		m.mapFieldToDBColumn("Country"),
		m.mapFieldToDBColumn("WebSite"),
		m.mapFieldToDBColumn("Phone"),
		m.mapFieldToDBColumn("Email"),
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
		filter.OrderBy = pointer.StringPtr("SortOrder")
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = pointer.StringPtr("ASC")
	}

	// PerPage
	if filter.PerPage == nil {
		if filter.Page == nil {
			filter.PerPage = pointer.Uint64Ptr(999999999999999999)
		} else {
			filter.PerPage = pointer.Uint64Ptr(10)
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = pointer.Uint64Ptr(1)
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

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Abbreviations")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("AddressLine1")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("AddressLine2")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("City")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("State")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(m.mapFieldToDBColumn(*filter.OrderBy) + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

func (m *Model) makeCountStatementByFilter(filter *Filter) sq.SelectBuilder {
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

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("Abbreviations")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("AddressLine1")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("AddressLine2")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("City")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.mapFieldToDBColumn("State")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, name, url, abbreviation, addressLine1, addressLine2, city, state, zipCode, country, webSite, phone, email, createdBy, updatedBy sql.NullString
	var sortOrder sql.NullInt64
	var active sql.NullBool
	var createdAt, updatedAt sql.NullTime

	err := rows.Scan(
		&id,
		&name,
		&url,
		&abbreviation,
		&sortOrder,
		&active,
		&addressLine1,
		&addressLine2,
		&city,
		&state,
		&zipCode,
		&country,
		&webSite,
		&phone,
		&email,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, errors.AddCode(err, "740849")
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

	if abbreviation.Valid {
		item.Abbreviation = abbreviation.String
	}

	if sortOrder.Valid {
		item.SortOrder = uint64(sortOrder.Int64)
	}

	if active.Valid {
		item.Active = active.Bool
	}

	if addressLine1.Valid {
		item.AddressLine1 = addressLine1.String
	}

	if addressLine2.Valid {
		item.AddressLine2 = addressLine2.String
	}

	if city.Valid {
		item.City = city.String
	}

	if state.Valid {
		item.State = state.String
	}

	if zipCode.Valid {
		item.ZipCode = zipCode.String
	}

	if country.Valid {
		item.Country = country.String
	}

	if webSite.Valid {
		item.WebSite = webSite.String
	}

	if phone.Valid {
		item.Phone = phone.String
	}

	if email.Valid {
		item.Email = email.String
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
		m.mapFieldToDBColumn("Abbreviation"),
		m.mapFieldToDBColumn("SortOrder"),
		m.mapFieldToDBColumn("Active"),
		m.mapFieldToDBColumn("AddressLine1"),
		m.mapFieldToDBColumn("AddressLine2"),
		m.mapFieldToDBColumn("City"),
		m.mapFieldToDBColumn("State"),
		m.mapFieldToDBColumn("ZipCode"),
		m.mapFieldToDBColumn("Country"),
		m.mapFieldToDBColumn("WebSite"),
		m.mapFieldToDBColumn("Phone"),
		m.mapFieldToDBColumn("Email"),
		m.mapFieldToDBColumn("CreatedAt"),
		m.mapFieldToDBColumn("CreatedBy"),
		m.mapFieldToDBColumn("UpdatedAt"),
		m.mapFieldToDBColumn("UpdatedBy"),
	).Values(
		item.Id,
		item.Name,
		item.Url,
		item.Abbreviation,
		item.SortOrder,
		item.Active,
		item.AddressLine1,
		item.AddressLine2,
		item.City,
		item.State,
		item.ZipCode,
		item.Country,
		item.WebSite,
		item.Phone,
		item.Email,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	// done
	return &insertItem, &item.Id
}

func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.mapFieldToDBColumn("Name"), item.Name).
		Set(m.mapFieldToDBColumn("Url"), item.Name).
		Set(m.mapFieldToDBColumn("Abbreviation"), item.Abbreviation).
		Set(m.mapFieldToDBColumn("SortOrder"), item.SortOrder).
		Set(m.mapFieldToDBColumn("Active"), item.Active).
		Set(m.mapFieldToDBColumn("AddressLine1"), item.AddressLine1).
		Set(m.mapFieldToDBColumn("AddressLine2"), item.AddressLine2).
		Set(m.mapFieldToDBColumn("City"), item.City).
		Set(m.mapFieldToDBColumn("State"), item.State).
		Set(m.mapFieldToDBColumn("ZipCode"), item.ZipCode).
		Set(m.mapFieldToDBColumn("Country"), item.Country).
		Set(m.mapFieldToDBColumn("WebSite"), item.WebSite).
		Set(m.mapFieldToDBColumn("Phone"), item.Phone).
		Set(m.mapFieldToDBColumn("Email"), item.Email).
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
