package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	"github.com/dmRusakov/tonoco/internal/domain/file/model"
	"time"
)

type Item = entity.File
type Filter = entity.FileFilter

type repository interface {
	Get(ctx context.Context, id *string, url *string) (*Item, []error)
	List(context.Context, *Filter) ([]*Item, []error)
	Create(context.Context, *Item) (*string, []error)
	Update(context.Context, *Item) []error
	Patch(context.Context, *string, *map[string]interface{}) []error
	UpdatedAt(context.Context, *string) (*time.Time, []error)
	TableIndexCount(context.Context) (*uint64, []error)
	MaxSortOrder(context.Context) (*uint64, []error)
	Delete(context.Context, *string) []error
}

type Service struct {
	repository model.Storage
}

func NewService(repository model.Storage) *Service {
	return &Service{repository: repository}
}

func (s *Service) Get(ctx context.Context, id *string, url *string) (*Item, error) {
	return s.repository.Get(ctx, id, url)
}

func (s *Service) List(ctx context.Context, filter *Filter) ([]*Item, error) {
	return s.repository.List(ctx, filter)
}

func (s *Service) Create(ctx context.Context, item *Item) (*string, error) {
	return s.repository.Create(ctx, item)
}

func (s *Service) Update(ctx context.Context, item *Item) error {
	return s.repository.Update(ctx, item)
}

func (s *Service) Patch(ctx context.Context, id *string, fields *map[string]interface{}) error {
	return s.repository.Patch(ctx, id, fields)
}

func (s *Service) UpdatedAt(ctx context.Context, id *string) (*time.Time, error) {
	return s.repository.UpdatedAt(ctx, id)
}

func (s *Service) TableIndexCount(ctx context.Context) (*uint64, error) {
	return s.repository.TableIndexCount(ctx)
}

func (s *Service) MaxSortOrder(ctx context.Context) (*uint64, error) {
	return s.repository.MaxSortOrder(ctx)
}

func (s *Service) Delete(ctx context.Context, id *string) error {
	return s.repository.Delete(ctx, id)
}
