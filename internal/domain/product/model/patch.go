package model

import (
	"context"
	"fmt"
)

// Patch is a method on the ProductModel struct that updates a Product in the database by its ID.
func (repo *ProductModel) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*Product, error) {
	// build query
	statement := repo.qb.Update(repo.table).Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), id)

	// iterate over the fields map and add each field to the update statement
	for field, value := range fields {
		// get DB field name from Product struct and db tag
		field = fieldMap[field]
		statement = statement.Set(field, value)
	}

	// add the updated_at field
	statement = statement.Set(fieldMap["UpdatedAt"], "NOW()")

	// add the updated_by field
	statement = statement.Set(fieldMap["UpdatedBy"], by)

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

	// retrieve the updated Product
	product, err := repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
