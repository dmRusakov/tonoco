package model

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

// Delete removes a ProductStatus from the database by its ID.
func (repo *ProductStatusModel) Delete(ctx context.Context, id string) error {
	// build query
	statement := repo.qb.Delete(repo.table).Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), id)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		tracing.Error(ctx, err)
		return err
	}

	// trace the SQL query and arguments
	tracing.SpanEvent(ctx, "Delete ProductStatus query")
	tracing.TraceVal(ctx, "sql", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the query
	cmd, err := repo.client.Exec(ctx, query, args...)
	if err != nil {
		tracing.Error(ctx, err)
		return err
	}

	// check if any rows were affected
	if cmd.RowsAffected() == 0 {
		err := fmt.Errorf("no rows affected")
		tracing.Error(ctx, err)
		return err
	}

	// return the result
	return nil
}

// Drop all ProductStatus records from the database.
func (repo *ProductStatusModel) Drop(ctx context.Context) error {
	// build query
	statement := repo.qb.Delete(repo.table)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		tracing.Error(ctx, err)
		return err
	}

	// trace the SQL query and arguments
	tracing.SpanEvent(ctx, "Drop ProductStatus query")
	tracing.TraceVal(ctx, "sql", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the query
	cmd, err := repo.client.Exec(ctx, query, args...)
	if err != nil {
		tracing.Error(ctx, err)
		return err
	}

	// check if any rows were affected
	if cmd.RowsAffected() == 0 {
		err := fmt.Errorf("no rows affected")
		tracing.Error(ctx, err)
		return err
	}

	// return the result
	return nil
}
