package model

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
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
		m.fieldMap("Id"),
		m.fieldMap("Name"),
		m.fieldMap("Url"),
		m.fieldMap("Abbreviation"),
		m.fieldMap("SortOrder"),
		m.fieldMap("Active"),
		m.fieldMap("AddressLine1"),
		m.fieldMap("AddressLine2"),
		m.fieldMap("City"),
		m.fieldMap("State"),
		m.fieldMap("ZipCode"),
		m.fieldMap("Country"),
		m.fieldMap("WebSite"),
		m.fieldMap("Phone"),
		m.fieldMap("Email"),
		m.fieldMap("CurrencyUrl"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
	).From(m.table + " p")
}

// make Get statement
func (m *Model) makeGetStatement(
	filter *Filter,
) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if filter.Ids != nil {
		statement = statement.Where(m.fieldMap("Id")+" = ?", (*filter.Ids)[0])
	}

	// url
	if filter.Urls != nil {
		statement = statement.Where(m.fieldMap("Url")+" = ?", (*filter.Urls)[0])
	}

	// abbreviation
	if filter.Abbreviations != nil {
		statement = statement.Where(m.fieldMap("Abbreviation")+" = ?", (*filter.Abbreviations)[0])

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
			statement = statement.Where(sq.Eq{m.fieldMap("Id"): *filter.Ids})
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

	// Abbreviations
	if filter.Abbreviations != nil {
		countAbbreviations := len(*filter.Abbreviations)

		if countAbbreviations > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("Abbreviation"): *filter.Abbreviations})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countAbbreviations) {
			*filter.PerPage = uint64(countAbbreviations)
		}

	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.fieldMap("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Abbreviations")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("AddressLine1")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("AddressLine2")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("City")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("State")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
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
			statement = statement.Where(sq.Eq{m.fieldMap("Id"): *filter.Ids})
		}
	}

	// Urls
	if filter.Urls != nil {
		countUrls := len(*filter.Urls)

		if countUrls > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("Url"): *filter.Urls})
		}
	}

	// Abbreviations
	if filter.Abbreviations != nil {
		countAbbreviations := len(*filter.Abbreviations)

		if countAbbreviations > 0 {
			statement = statement.Where(sq.Eq{m.fieldMap("Abbreviation"): *filter.Abbreviations})
		}
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{m.fieldMap("Active"): *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+m.fieldMap("Name")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Url")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("Abbreviations")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("AddressLine1")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("AddressLine2")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("City")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+m.fieldMap("State")+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var id, url, name, abbreviation, addressLine1, addressLine2, city, state, zipCode, country, webSite, phone, email, currencyUrl, createdBy, updatedBy sql.NullString
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
		&currencyUrl,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	var item = Item{}

	if id.Valid {
		item.Id = id.String
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

	if currencyUrl.Valid {
		item.CurrencyUrl = currencyUrl.String
	}

	if createdAt.Valid {
		item.CreatedAt = createdAt.Time
	}

	if createdBy.Valid {
		item.CreatedBy = createdBy.String
	}

	if updatedAt.Valid {
		item.UpdatedAt = updatedAt.Time
	}

	if updatedBy.Valid {
		item.UpdatedBy = updatedBy.String
	}

	return &item, nil
}

// makeInsertStatement
func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *string) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if Id is not set, generate a new UUID
	if item.Id == "" {
		item.Id = uuid.New().String()
	}

	// set Id to context
	ctx = context.WithValue(ctx, "itemId", item.Id)

	insertItem := m.qb.Insert(m.table).Columns(
		m.fieldMap("Id"),
		m.fieldMap("Name"),
		m.fieldMap("Url"),
		m.fieldMap("Abbreviation"),
		m.fieldMap("SortOrder"),
		m.fieldMap("Active"),
		m.fieldMap("AddressLine1"),
		m.fieldMap("AddressLine2"),
		m.fieldMap("City"),
		m.fieldMap("State"),
		m.fieldMap("ZipCode"),
		m.fieldMap("Country"),
		m.fieldMap("WebSite"),
		m.fieldMap("Phone"),
		m.fieldMap("Email"),
		m.fieldMap("CurrencyUrl"),
		m.fieldMap("CreatedAt"),
		m.fieldMap("CreatedBy"),
		m.fieldMap("UpdatedAt"),
		m.fieldMap("UpdatedBy"),
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
		item.CurrencyUrl,
		"NOW()",
		by,
		"NOW()",
		by,
	)

	// get itemId from context
	return &insertItem, &item.Id
}

// makeUpdateStatement
func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(m.fieldMap("Name"), item.Name).
		Set(m.fieldMap("Url"), item.Url).
		Set(m.fieldMap("Abbreviation"), item.Abbreviation).
		Set(m.fieldMap("SortOrder"), item.SortOrder).
		Set(m.fieldMap("Active"), item.Active).
		Set(m.fieldMap("AddressLine1"), item.AddressLine1).
		Set(m.fieldMap("AddressLine2"), item.AddressLine2).
		Set(m.fieldMap("City"), item.City).
		Set(m.fieldMap("State"), item.State).
		Set(m.fieldMap("ZipCode"), item.ZipCode).
		Set(m.fieldMap("Country"), item.Country).
		Set(m.fieldMap("WebSite"), item.WebSite).
		Set(m.fieldMap("Phone"), item.Phone).
		Set(m.fieldMap("Email"), item.Email).
		Set(m.fieldMap("CurrencyUrl"), item.CurrencyUrl).
		Set(m.fieldMap("UpdatedAt"), "NOW()").
		Set(m.fieldMap("UpdatedBy"), by)
}

// makePatchStatement
func (m *Model) makePatchStatement(ctx context.Context, id *string, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		field = m.fieldMap(field)
		statement = statement.Set(field, value)
	}

	return statement.Set(m.fieldMap("UpdatedAt"), "NOW()").Set(m.fieldMap("UpdatedBy"), by)
}
