package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/price_type/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
	"time"
)

type Item = entity.PriceType
type Filter = entity.PriceTypeFilter

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
	repository          model.Storage
	RegularPriceTypeIds []uuid.UUID
	SpecialPriceTypeIds []uuid.UUID
}

func NewService(repository *model.Model) *Service {
	service := &Service{
		repository: repository,
	}

	// regular price
	regularPriceTypeFilter := &entity.PriceTypeFilter{
		Urls:           &[]string{"regular"},
		IsPublic:       entity.BoolPtr(true),
		IsIdsOnly:      entity.BoolPtr(true),
		IsCount:        entity.BoolPtr(false),
		IsUpdateFilter: entity.BoolPtr(true),
	}

	_, err := service.List(context.Background(), regularPriceTypeFilter)
	if err != nil {
		panic(err)
	}

	if regularPriceTypeFilter.Ids != nil {
		service.RegularPriceTypeIds = *regularPriceTypeFilter.Ids
	} else {
		service.RegularPriceTypeIds = []uuid.UUID{}
	}

	// special price
	specialPriceTypeFilter := &entity.PriceTypeFilter{
		Urls:           &[]string{"special", "sale"},
		IsPublic:       entity.BoolPtr(true),
		IsIdsOnly:      entity.BoolPtr(true),
		IsCount:        entity.BoolPtr(false),
		IsUpdateFilter: entity.BoolPtr(true),
	}

	_, err = service.List(context.Background(), specialPriceTypeFilter)
	if err != nil {
		panic(err)
	}

	if specialPriceTypeFilter.Ids != nil {
		service.SpecialPriceTypeIds = *specialPriceTypeFilter.Ids
	} else {
		service.SpecialPriceTypeIds = []uuid.UUID{}
	}

	// done
	return service
}

func (s *Service) Get(ctx context.Context, filter *Filter) (*Item, error) {
	return s.repository.Get(ctx, filter)
}

func (s *Service) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	return s.repository.List(ctx, filter)
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
