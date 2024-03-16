package model

import (
	"context"
	"fmt"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *Model) Update(ctx context.Context, item *Item) (*Item, error) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// build query
	statement := repo.qb.Update(repo.table).
		Set(fieldMap["Name"], item.Name).
		Set(fieldMap["Url"], item.Url).
		Set(fieldMap["SortOrder"], item.SortOrder).
		Set(fieldMap["Active"], item.Active).
		Set(fieldMap["UpdatedAt"], "NOW()").
		Set(fieldMap["UpdatedBy"], by).
		Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), item.ID)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	// add tracing
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

	// return the updated item
	return repo.Get(ctx, item.ID)
}
