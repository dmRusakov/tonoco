package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
)

type Item = entity.Store
type Filter = entity.StoreFilter

// fieldMap
var fieldMap = map[string]string{
	"ID":           "id",
	"Name":         "name",
	"Abbreviation": "abbreviation",
	"Config":       "config",
	"SortOrder":    "sort_order",
	"Active":       "active",
	"AddressLine1": "address_line1",
	"AddressLine2": "address_line2",
	"City":         "city",
	"State":        "state",
	"ZipCode":      "zip_code",
	"Country":      "country",
	"WebSite":      "web_site",
	"Phone":        "phone",
	"Email":        "email",
	"CreatedAt":    "created_at",
	"CreatedBy":    "created_by",
	"UpdatedAt":    "updated_at",
	"UpdatedBy":    "updated_by",
}

// makeStatement
func (m *Model) makeStatement() sq.SelectBuilder {
	return m.qb.Select(
		fieldMap["ID"],
		fieldMap["Name"],
		fieldMap["Abbreviation"],
		fieldMap["Config"],
		fieldMap["SortOrder"],
		fieldMap["Active"],
		fieldMap["AddressLine1"],
		fieldMap["AddressLine2"],
		fieldMap["City"],
		fieldMap["State"],
		fieldMap["ZipCode"],
		fieldMap["Country"],
		fieldMap["WebSite"],
		fieldMap["Phone"],
		fieldMap["Email"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).From(m.table + " p")
}

// make Get statement
func (m *Model) makeGetStatement(
	id *string,
	abbreviation *string,
) sq.SelectBuilder {
	// build query
	statement := m.makeStatement()

	// id
	if id != nil {
		statement = statement.Where(fieldMap["ID"]+" = ?", *id)
	}

	// abbreviation
	if abbreviation != nil {
		statement = statement.Where(fieldMap["Abbreviation"]+" = ?", *abbreviation)
	}

	return statement
}

// makeStatementByFilter
func (m *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
	// Build query
	statement := m.makeStatement()

	// Ids
	if filter.IDs != nil {
		countIds := len(*filter.IDs)

		if countIds > 0 {
			statement = statement.Where(sq.Eq{fieldMap["ID"]: *filter.IDs})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countIds) {
			*filter.PerPage = uint64(countIds)
		}
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{fieldMap["Active"]: *filter.Active})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+fieldMap["Name"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["Abbreviations"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["AddressLine1"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["AddressLine2"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["City"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["State"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	// OrderBy
	if filter.OrderBy == nil {
		filter.OrderBy = new(string)
		*filter.OrderBy = "SortOrder"
	}

	// OrderDir
	if filter.OrderDir == nil {
		filter.OrderDir = new(string)
		*filter.OrderDir = "ASC"
	}

	// PerPage
	if filter.PerPage == nil {
		filter.PerPage = new(uint64)
		if filter.Page == nil {
			*filter.PerPage = 999999999999999999
		} else {
			*filter.PerPage = 10
		}
	}

	// Page
	if filter.Page == nil {
		filter.Page = new(uint64)
		*filter.Page = 1
	}

	// Add OrderBy, OrderDir, Page, Limit and return
	return statement.OrderBy(fieldMap[*filter.OrderBy] + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)
}

// scanOneRow
func (m *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = &Item{}
	err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.Abbreviation,
		&item.Config,
		&item.SortOrder,
		&item.Active,
		&item.AddressLine1,
		&item.AddressLine2,
		&item.City,
		&item.State,
		&item.ZipCode,
		&item.Country,
		&item.WebSite,
		&item.Phone,
		&item.Email,
		&item.CreatedAt,
		&item.CreatedBy,
		&item.UpdatedAt,
		&item.UpdatedBy,
	)

	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	return item, nil
}

// makeInsertStatement
func (m *Model) makeInsertStatement(ctx context.Context, item *Item) (*sq.InsertBuilder, *string) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if ID is not set, generate a new UUID
	if item.ID == "" {
		item.ID = uuid.New().String()
	}

	// set ID to context
	ctx = context.WithValue(ctx, "itemId", item.ID)

	insertItem := m.qb.Insert(m.table).Columns(
		fieldMap["ID"],
		fieldMap["Name"],
		fieldMap["Abbreviation"],
		fieldMap["Config"],
		fieldMap["SortOrder"],
		fieldMap["Active"],
		fieldMap["AddressLine1"],
		fieldMap["AddressLine2"],
		fieldMap["City"],
		fieldMap["State"],
		fieldMap["ZipCode"],
		fieldMap["Country"],
		fieldMap["WebSite"],
		fieldMap["Phone"],
		fieldMap["Email"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).Values(
		item.ID,
		item.Name,
		item.Abbreviation,
		item.Config,
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

	// get itemId from context
	return &insertItem, &item.ID
}

// makeUpdateStatement
func (m *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return m.qb.Update(m.table).
		Set(fieldMap["Name"], item.Name).
		Set(fieldMap["Abbreviation"], item.Abbreviation).
		Set(fieldMap["Config"], item.Config).
		Set(fieldMap["SortOrder"], item.SortOrder).
		Set(fieldMap["Active"], item.Active).
		Set(fieldMap["AddressLine1"], item.AddressLine1).
		Set(fieldMap["AddressLine2"], item.AddressLine2).
		Set(fieldMap["City"], item.City).
		Set(fieldMap["State"], item.State).
		Set(fieldMap["ZipCode"], item.ZipCode).
		Set(fieldMap["Country"], item.Country).
		Set(fieldMap["WebSite"], item.WebSite).
		Set(fieldMap["Phone"], item.Phone).
		Set(fieldMap["Email"], item.Email).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by)
}

// makePatchStatement
func (m *Model) makePatchStatement(ctx context.Context, id *string, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := m.qb.Update(m.table).Where("id = ?", id)

	for field, value := range *fields {
		field = fieldMap[field]
		statement = statement.Set(field, value)
	}

	return statement.Set(fieldMap["UpdatedAt"], "NOW()").Set(fieldMap["UpdatedBy"], by)
}
