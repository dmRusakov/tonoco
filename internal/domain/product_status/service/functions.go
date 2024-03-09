package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"time"
)

func (s *ProductStatusService) All(ctx context.Context, filter *model.Filter) ([]*model.ProductStatus, error) {
	return s.repository.All(ctx, filter)
}

func (s *ProductStatusService) Create(ctx context.Context, productStatus *model.ProductStatus, by string) (*model.ProductStatus, error) {
	return s.repository.Create(ctx, productStatus)
}

func (s *ProductStatusService) Get(ctx context.Context, id string) (*model.ProductStatus, error) {
	return s.repository.Get(ctx, id)
}

func (s *ProductStatusService) GetByURL(ctx context.Context, url string) (*model.ProductStatus, error) {
	return s.repository.GetByURL(ctx, url)
}

func (s *ProductStatusService) Update(ctx context.Context, productStatus *model.ProductStatus, by string) (*model.ProductStatus, error) {
	return s.repository.Update(ctx, productStatus)
}

func (s *ProductStatusService) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.ProductStatus, error) {
	return s.repository.Patch(ctx, id, fields)
}

func (s *ProductStatusService) UpdatedAt(ctx context.Context, id string) (time.Time, error) {
	return s.repository.UpdatedAt(ctx, id)
}

func (s *ProductStatusService) TableUpdated(ctx context.Context) (time.Time, error) {
	return s.repository.TableUpdated(ctx)
}

func (s *ProductStatusService) MaxSortOrder(ctx context.Context) (*uint32, error) {
	return s.repository.MaxSortOrder(ctx)
}

func (s *ProductStatusService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
