package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/domain/store/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"time"
)

type Item = entity.Store
type Filter = entity.StoreFilter

type Repository interface {
	InitStore(context.Context, string) (*Item, error)
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
	DefaultStore *Item
	repository   model.Storage
	itemCash     map[string]Item
	itemsCash    map[string]map[string]Item
	countCash    map[string]uint64
}

func NewService(repository model.Storage, cfg *config.Config) (*Service, error) {
	// Init service
	service := &Service{
		repository: repository,
		itemCash:   make(map[string]Item),
		itemsCash:  make(map[string]map[string]Item),
		countCash:  make(map[string]uint64),
	}

	// If not in cache, proceed with initialization
	item, err := repository.Get(context.Background(), &Filter{Urls: &[]string{cfg.StoreUrl}})
	if err != nil {
		return nil, err
	}

	// Store the newly initialized item in the cache
	service.DefaultStore = item

	return service, nil
}
