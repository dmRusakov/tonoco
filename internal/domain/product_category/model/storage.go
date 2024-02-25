package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
)

type ProductCategoryStorage interface {
	Create(ctx context.Context, productCategory *ProductCategory, by string) (*ProductCategory, error)
	Get(ctx context.Context, id string) (*ProductCategory, error)
	Update(ctx context.Context, product *ProductCategory, by string) (*ProductCategory, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*ProductCategory, error)
	Delete(ctx context.Context, id string) error
}

// ProductCategoryModel is a struct that contains the SQL statement builder and the PostgreSQL client.
type ProductCategoryModel struct {
	table  string
	qb     sq.StatementBuilderType
	client psql.Client
}

// ProductCategoryStorage is a constructor function that returns a new instance of the ProductCategoryModel.
func NewProductCategoryStorage(client psql.Client) *ProductCategoryModel {
	return &ProductCategoryModel{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
		table:  "public.product_category",
	}
}
