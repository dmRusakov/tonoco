package model

import "context"

// MaxSortOrder - get the maximum value for the sort order in the table
func (repo *ProductCategoryModel) MaxSortOrder(ctx context.Context) (*int, error) {
	// build query
	statement := repo.qb.
		Select("max(sort_order)").
		From(repo.table).
		GroupBy("sort_order")

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

	if !rows.Next() {
		return nil, nil
	}

	// scan the result set into a slice of Product structs
	var sortOrder int
	if err = rows.Scan(
		&sortOrder,
	); err != nil {
		return nil, err
	}

	// return the max sort order
	return &sortOrder, nil
}
