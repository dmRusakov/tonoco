package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/price_type/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"time"
)

type Item = entity.PriceType
type Filter = entity.PriceTypeFilter

type Repository interface {
	Get(context.Context, *Filter) (*Item, error)
	List(context.Context, *Filter, bool) (*map[string]Item, *uint64, error)
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
	itemCash   map[string]Item
	itemsCash  map[string]map[string]Item
	countCash  map[string]uint64
}

func NewService(repository model.Storage) *Service {
	return &Service{
		repository: repository,
		itemCash:   make(map[string]Item),
		itemsCash:  make(map[string]map[string]Item),
		countCash:  make(map[string]uint64),
	}
}
