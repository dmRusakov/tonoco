package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (repo *ProductStatusModel) All(
	ctx context.Context,
	filter *Filter,
) ([]*ProductStatus, error) {
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
			fieldMap["SortOrder"],
			fieldMap["Active"],
		).
		From(repo.table + " p")

	// add the active filter if it is not nil
	if filter.Active != nil {
		statement = statement.Where(sq.Eq{fieldMap["Active"]: *filter.Active})
	}

	statement = statement.OrderBy(fieldMap[*filter.SortBy] + " " + *filter.SortOrder).
		Offset((*filter.Page - 1) * *filter.PerPage).Limit(*filter.PerPage)

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
			&productStatus.Url,
			&productStatus.SortOrder,
			&productStatus.Active,
		); err != nil {
			return nil, err
		}

		productStatuses = append(productStatuses, productStatus)
	}

	return productStatuses, nil
}
