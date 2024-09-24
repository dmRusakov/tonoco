package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_image/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
	"time"
)

type Item = entity.ProductImage
type Filter = entity.ProductImageFilter

type Repository interface {
	Get(context.Context, *Filter) (*Item, error)
	List(context.Context, *Filter) (*map[uuid.UUID]Item, error)
	Create(context.Context, *Item) (*uuid.UUID, error)
	Update(context.Context, *Item) error
	Patch(context.Context, *uuid.UUID, *map[string]interface{}) error
	UpdatedAt(context.Context, *uuid.UUID) (*time.Time, error)
	TableIndexCount(context.Context) (*uint64, error)
	MaxSortOrder(context.Context) (*uint64, error)
	Delete(context.Context, *uuid.UUID) error
}

type Service struct {
	repository model.Storage
	itemCash   map[string]Item
	itemsCash  map[string]map[uuid.UUID]Item
	countCash  map[string]uint64
}

func NewService(repository *model.Model) *Service {
	return &Service{
		repository: repository,
		itemCash:   make(map[string]Item),
		itemsCash:  make(map[string]map[uuid.UUID]Item),
		countCash:  make(map[string]uint64),
	}
}