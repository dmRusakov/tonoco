package dao

import (
	"github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/pkg/postgresql"
)

type ProductDAO struct {
	qb     squirrel.StatementBuilderType
	client postgresql.Client
}

func NewProductStorage(client postgresql.Client) *ProductDAO {
	return &ProductDAO{
		qb:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client: client,
	}
}
