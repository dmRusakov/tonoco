package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/tag_type/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
	"time"
)

type Item = entity.TagType
type Filter = entity.TagTypeFilter

type Repository interface {
	Get(ctx context.Context, filter *Filter) (*Item, error)
	GetDefaultIds(name string) (*entity.DefaultTagTypes, error)
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
	defaultIds  map[string]*entity.DefaultTagTypes
}

func NewService(repository *model.Model) *Service {
	return &Service{
		repository:  repository,
		defaultItem: make(map[string]*Item),
		defaultIds:  make(map[string]*entity.DefaultTagTypes),
	}
}

func (s *Service) Get(ctx context.Context, filter *Filter) (*Item, error) {
	if filter == nil {
		return nil, entity.ErrFilterIsNil
	}
	return s.repository.Get(ctx, filter)
}

func (s *Service) GetDefault(name string) (*Item, error) {
	if s.defaultItem[name] != nil {
		return s.defaultItem[name], nil
	}

	item, err := s.Get(context.Background(), &Filter{
		Urls: &[]string{name},
	})
	if err != nil {
		return nil, err
	}

	s.defaultItem[name] = item

	return s.defaultItem[name], nil
}

func (s *Service) GetDefaultIds(name string) (*entity.DefaultTagTypes, error) {
	if s.defaultIds[name] != nil {
		return s.defaultIds[name], nil
	}

	switch name {
	case "list":
		item := &entity.DefaultTagTypes{
			TagTypes:    &map[uuid.UUID]entity.TagType{},
			TagOrder:    &map[uuid.UUID]uint64{},
			TagTypesIds: &[]uuid.UUID{},
		}

		var err error
		tagTypeFilter := &entity.TagTypeFilter{
			OrderBy:        entity.StringPtr("SortOrder"),
			OrderDir:       entity.StringPtr("ASC"),
			ListItem:       entity.BoolPtr(true),
			Active:         entity.BoolPtr(true),
			IsCount:        entity.BoolPtr(false),
			IsUpdateFilter: entity.BoolPtr(true),
		}

		item.TagTypes, err = s.List(context.Background(), tagTypeFilter)
		if err != nil {
			return nil, err
		}

		// tag order
		for i, tagType := range *tagTypeFilter.Ids {
			(*item.TagOrder)[tagType] = uint64(i)
		}

		// get tag types ids
		item.TagTypesIds = tagTypeFilter.Ids

		s.defaultIds[name] = item
	default:
		return nil, entity.ErrNotFound
	}

	return s.defaultIds[name], nil
}

func (s *Service) List(ctx context.Context, filter *Filter) (*map[uuid.UUID]Item, error) {
	if filter == nil {
		return nil, entity.ErrFilterIsNil
	}
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
