package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
)

type Item = entity.ProductCategory
type Filter = entity.ProductCategoryFilter

// fieldMap
var fieldMap = map[string]string{
	"ID":               "id",
	"Name":             "name",
	"Url":              "url",
	"ShortDescription": "short_description",
	"Description":      "description",
	"SortOrder":        "sort_order",
	"Prime":            "prime",
	"Active":           "active",
	"CreatedAt":        "created_at",
	"CreatedBy":        "created_by",
	"UpdatedAt":        "updated_at",
	"UpdatedBy":        "updated_by",
}

// makeStatement
func (repo *Model) makeStatement() sq.SelectBuilder {
	return repo.qb.Select(
		fieldMap["ID"],
		fieldMap["Name"],
		fieldMap["Url"],
		fieldMap["ShortDescription"],
		fieldMap["Description"],
		fieldMap["SortOrder"],
		fieldMap["Prime"],
		fieldMap["Active"],
		fieldMap["CreatedAt"],
		fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"],
		fieldMap["UpdatedBy"],
	).From(repo.table + " p")
}

// makeStatementByFilter
func (repo *Model) makeStatementByFilter(filter *Filter) sq.SelectBuilder {
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

	// Page
	if filter.Page == nil {
		filter.Page = new(uint64)
		*filter.Page = 1
	}

	// PerPage
	if filter.PerPage == nil {
		filter.PerPage = new(uint64)
		*filter.PerPage = 10
	}

	// Build query and OrderBy, OrderDir, Page, Limit
	statement := repo.makeStatement().OrderBy(fieldMap[*filter.OrderBy] + " " + *filter.OrderDir).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)

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

	// Urls
	if filter.Urls != nil {
		countUrls := len(*filter.Urls)

		if countUrls > 0 {
			statement = statement.Where(sq.Eq{fieldMap["Url"]: *filter.Urls})
		}

		*filter.Page = 1
		if (*filter.PerPage) > uint64(countUrls) {
			*filter.PerPage = uint64(countUrls)
		}
	}

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{fieldMap["Active"]: *filter.Active})
	}

	// Prime
	if filter.Prime != nil {
		statement = statement.Where(sq.Eq{fieldMap["Prime"]: *filter.Prime})
	}

	// Search
	if filter.Search != nil {
		statement = statement.Where(
			sq.Or{
				sq.Expr("LOWER("+fieldMap["Name"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["Url"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
			},
		)
	}

	return statement
}

// scanOneRow
func (repo *Model) scanOneRow(ctx context.Context, rows sq.RowScanner) (*Item, error) {
	var item = &Item{}
	err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.Url,
		&item.ShortDescription,
		&item.Description,
		&item.SortOrder,
		&item.Prime,
		&item.Active,
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
func (repo *Model) makeInsertStatement(ctx context.Context, item *Item) sq.InsertBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if ID is not set, generate a new UUID
	if item.ID == "" {
		item.ID = uuid.New().String()
	}

	return repo.qb.Insert(repo.table).
		Columns(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["ShortDescription"],
			fieldMap["Description"],
			fieldMap["SortOrder"],
			fieldMap["Prime"],
			fieldMap["Active"],
			fieldMap["CreatedAt"],
			fieldMap["CreatedBy"],
			fieldMap["UpdatedAt"],
			fieldMap["UpdatedBy"],
		).
		Values(
			item.ID,
			item.Name,
			item.Url,
			item.ShortDescription,
			item.Description,
			item.SortOrder,
			item.Prime,
			item.Active,
			"NOW()",
			by,
			"NOW()",
			by,
		)
}

// makeUpdateStatement
func (repo *Model) makeUpdateStatement(ctx context.Context, item *Item) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	return repo.qb.Update(repo.table).
		Set(fieldMap["Name"], item.Name).
		Set(fieldMap["Url"], item.Url).
		Set(fieldMap["ShortDescription"], item.ShortDescription).
		Set(fieldMap["Description"], item.Description).
		Set(fieldMap["SortOrder"], item.SortOrder).
		Set(fieldMap["Prime"], item.Prime).
		Set(fieldMap["Active"], item.Active).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by).
		Where(sq.Eq{fieldMap["ID"]: item.ID})
}

// makePatchStatement
func (repo *Model) makePatchStatement(ctx context.Context, id *string, fields *map[string]interface{}) sq.UpdateBuilder {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	statement := repo.qb.Update(repo.table).Where("id = ?", id)

	for field, value := range *fields {
		field = fieldMap[field]
		statement = statement.Set(field, value)
	}

	return statement.Set(fieldMap["UpdatedAt"], "NOW()").Set(fieldMap["UpdatedBy"], by)
}
