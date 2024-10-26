package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/stock_quantity/model"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/google/uuid"
	"time"
)

type Item = db.StockQuantity
type Filter = db.StockQuantityFilter

type Repository interface {
	Get(ctx context.Context, filter *Filter) (*Item, error)
	List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error)
	Create(ctx context.Context, item *Item) (*uuid.UUID, error)
	Update(ctx context.Context, item *Item) error
	Patch(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) error
	UpdatedAt(ctx context.Context, id *uuid.UUID) (*time.Time, error)
	TableIndexCount(ctx context.Context) (*uint64, error)
	MaxSortOrder(ctx context.Context) (*uint64, error)
	Delete(ctx context.Context, id *uuid.UUID) error
}

type Service struct {
	repository model.Storage
}

func NewService(repository *model.Model) *Service {
	return &Service{
		repository: repository,
	}
}
