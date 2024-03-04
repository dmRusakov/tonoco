package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (repo *ProductCategoryModel) All(
	ctx context.Context,
	filter *Filter,
) ([]*ProductCategory, error) {
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
	statement := repo.qb.
		Select(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["ShortDescription"],
			fieldMap["Description"],
			fieldMap["SortOrder"],
			fieldMap["Prime"],
			fieldMap["Active"],
		).
		From(repo.table + " p").
		OrderBy(fieldMap[*filter.SortBy] + " " + *filter.SortOrder).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)

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
				sq.Expr("LOWER("+fieldMap["ShortDescription"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
				sq.Expr("LOWER("+fieldMap["Description"]+") ILIKE LOWER(?)", "%"+*filter.Search+"%"),
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

	var productCategories []*ProductCategory
	for rows.Next() {
		var productCategory ProductCategory
		err = rows.Scan(
			&productCategory.ID,
			&productCategory.Name,
			&productCategory.Url,
			&productCategory.ShortDescription,
			&productCategory.Description,
			&productCategory.SortOrder,
			&productCategory.Prime,
			&productCategory.Active,
		)
		if err != nil {
			return nil, err
		}
		productCategories = append(productCategories, &productCategory)
	}

	return productCategories, nil
}
