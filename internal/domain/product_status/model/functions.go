package model

import (
	"context"
	"fmt"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/dmRusakov/tonoco/pkg/tracing"
	"strconv"
	"time"
)

func (repo *Model) Get(ctx context.Context, id *string, url *string) (*Item, error) {
	// build query
	statement := repo.makeStatement()

	isParamSet := false

	// id
	if id != nil {
		statement = statement.Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), *id)
		isParamSet = true
	}

	// url
	if url != nil {
		statement = statement.Where(fmt.Sprintf("%s = ?", fieldMap["Url"]), *url)
		isParamSet = true
	}

	if !isParamSet {
		return nil, fmt.Errorf("id or url must be set")
	}

	// execute the query
	rows, err := psql.Get(ctx, repo.client, statement)
	if err != nil {
		return nil, err
	}

	// return the Item
	return repo.scanOneRow(ctx, rows)
}

func (repo *Model) List(ctx context.Context, filter *Filter) ([]*Item, error) {
	// build query
	statement := repo.makeStatementByFilter(filter)

	// execute the query
	rows, err := psql.List(ctx, repo.client, statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over the result set
	var items []*Item
	for rows.Next() {
		item, err := repo.scanOneRow(ctx, rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (repo *Model) Create(ctx context.Context, item *Item) error {
	// build query
	statement, err := repo.makeInsertStatement(ctx, item)
	if err != nil {
		return err
	}

	// execute the query
	err = psql.Create(ctx, repo.client, *statement)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Model) Update(ctx context.Context, item *Item) error {
	// build query
	statement := repo.makeUpdateStatement(ctx, item).Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), item.ID)

	// execute the query
	err := psql.Update(ctx, repo.client, statement)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Model) Patch(ctx context.Context, id *string, fields *map[string]interface{}) error {
	// build query
	statement := repo.makePatchStatement(ctx, id, fields)

	err := psql.Update(ctx, repo.client, statement)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Model) Delete(ctx context.Context, id *string) error {
	// build query
	statement := repo.qb.Delete(repo.table).Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), id)

	// execute the query to delete the item
	return psql.Delete(ctx, repo.client, statement)
}

func (repo *Model) UpdatedAt(ctx context.Context, id *string) (*time.Time, error) {
	// build query
	statement := repo.qb.Select(fieldMap["UpdatedAt"]).From(repo.table).Where("id = ?", id)

	rows, err := psql.Get(ctx, repo.client, statement)
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

func (repo *Model) TableIndexCount(ctx context.Context) (*uint64, error) {
	// build query
	statement := repo.qb.Select("n_tup_upd").From("pg_stat_user_tables").Where("relname = ?", repo.table)

	// execute the query
	rows, err := psql.Get(ctx, repo.client, statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var updatedAt string
	err = rows.Scan(&updatedAt)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	// convert the string to a uint64
	count, err := strconv.ParseUint(updatedAt, 10, 64)
	if err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)
		return nil, err
	}

	// return the updated at
	return &count, nil
}

func (repo *Model) MaxSortOrder(ctx context.Context) (*uint64, error) {
	// build query
	statement := repo.qb.
		Select("max(sort_order)").
		From(repo.table).
		GroupBy("sort_order")

	// execute the query
	rows, err := psql.Get(ctx, repo.client, statement)
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
