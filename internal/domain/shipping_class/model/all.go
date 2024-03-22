package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (repo *Model) All(
	ctx context.Context,
	filter *Filter,
) ([]*Item, error) {
	// check standard filter parameter
	if filter.SortBy == nil {
		filter.SortBy = new(string)
		*filter.SortBy = "SortOrder"
	}

	if filter.SortOrder == nil {
		filter.SortOrder = new(string)
		*filter.SortOrder = "ASC"
	}

	if filter.Page == nil {
		filter.Page = new(uint64)
		*filter.Page = 1
	}

	if filter.PerPage == nil {
		filter.PerPage = new(uint64)
		*filter.PerPage = 10
	}

	// build query
	statement := repo.makeSelect().OrderBy(fieldMap[*filter.SortBy] + " " + *filter.SortOrder).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)

	// Active
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{fieldMap["Active"]: *filter.Active})
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

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	// execute the SQL query
	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over the result set
	var items []*Item
	for rows.Next() {
		item, err := repo.scanGet(ctx, rows)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	// done
	return items, nil
}
