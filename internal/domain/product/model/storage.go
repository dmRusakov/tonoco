package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
)

type ProductStorage interface {
	Create(ctx context.Context, product *Product, by string) (*Product, error)
	Get(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product, by string) (*Product, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*Product, error)
	Delete(ctx context.Context, id string) error
}

// ProductModel is a struct that contains the SQL statement builder and the PostgreSQL client.
type ProductModel struct {
	table  string
	qb     sq.StatementBuilderType
	client psql.Client
}

// NewProductStorage is a constructor function that returns a new instance of ProductModel.
func NewProductStorage(client psql.Client) *ProductModel {
	return &ProductModel{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
		table:  "public.product",
	}
}
