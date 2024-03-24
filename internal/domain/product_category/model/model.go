package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

type Storage interface {
	All(ctx context.Context, filter *Filter) ([]*Item, error)
	Create(ctx context.Context, productCategory *Item) (*Item, error)
	Get(ctx context.Context, id string) (*Item, error)
	GetByURL(ctx context.Context, url string) (*Item, error)
	Update(ctx context.Context, product *Item) (*Item, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}) (*Item, error)
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
		table:  "product_category",
	}
}
