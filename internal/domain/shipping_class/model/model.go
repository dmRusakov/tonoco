package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

type ShippingClassStorage interface {
	All(ctx context.Context, filter *entity.ShippingClassFilter) ([]*entity.ShippingClass, error)
	Create(ctx context.Context, productCategory *entity.ShippingClass) (*entity.ShippingClass, error)
	Get(ctx context.Context, id string) (*entity.ShippingClass, error)
	GetByURL(ctx context.Context, url string) (*entity.ShippingClass, error)
	Update(ctx context.Context, product *entity.ShippingClass) (*entity.ShippingClass, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}) (*entity.ShippingClass, error)
	UpdatedAt(ctx context.Context, id string) (time.Time, error)
	TableUpdated(ctx context.Context) (time.Time, error)
	MaxSortOrder(ctx context.Context) (*uint32, error)
	Delete(ctx context.Context, id string) error
	Drop(ctx context.Context) error
}

// Model is a struct that contains the SQL statement builder and the PostgreSQL client.
type Model struct {
	table  string
	qb     sq.StatementBuilderType
	client psql.Client
}

// ShippingClassStorage is a constructor function that returns a new instance of the Model.
func NewShippingClassStorage(client psql.Client) *Model {
	return &Model{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
		table:  "public.product_status",
	}
}
