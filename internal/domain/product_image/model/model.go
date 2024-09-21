package model

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/dmRusakov/tonoco/internal/entity"
	psql "github.com/dmRusakov/tonoco/pkg/postgresql"
	"github.com/google/uuid"
	"time"
)

type ProductImage = entity.ProductImage
type ProductImageFilter = entity.ProductImageFilter

type Storage interface {
	Get(context.Context, *ProductImageFilter) (*ProductImage, error)
	List(context.Context, *ProductImageFilter, bool) (*map[uuid.UUID]ProductImage, *uint64, error)
	Create(context.Context, *ProductImage) (*uuid.UUID, error)
	Update(context.Context, *ProductImage) error
	Patch(context.Context, *uuid.UUID, *map[string]interface{}) error
	UpdatedAt(context.Context, *uuid.UUID) (*time.Time, error)
	MaxSortOrder(context.Context) (*uint64, error)
	TableIndexCount(context.Context) (*uint64, error)
	Delete(context.Context, *uuid.UUID) error
}

type Model struct {
	table       string
	qb          sq.StatementBuilderType
	client      psql.Client
	dbFieldCash map[string]string
}

func NewStorage(client psql.Client) *Model {
	return &Model{
		qb:          sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client:      client,
		table:       "product_image",
		dbFieldCash: map[string]string{},
	}
}
