package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

type Storage interface {
	Get(ctx context.Context, id *string, url *string) (*Item, error)
	List(context.Context, *Filter) ([]*Item, error)
	Create(context.Context, *Item) (*string, error)
	Update(context.Context, *Item) error
	Patch(context.Context, *string, *map[string]interface{}) error
	UpdatedAt(context.Context, *string) (*time.Time, error)
	MaxSortOrder(context.Context) (*uint64, error)
	TableIndexCount(context.Context) (*uint64, error)
	Delete(context.Context, *string) error
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
