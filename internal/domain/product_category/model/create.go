package model

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"strconv"
)

func (repo *Model) Create(ctx context.Context, productCategory *entity.ProductCategory) (*entity.ProductCategory, error) {
	// get user_id from context
	by := ctx.Value("user_id").(string)

	// if ID is not set, generate a new UUID
	if productCategory.ID == "" {
		productCategory.ID = uuid.New().String()
	}

	// build query
	statement := repo.qb.Insert(repo.table).
		Columns(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["ShortDescription"],
			fieldMap["Description"],
			fieldMap["SortOrder"],
			fieldMap["Prime"],
			fieldMap["Active"],
			fieldMap["CreatedAt"],
			fieldMap["CreatedBy"],
			fieldMap["UpdatedAt"],
			fieldMap["UpdatedBy"]).
		Values(
			productCategory.ID,
			productCategory.Name,
			productCategory.Url,
			productCategory.ShortDescription,
			productCategory.Description,
			productCategory.SortOrder,
			productCategory.Prime,
			productCategory.Active,
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

	return productCategory, nil
}
