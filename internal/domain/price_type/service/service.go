package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/price_type/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/google/uuid"
	"time"
)

type Item = db.PriceType
type Filter = db.PriceTypeFilter

type Repository interface {
	Get(ctx context.Context, filter *Filter) (*Item, error)
	GetDefault(name string) (*Item, error)
	GetDefaultIds(name string) ([]uuid.UUID, error)
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
	repository  model.Storage
	defaultItem map[string]*Item
	defaultIds  map[string]*[]uuid.UUID
}

func NewService(repository *model.Model) *Service {
	return &Service{
		repository:  repository,
		defaultItem: map[string]*Item{},
		defaultIds:  map[string]*[]uuid.UUID{},
	}
}

func (s *Service) Get(ctx context.Context, filter *Filter) (*Item, error) {
	return s.repository.Get(ctx, filter)
}

func (s *Service) GetDefault(name string) (*Item, error) {
	if s.defaultItem[name] != nil {
		return s.defaultItem[name], nil
	}

	item, err := s.Get(context.Background(), &Filter{
		Urls:     &[]string{name},
		IsPublic: pointer.BoolPtr(true),
	})
	if err != nil {
		return nil, err
	}

	s.defaultItem[name] = item

	return s.defaultItem[name], nil
}

func (s *Service) GetDefaultIds(name string) (*[]uuid.UUID, error) {
	if s.defaultIds[name] != nil {
		return s.defaultIds[name], nil
	}

	switch name {
	case "regular":
		regularPriceTypeFilter := &Filter{
			Urls:           &[]string{"regular"},
			IsPublic:       pointer.BoolPtr(true),
			IsIdsOnly:      pointer.BoolPtr(true),
			IsCount:        pointer.BoolPtr(false),
			IsUpdateFilter: pointer.BoolPtr(true),
		}

		_, err := s.List(context.Background(), regularPriceTypeFilter)
		if err != nil {
			return nil, err
		}

		if regularPriceTypeFilter.Ids != nil {
			s.defaultIds[name] = regularPriceTypeFilter.Ids
		} else {
			s.defaultIds[name] = nil
		}
	case "special":
		filter := &Filter{
			Urls:           &[]string{"special", "sale"},
			IsPublic:       pointer.BoolPtr(true),
			IsIdsOnly:      pointer.BoolPtr(true),
			IsCount:        pointer.BoolPtr(false),
			IsUpdateFilter: pointer.BoolPtr(true),
		}

		_, err := s.List(context.Background(), filter)
		if err != nil {
			return nil, err
		}

		if filter.Ids != nil {
			s.defaultIds[name] = filter.Ids
		} else {
			s.defaultIds[name] = nil
		}
	default:
		return nil, entity.ErrNotFound
	}

	return s.defaultIds[name], nil
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
