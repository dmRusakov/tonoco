package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/shipping_class/model"
	"time"
)

type repository interface {
	All(context.Context, *Filter) ([]*Item, error)
	Create(context.Context, *Item) (*Item, error)
	Get(context.Context, string) (*Item, error)
	GetByURL(context.Context, string) (*Item, error)
	Update(context.Context, *Item) (*Item, error)
	Patch(context.Context, string, map[string]interface{}) (*Item, error)
	UpdatedAt(context.Context, string) (time.Time, error)
	TableUpdated(context.Context) (time.Time, error)
	MaxSortOrder(context.Context) (*uint32, error)
	Delete(context.Context, string) error
}

type Service struct {
	repository model.Storage
}

func NewService(repository model.Storage) *Service {
	return &Service{repository: repository}
}
