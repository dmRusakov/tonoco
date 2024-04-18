package service

import (
	"context"
	"time"
)

func (s *Service) Get(ctx context.Context, id *string, url *string) (*Item, error) {
	return s.repository.Get(ctx, id, url)
}

func (s *Service) List(ctx context.Context, filter *Filter) ([]*Item, error) {
	return s.repository.List(ctx, filter)
}

func (s *Service) Create(ctx context.Context, productStatus *Item) (*string, error) {
	return s.repository.Create(ctx, productStatus)
}

func (s *Service) Update(ctx context.Context, productStatus *Item) error {
	return s.repository.Update(ctx, productStatus)
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
