package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (repo *ProductStatusModel) All(
	ctx context.Context,
	active *bool,
	sortBy *string, // field to sort by
	sortOrder *string, // sort order (asc, desc)
	page *uint32, // page number
	perPage *uint32, // number of items per page
) ([]*ProductStatus, error) {
	// build query
	statement := repo.qb.
		Select(repo.makeDbRequestColumns()...).
		From(repo.table + " p")

	// add the active filter if it is not nil
	if active != nil {
		statement = statement.Where(sq.Eq{fieldMap["Active"]: *active})
	}

	// add the sort order if it is not nil
	if sortBy == nil {
		sort := fieldMap["SortOrder"]
		sortBy = &sort
	}

	if sortOrder == nil {
		sort := "asc"
		sortOrder = &sort
	}

	statement = statement.OrderBy(*sortBy + " " + *sortOrder)

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

	var productStatuses []*ProductStatus
	for rows.Next() {
		// scan the result set into a ProductStatus struct
		productStatus := &ProductStatus{}
		if err = rows.Scan(
			&productStatus.ID,
			&productStatus.Name,
			&productStatus.Slug,
			&productStatus.SortOrder,
			&productStatus.Active,
		); err != nil {
			return nil, err
		}

		productStatuses = append(productStatuses, productStatus)
	}

	return productStatuses, nil
}
