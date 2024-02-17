package model

import (
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
)

const table = "public.product"

// ProductModel is a struct that contains the SQL statement builder and the PostgreSQL client.
type ProductModel struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

// NewProductStorage is a constructor function that returns a new instance of ProductModel.
func NewProductStorage(client psql.Client) *ProductModel {
	return &ProductModel{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}
