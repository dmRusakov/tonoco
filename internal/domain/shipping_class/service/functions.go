package service

import (
	"context"
	"time"
)

func (s *Service) List(ctx context.Context, filter *Filter) ([]*Item, error) {
	return s.repository.List(ctx, filter)
}

func (s *Service) Create(ctx context.Context, productStatus *Item) (*Item, error) {
	return s.repository.Create(ctx, productStatus)
}

func (s *Service) Get(ctx context.Context, id *string, url *string) (*Item, error) {
	return s.repository.Get(ctx, id, url)
}

func (s *Service) Update(ctx context.Context, productStatus *Item) (*Item, error) {
	return s.repository.Update(ctx, productStatus)
}

func (s *Service) Patch(ctx context.Context, id *string, fields *map[string]interface{}) (*Item, error) {
	return s.repository.Patch(ctx, id, fields)
}

func (s *Service) UpdatedAt(ctx context.Context, id *string) (*time.Time, error) {
	return s.repository.UpdatedAt(ctx, id)
}

func (s *Service) TableUpdated(ctx context.Context) (*time.Time, error) {
	return s.repository.TableUpdated(ctx)
}

func (s *Service) MaxSortOrder(ctx context.Context) (*uint32, error) {
	return s.repository.MaxSortOrder(ctx)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
