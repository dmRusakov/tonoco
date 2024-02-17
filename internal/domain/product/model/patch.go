package model

import (
	"context"
)

// Patch is a method on the ProductModel struct that updates a product in the database by its ID.
func (repo *ProductModel) Patch(ctx context.Context, id string, fields map[string]interface{}) (*ProductStorage, error) {
	// build query
	statement := repo.qb.Update(table).Where("id = ?", id)

	// iterate over the fields map and add each field to the update statement
	for field, value := range fields {
		// get DB field name from ProductStorage struct and db tag
		field = fieldMap[field]
		statement = statement.Set(field, value)
	}

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	// execute the query
	_, err = repo.client.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	// retrieve the updated product
	product, err := repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
