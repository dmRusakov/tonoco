package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
)

func (repo *ProductStatusModel) All(
	ctx context.Context,
	filter *entity.ProductStatusFilter,
) ([]*entity.ProductStatus, error) {
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
		From(repo.table + " p").
		OrderBy(fieldMap[*filter.SortBy] + " " + *filter.SortOrder).
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

	var productStatuses []*entity.ProductStatus
	for rows.Next() {
		// scan the result set into a ProductStatus struct
		productStatus := &entity.ProductStatus{}
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
