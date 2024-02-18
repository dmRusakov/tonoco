package model

import (
	"context"
	"fmt"
)

func (repo *ProductModel) Delete(ctx context.Context, id string) error {
	// build query
	statement := repo.qb.Delete(repo.table).Where(fmt.Sprintf("%s = ?", fieldMap["ID"]), id)

	// convert the SQL statement to a string
	query, args, err := statement.ToSql()
	if err != nil {
		return err
	}

	// execute the query
	_, err = repo.client.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
