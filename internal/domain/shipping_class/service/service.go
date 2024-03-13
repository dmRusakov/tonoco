package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"time"
)

type repository interface {
	All(ctx context.Context, filter Filter) ([]Item, error)
	Create(ctx context.Context, productStatus Item, by string) (Item, error)
	Get(ctx context.Context, id string) (Item, error)
	GetByURL(ctx context.Context, url string) (Item, error)
	Update(ctx context.Context, productStatus Item) (Item, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (Item, error)
	UpdatedAt(ctx context.Context, id string) (time.Time, error)
	TableUpdated(ctx context.Context) (time.Time, error)
	MaxSortOrder(ctx context.Context) (*uint32, error)
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repository model.ProductStatusStorage
}

func NewService(repository model.ProductStatusStorage) *Service {
	return &Service{repository: repository}
}
