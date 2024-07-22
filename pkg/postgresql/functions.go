package postgresql

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/jackc/pgx/v5"
	"strconv"
	"time"
)

// List - List items.
func List(ctx context.Context, client Client, statement sq.SelectBuilder) (pgx.Rows, error) {
	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	// execute the SQL query
	rows, err := client.Query(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	// return the rows
	return rows, nil
}

// Get - Get item
func Get(ctx context.Context, client Client, statement sq.SelectBuilder) (pgx.Rows, error) {
	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	fmt.Println(query, "functions:36")
	if err != nil {
		err = ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Select Item")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the SQL query
	rows, err := client.Query(ctx, query, args...)
	if err != nil {
		err = ErrDoQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		err = ErrNoRows()
		tracing.Error(ctx, err)
		return nil, err
	}

	// return the Item
	return rows, nil
}

// Create - Create item.
func Create(ctx context.Context, client Client, statement *sq.InsertBuilder) error {
	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}

	tracing.SpanEvent(ctx, "Insert Product query")
	tracing.TraceVal(ctx, "sql", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the query
	cmd, execErr := client.Exec(ctx, query, args...)
	if execErr != nil {
		execErr = ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	// check if the item was created
	if cmd.RowsAffected() == 0 {
		return ErrNothingInserted
	}

	// done successfully
	return nil
}

// Update - Update item.
func Update(ctx context.Context, client Client, statement sq.UpdateBuilder) error {
	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
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

	// start a transaction
	tx, err := client.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// execute the query
	cmd, err := client.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		err = ErrNothingInserted
		tracing.Error(ctx, err)

		return err
	}

	// commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	// done successfully
	return nil
}

// Delete - Delete item.
func Delete(ctx context.Context, client Client, statement sq.DeleteBuilder) error {
	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		tracing.Error(ctx, err)
		return err
	}

	// trace the SQL query and arguments
	tracing.SpanEvent(ctx, "Delete ShippingClass query")
	tracing.TraceVal(ctx, "sql", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the query
	cmd, err := client.Exec(ctx, query, args...)
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

func UpdatedAt(ctx context.Context, client Client, statement sq.SelectBuilder) (*time.Time, error) {
	rows, err := Get(ctx, client, statement)
	if err != nil {
		return nil, err
	}

	// scan the result set into a slice of Item structs
	var updatedAt *time.Time
	if err = rows.Scan(
		&updatedAt,
	); err != nil {
		return nil, err
	}

	// return the updated at
	return updatedAt, nil
}

func TableIndexCount(ctx context.Context, client Client, statement sq.SelectBuilder) (*uint64, error) {
	// execute the query
	rows, err := Get(ctx, client, statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var updatedAt string
	err = rows.Scan(&updatedAt)
	if err != nil {
		err = ErrScan(ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	// convert the string to a uint64
	count, err := strconv.ParseUint(updatedAt, 10, 64)
	if err != nil {
		err = ErrScan(ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	// return the updated at
	return &count, nil
}

func MaxSortOrder(ctx context.Context, client Client, qb sq.StatementBuilderType, tableName *string) (*uint64, error) {
	// build query
	statement := qb.
		Select("max(sort_order)").
		From(*tableName).
		GroupBy("sort_order")

	// execute the query
	rows, err := Get(ctx, client, statement)
	if err != nil {
		return nil, err
	}

	// scan the result set into a slice of Item structs
	var sortOrder uint64
	if err = rows.Scan(
		&sortOrder,
	); err != nil {
		return nil, err
	}

	// return the max sort order
	return &sortOrder, nil
}
