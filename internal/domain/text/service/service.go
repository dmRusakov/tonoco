package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/text/model"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/google/uuid"
	"time"
)

type Item = db.Text
type Filter = db.TextFilter

type Service struct {
	repository model.Storage
}

type Repository interface {
	Get(context.Context, *Filter) (*Item, error)
	List(context.Context, *Filter) (*map[uuid.UUID]Item, error)
	Ids(context.Context, *Filter) (*[]uuid.UUID, error)
	Create(context.Context, *Item) (*uuid.UUID, error)
	Update(context.Context, *Item) error
	Patch(context.Context, *uuid.UUID, *map[string]interface{}) error
	UpdatedAt(context.Context, *uuid.UUID) (*time.Time, error)
	MaxSortOrder(context.Context) (*uint64, error)
	TableIndexCount(context.Context) (*uint64, error)
	Delete(context.Context, *uuid.UUID) error
}

func NewService(repository *model.Model) *Service {
	return &Service{
		repository: repository,
	}
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