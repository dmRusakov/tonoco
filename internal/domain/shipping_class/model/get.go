package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

func (repo *Model) Get(ctx context.Context, id string) (*Item, error) {
	// build query
	statement := repo.qb.
		Select(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["SortOrder"],
			fieldMap["Active"],
		).
		From(repo.table + " p").
		Where(sq.Eq{fieldMap["ID"]: id})

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Select Product")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the SQL query
	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, psql.ErrNoRowForID(id)
	}

	// scan the result set into a ShippingClass struct
	productStatus := &Item{}
	if err = rows.Scan(
		&productStatus.ID,
		&productStatus.Name,
		&productStatus.Url,
		&productStatus.SortOrder,
		&productStatus.Active,
	); err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)

		return nil, err
	}

	return (*Item)(productStatus), nil
}

// GetByURL - get a product status by URL
func (repo *Model) GetByURL(ctx context.Context, url string) (*Item, error) {
	// build query
	statement := repo.qb.
		Select(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["SortOrder"],
			fieldMap["Active"],
		).
		From(repo.table + " p").
		Where(sq.Eq{fieldMap["Url"]: url})

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Select Product")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	// execute the SQL query
	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, psql.ErrNoRowForURL(url)
	}

	// scan the result set into a ShippingClass struct
	productStatus := &Item{}
	if err = rows.Scan(
		&productStatus.ID,
		&productStatus.Name,
		&productStatus.Url,
		&productStatus.SortOrder,
		&productStatus.Active,
	); err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)

		return nil, err
	}

	return (*Item)(productStatus), nil
}
