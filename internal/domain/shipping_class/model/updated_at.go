package model

import (
	"context"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"time"
)

// UpdatedAt - Get item updated at by id
func (repo *Model) UpdatedAt(
	ctx context.Context,
	id string,
) (*time.Time, error) {
	// build query
	statement := repo.makeUpdatedAt(id)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	// execute the SQL query
	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
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

// TableUpdated - Get table updated at
func (repo *Model) TableUpdated(
	ctx context.Context,
) (*time.Time, error) {
	// build query
	statement := repo.makeTableUpdated()

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)
		return nil, err
	}

	tracing.SpanEvent(ctx, "Select Table Updated")
	tracing.TraceVal(ctx, "SQL", query)
	// execute the SQL query
	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		tracing.Error(ctx, err)
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		err = psql.ErrNoRowForTableUpdated()
		tracing.Error(ctx, err)
		return nil, err
	}

	// return the updated at
	return repo.scanUpdatedAt(ctx, rows)
}
