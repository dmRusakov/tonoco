package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"time"
)

type repository interface {
	Get(ctx context.Context, id *string, url *string) (*Item, error)
	List(context.Context, *Filter) ([]*Item, error)
	Create(context.Context, *Item) (*string, error)
	Update(context.Context, *Item) error
	Patch(context.Context, *string, *map[string]interface{}) error
	UpdatedAt(context.Context, *string) (*time.Time, error)
	TableIndexCount(context.Context) (*uint64, error)
	MaxSortOrder(context.Context) (*uint64, error)
	Delete(context.Context, *string) error
}

type Service struct {
	repository model.Storage
}

func NewService(repository model.Storage) *Service {
	return &Service{repository: repository}
}
