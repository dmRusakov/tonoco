package model

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *Model) Update(ctx context.Context, product *entity.ShippingClass) (*entity.ShippingClass, error) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// build query
	statement := repo.qb.Update(repo.table).
		Set(fieldMap["Name"], product.Name).
		Set(fieldMap["Url"], product.Url).
		Set(fieldMap["SortOrder"], product.SortOrder).
		Set(fieldMap["Active"], product.Active).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by).
		Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), product.ID)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	// Add tracing
	tracing.SpanEvent(ctx, "Update ShippingClass")
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
		return nil, err
	}

	if cmd.RowsAffected() == 0 {
		err = psql.ErrNothingInserted
		tracing.Error(ctx, err)

		return nil, err
	}

	// retrieve the updated Product
	productStatus, err := repo.Get(ctx, product.ID)
	if err != nil {
		return nil, err
	}

	return productStatus, nil
}