package dao

import (
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
)

type ProductDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewProductStorage(client psql.Client) *ProductDAO {
	return &ProductDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}
