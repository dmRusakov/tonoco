package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *ProductCategoryModel) Update(ctx context.Context, product *ProductCategory, by string) (*ProductCategory, error) {
	// build query
	statement := repo.qb.
		Update(repo.table).
		Set(fieldMap["Name"], product.Name).
		Set(fieldMap["Slug"], product.Slug).
		Set(fieldMap["SortDescription"], product.SortDescription).
		Set(fieldMap["Description"], product.Description).
		Set(fieldMap["SortOrder"], product.SortOrder).
		Set(fieldMap["Prime"], product.Prime).
		Set(fieldMap["Active"], product.Active).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by).
		Where(fieldMap["ID"]+" = ?", product.ID)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Update Product")
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
		err = psql.ErrExec(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	return repo.Get(ctx, product.ID)
}
