package model

import (
	"context"
	"fmt"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *Model) Patch(ctx context.Context, id string, fields map[string]interface{}) (*Item, error) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

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

	// get item
	item, err := repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// return the item
	return item, nil
}
