package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

type Storage interface {
	List(context.Context, *Filter) ([]*Item, error)
	Create(context.Context, *Item) (*Item, error)
	Get(context.Context, *string, *string) (*Item, error)
	Update(context.Context, *Item) (*Item, error)
	Patch(context.Context, *string, *map[string]interface{}) (*Item, error)
	UpdatedAt(context.Context, *string) (*time.Time, error)
	TableIndexCount(context.Context) (*uint64, error)
	MaxSortOrder(context.Context) (*uint32, error)
	Delete(context.Context, string) error
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
		table:  "shipping_class",
	}
}
