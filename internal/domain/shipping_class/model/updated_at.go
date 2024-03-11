package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"time"
)

// UpdatedAt - Get item updated at by id
func (repo *ShippingClassModel) UpdatedAt(
	ctx context.Context,
	id string,
) (time.Time, error) {
	// build query
	statement := repo.qb.
		Select(fieldMap["UpdatedAt"]).
		From(repo.table).
		Where(sq.Eq{fieldMap["ID"]: id})

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return time.Time{}, err
	}

	// execute the SQL query
	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		return time.Time{}, err
	}

	defer rows.Close()

	if !rows.Next() {
		return time.Time{}, nil
	}

	// scan the result set into a slice of Product structs
	var updatedAt time.Time
	if err = rows.Scan(
		&updatedAt,
	); err != nil {
		return time.Time{}, err
	}

	// return the updated at
	return updatedAt, nil
}

// TableUpdated - Get table updated at
func (repo *ShippingClassModel) TableUpdated(
	ctx context.Context,
) (time.Time, error) {
	// build query
	statement := repo.qb.
		Select("max(updated_at)").
		From(repo.table)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return time.Time{}, err
	}

	// execute the SQL query
	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		return time.Time{}, err
	}

	defer rows.Close()

	if !rows.Next() {
		return time.Time{}, nil
	}

	// scan the result set into a slice of Product structs
	var updatedAt time.Time
	if err = rows.Scan(
		&updatedAt,
	); err != nil {
		return time.Time{}, err
	}

	// return the updated at
	return updatedAt, nil
}
