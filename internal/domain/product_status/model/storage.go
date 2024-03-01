package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
)

type ProductCategoryStorage interface {
	All(ctx context.Context, filter *Filter) ([]*ProductStatus, error)
	Create(ctx context.Context, productCategory *ProductStatus, by string) (*ProductStatus, error)
	Get(ctx context.Context, id string) (*ProductStatus, error)
	Update(ctx context.Context, product *ProductStatus, by string) (*ProductStatus, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*ProductStatus, error)
	Delete(ctx context.Context, id string) error
}

// ProductStatusModel is a struct that contains the SQL statement builder and the PostgreSQL client.
type ProductStatusModel struct {
	table  string
	qb     sq.StatementBuilderType
	client psql.Client
}

// ProductStatusStorage is a constructor function that returns a new instance of the ProductStatusModel.
func NewProductStatusStorage(client psql.Client) *ProductStatusModel {
	return &ProductStatusModel{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
		table:  "public.product_status",
	}
}
