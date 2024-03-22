package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *Model) Patch(ctx context.Context, id string, fields map[string]interface{}) (*Item, error) {
	// build query
	statement := repo.makePatch(ctx, id, fields)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	// add tracing
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
	return repo.Get(ctx, id)
}
