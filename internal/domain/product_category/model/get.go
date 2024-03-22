package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
)

// Get returns a single product category by ID.
func (repo *Model) Get(ctx context.Context, id string) (*Item, error) {
	// build query
	statement := repo.qb.
		Select(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["ShortDescription"],
			fieldMap["Description"],
			fieldMap["SortOrder"],
			fieldMap["Prime"],
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

	tracing.SpanEvent(ctx, "Select Item")
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

	// check if the result set is empty
	if !rows.Next() {
		return nil, psql.ErrNoRowForID(id)
	}

	// scan the result set into a slice of Item structs
	item := &Item{}
	if err = rows.Scan(
		&item.ID,
		&item.Name,
		&item.Url,
		&item.ShortDescription,
		&item.Description,
		&item.SortOrder,
		&item.Prime,
		&item.Active,
	); err != nil {
		err = psql.ErrScan(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	// done
	return item, nil
}

// GetByURL returns a single product category by URL.
func (repo *Model) GetByURL(ctx context.Context, url string) (*Item, error) {
	// build query
	statement := repo.qb.
		Select(
			fieldMap["ID"],
			fieldMap["Name"],
			fieldMap["Url"],
			fieldMap["ShortDescription"],
			fieldMap["Description"],
			fieldMap["SortOrder"],
			fieldMap["Prime"],
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

	tracing.SpanEvent(ctx, "Select Item")
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

	// scan the result set into a slice of Item structs
	productCategory := &Item{}
	if err = rows.Scan(
		&productCategory.ID,
		&productCategory.Name,
		&productCategory.Url,
		&productCategory.ShortDescription,
		&productCategory.Description,
		&productCategory.SortOrder,
		&productCategory.Prime,
		&productCategory.Active,
	); err != nil {
		err = psql.ErrScan(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	return productCategory, nil
}
