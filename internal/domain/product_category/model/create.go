package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"github.com/google/uuid"
	"strconv"
)

func (repo *ProductCategoryModel) Create(ctx context.Context, productCategory *ProductCategory, by string) (*ProductCategory, error) {
	// if ID is not set, generate a new UUID
	if productCategory.ID == "" {
		productCategory.ID = uuid.New().String()
	}

	// build query
	statement := repo.qb.Insert(repo.table).Columns(
		fieldMap["ID"], fieldMap["Name"], fieldMap["Slug"], fieldMap["SortDescription"], fieldMap["Description"],
		fieldMap["SortOrder"], fieldMap["Prime"], fieldMap["Active"], fieldMap["CreatedAt"], fieldMap["CreatedBy"],
		fieldMap["UpdatedAt"], fieldMap["UpdatedBy"]).Values(productCategory.ID, productCategory.Name, productCategory.Slug,
		productCategory.SortDescription, productCategory.Description, productCategory.SortOrder, productCategory.Prime,
		productCategory.Active, "NOW()", by, "NOW()", by)

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
