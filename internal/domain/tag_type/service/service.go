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
	GetTagTypesForList(ctx context.Context) *entity.DefaultTagTypes
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
	repository     model.Storage
	defaultForList *entity.DefaultTagTypes
}

func NewService(repository *model.Model) *Service {
	return &Service{
		repository:     repository,
		defaultForList: nil,
	}
}

func (s *Service) Get(ctx context.Context, filter *Filter) (*Item, error) {
	if filter == nil {
		return nil, entity.ErrFilterIsNil
	}
	return s.repository.Get(ctx, filter)
}

func (s *Service) GetTagTypesForList(ctx context.Context) *entity.DefaultTagTypes {
	if s.defaultForList != nil {
		return s.defaultForList
	}

	var tagTypes *map[uuid.UUID]entity.TagType
	var tagTypesIds *[]uuid.UUID
	var tagOrder map[uuid.UUID]uint32 = make(map[uuid.UUID]uint32)
	var err error
	tagTypeFilter := &entity.TagTypeFilter{
		OrderBy:        entity.StringPtr("SortOrder"),
		OrderDir:       entity.StringPtr("ASC"),
		ListItem:       entity.BoolPtr(true),
		Active:         entity.BoolPtr(true),
		IsCount:        entity.BoolPtr(false),
		IsUpdateFilter: entity.BoolPtr(true),
	}

	tagTypes, err = s.List(ctx, tagTypeFilter)
	if err != nil {
		s.defaultForList = nil
		return nil
	}

	// tag order
	for i, tagType := range *tagTypeFilter.Ids {
		tagOrder[tagType] = uint32(i)
	}

	// get tag types ids
	tagTypesIds = tagTypeFilter.Ids

	s.defaultForList = &entity.DefaultTagTypes{
		TagTypes:    tagTypes,
		TagOrder:    &tagOrder,
		TagTypesIds: tagTypesIds,
	}

	return s.defaultForList
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
