package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	"time"
)

func (s *Service) All(ctx context.Context, filter *entity.ProductStatusFilter) ([]*entity.ProductStatus, error) {
	return s.repository.All(ctx, filter)
}

func (s *Service) Create(ctx context.Context, productStatus *entity.ProductStatus, by string) (*entity.ProductStatus, error) {
	return s.repository.Create(ctx, productStatus)
}

func (s *Service) Get(ctx context.Context, id string) (*entity.ProductStatus, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) GetByURL(ctx context.Context, url string) (*entity.ProductStatus, error) {
	return s.repository.GetByURL(ctx, url)
}

func (s *Service) Update(ctx context.Context, productStatus *entity.ProductStatus, by string) (*entity.ProductStatus, error) {
	return s.repository.Update(ctx, productStatus)
}

func (s *Service) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*entity.ProductStatus, error) {
	return s.repository.Patch(ctx, id, fields)
}

func (s *Service) UpdatedAt(ctx context.Context, id string) (time.Time, error) {
	return s.repository.UpdatedAt(ctx, id)
}

func (s *Service) TableUpdated(ctx context.Context) (time.Time, error) {
	return s.repository.TableUpdated(ctx)
}

func (s *Service) MaxSortOrder(ctx context.Context) (*uint32, error) {
	return s.repository.MaxSortOrder(ctx)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
