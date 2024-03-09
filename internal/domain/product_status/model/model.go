package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

type ProductStatusStorage interface {
	All(ctx context.Context, filter *Filter) ([]*ProductStatus, error)
	Create(ctx context.Context, productCategory *ProductStatus) (*ProductStatus, error)
	Get(ctx context.Context, id string) (*ProductStatus, error)
	GetByURL(ctx context.Context, url string) (*ProductStatus, error)
	Update(ctx context.Context, product *ProductStatus) (*ProductStatus, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}) (*ProductStatus, error)
	UpdatedAt(ctx context.Context, id string) (time.Time, error)
	TableUpdated(ctx context.Context) (time.Time, error)
	MaxSortOrder(ctx context.Context) (*uint32, error)
	Delete(ctx context.Context, id string) error
	Drop(ctx context.Context) error
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
