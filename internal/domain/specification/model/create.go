package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *Model) Create(ctx context.Context, item *Item) (*Item, error) {
	// build query
	statement := repo.makeInsert(ctx, item)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Insert Product query")
	tracing.TraceVal(ctx, "sql", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the query
	cmd, execErr := repo.client.Exec(ctx, query, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return nil, execErr
	}

	// check if the item was created
	if cmd.RowsAffected() == 0 {
		return nil, psql.ErrNothingInserted
	}

	// return the newly created item
	return repo.Get(ctx, item.ID)
}
