package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"strconv"
)

func (repo *ProductStatusModel) Create(ctx context.Context, productStatus *ProductStatus, by string) (*ProductStatus, error) {
	// if ID is not set, generate a new UUID
	if productStatus.ID == "" {
		productStatus.ID = uuid.New().String()
	}

	// build query
	statement := repo.qb.Insert(repo.table).
		Columns(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Slug"],
			fieldMap["SortOrder"],
			fieldMap["Active"],
			fieldMap["CreatedAt"],
			fieldMap["CreatedBy"],
			fieldMap["UpdatedAt"],
			fieldMap["UpdatedBy"]).
		Values(
			productStatus.ID,
			productStatus.Name,
			productStatus.Slug,
			productStatus.SortOrder,
			productStatus.Active,
			"NOW()",
			by,
			"NOW()",
			by,
		)

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

	if cmd.RowsAffected() == 0 {
		return nil, psql.ErrNothingInserted
	}

	return productStatus, nil
}
