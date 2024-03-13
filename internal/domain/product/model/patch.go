package model

import (
	"context"
	"fmt"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
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
		tracing.Error(ctx, err)
		return nil, err
	}

	// trace the SQL query and arguments
	tracing.SpanEvent(ctx, "Patch Product query")
	tracing.TraceVal(ctx, "sql", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the query
	cmd, err := repo.client.Exec(ctx, query, args...)
	if err != nil {
		tracing.Error(ctx, err)
		return nil, err
	}

	if cmd.RowsAffected() == 0 {
		err = psql.ErrNothingInserted
		tracing.Error(ctx, err)
		return nil, err
	}

	// retrieve the updated Item
	product, err := repo.Get(ctx, id)
	if err != nil {
		tracing.Error(ctx, err)
		return nil, err
	}

	return product, nil
}
