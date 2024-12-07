package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/shop/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/google/uuid"
	"time"
)

type Item = db.Shop
type Filter = db.ShopFilter

type Repository interface {
	Get(ctx context.Context, filter *Filter) (*Item, error)
	List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error)
	Ids(ctx context.Context, filter *Filter) (*[]uuid.UUID, error)
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
	urls       map[uuid.UUID]string
}

func NewService(repository *model.Model) *Service {
	service := &Service{
		repository: repository,
		urls:       make(map[uuid.UUID]string),
	}
	list, err := service.repository.List(context.Background(), &Filter{
		DataPagination: &entity.DataPagination{},
		DataConfig:     &entity.DataConfig{},
	})

	if err != nil {
		panic(err)
	}

	for _, item := range *list {
		service.urls[item.Id] = item.Url
	}

	return service
}

func (s *Service) Get(ctx context.Context, filter *Filter) (*Item, error) {
	return s.repository.Get(ctx, filter)
}

func (s *Service) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	return s.repository.List(ctx, filter)
}

func (s *Service) Ids(ctx context.Context, filter *Filter) (*[]uuid.UUID, error) {
	return s.repository.Ids(ctx, filter)
}

func (s *Service) Create(ctx context.Context, item *Item) (*uuid.UUID, error) {
	return s.repository.Create(ctx, item)
}

func (s *Service) Update(ctx context.Context, item *Item) error {
	return s.repository.Update(ctx, item)
}

func (s *Service) Patch(ctx context.Context, id *uuid.UUID, fields *map[string]interface{}) error {
	return s.repository.Patch(ctx, id, fields)
}

func (s *Service) UpdatedAt(ctx context.Context, id *uuid.UUID) (*time.Time, error) {
	return s.repository.UpdatedAt(ctx, id)
}

func (s *Service) TableIndexCount(ctx context.Context) (*uint64, error) {
	return s.repository.TableIndexCount(ctx)
}

func (s *Service) MaxSortOrder(ctx context.Context) (*uint64, error) {
	return s.repository.MaxSortOrder(ctx)
}

func (s *Service) Delete(ctx context.Context, id *uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}