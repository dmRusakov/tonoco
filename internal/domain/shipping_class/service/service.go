package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/shipping_class/model"
	"time"
)

type repository interface {
	List(context.Context, *Filter) ([]*Item, error)
	Create(context.Context, *Item) (*Item, error)
	Get(context.Context, *string, *string) (*Item, error)
	Update(context.Context, *Item) (*Item, error)
	Patch(context.Context, *string, *map[string]interface{}) (*Item, error)
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
