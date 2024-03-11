package model

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *Model) Update(ctx context.Context, productCategory *entity.ProductCategory) (*entity.ProductCategory, error) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// build query
	statement := repo.qb.
		Update(repo.table).
		Set(fieldMap["Name"], productCategory.Name).
		Set(fieldMap["Url"], productCategory.Url).
		Set(fieldMap["ShortDescription"], productCategory.ShortDescription).
		Set(fieldMap["Description"], productCategory.Description).
		Set(fieldMap["SortOrder"], productCategory.SortOrder).
		Set(fieldMap["Prime"], productCategory.Prime).
		Set(fieldMap["Active"], productCategory.Active).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by).
		Where(fieldMap["ID"]+" = ?", productCategory.ID)

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
	cmd, err := repo.client.Exec(ctx, query, args...)
	if err != nil {
		err = psql.ErrExec(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	if cmd.RowsAffected() == 0 {
		err = psql.ErrNothingInserted
		tracing.Error(ctx, err)

		return nil, err
	}

	return repo.Get(ctx, productCategory.ID)
}
