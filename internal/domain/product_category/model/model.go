package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

type ProductCategoryStorage interface {
	All(ctx context.Context, filter *entity.ProductCategoryFilter) ([]*entity.ProductCategory, error)
	Create(ctx context.Context, productCategory *entity.ProductCategory) (*entity.ProductCategory, error)
	Get(ctx context.Context, id string) (*entity.ProductCategory, error)
	GetByURL(ctx context.Context, url string) (*entity.ProductCategory, error)
	Update(ctx context.Context, product *entity.ProductCategory) (*entity.ProductCategory, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}) (*entity.ProductCategory, error)
	UpdatedAt(ctx context.Context, id string) (time.Time, error)
	TableUpdated(ctx context.Context) (time.Time, error)
	MaxSortOrder(ctx context.Context) (*uint32, error)
	Delete(ctx context.Context, id string) error
}

// Model is a struct that contains the SQL statement builder and the PostgreSQL client.
type Model struct {
	table  string
	qb     sq.StatementBuilderType
	client psql.Client
}

// NewStorage is a constructor function that returns a new instance of the Model.
func NewStorage(client psql.Client) *Model {
	return &Model{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
		table:  "public.product_category",
	}
}
