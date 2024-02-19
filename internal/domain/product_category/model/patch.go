package model

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *ProductCategoryModel) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*ProductCategory, error) {
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

	// Add tracing
	tracing.SpanEvent(ctx, "Patch Product")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		argStr, ok := arg.(string)
		if !ok {
			// arg is not of type string, handle the error or continue to next iteration
			continue
		}
		tracing.TraceVal(ctx, "Arg"+strconv.Itoa(i), argStr)
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
