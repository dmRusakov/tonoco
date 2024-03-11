package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

type ShippingClassStorage interface {
	All(ctx context.Context, filter *ShippingClassFilter) ([]*ShippingClass, error)
	Create(ctx context.Context, productCategory *ShippingClass) (*ShippingClass, error)
	Get(ctx context.Context, id string) (*ShippingClass, error)
	GetByURL(ctx context.Context, url string) (*ShippingClass, error)
	Update(ctx context.Context, product *ShippingClass) (*ShippingClass, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}) (*ShippingClass, error)
	UpdatedAt(ctx context.Context, id string) (time.Time, error)
	TableUpdated(ctx context.Context) (time.Time, error)
	MaxSortOrder(ctx context.Context) (*uint32, error)
	Delete(ctx context.Context, id string) error
	Drop(ctx context.Context) error
}

// ShippingClassModel is a struct that contains the SQL statement builder and the PostgreSQL client.
type ShippingClassModel struct {
	table  string
	qb     sq.StatementBuilderType
	client psql.Client
}

// ShippingClassStorage is a constructor function that returns a new instance of the ShippingClassModel.
func NewShippingClassStorage(client psql.Client) *ShippingClassModel {
	return &ShippingClassModel{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
		table:  "public.product_status",
	}
}
