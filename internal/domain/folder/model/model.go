package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"time"
)

type Item = entity.Folder
type Filter = entity.FolderFilter

type Storage interface {
	Get(context.Context, *Filter) (*Item, error)
	List(context.Context, *Filter, bool) (*map[string]Item, error)
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
	table       string
	qb          sq.StatementBuilderType
	client      psql.Client
	dbFieldCash map[string]string
}

// NewStorage is a constructor function that returns a new instance of the Model.
func NewStorage(client psql.Client) *Model {
	return &Model{
		qb:          sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:      client,
		table:       "file",
		dbFieldCash: map[string]string{},
	}
}
