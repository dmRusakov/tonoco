package model

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *ProductModel) Delete(ctx context.Context, id string) error {
	// build query
	statement := repo.qb.Delete(repo.table).Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), id)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		tracing.Error(ctx, err)
		return err
	}

	// trace the SQL query and arguments
	tracing.SpanEvent(ctx, "Delete Product query")
	tracing.TraceVal(ctx, "sql", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the query
	_, execErr := repo.client.Exec(ctx, query, args...)
	if execErr != nil {
		tracing.Error(ctx, execErr)
		return execErr
	}

	return nil
}
