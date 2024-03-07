package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	"time"
)

func (s *ProductCategoryService) All(ctx context.Context, filter *model.Filter) ([]*model.ProductCategory, error) {
	return s.repository.All(ctx, filter)
}

func (s *ProductCategoryService) Create(ctx context.Context, productCategory *model.ProductCategory, by string) (*model.ProductCategory, error) {
	return s.repository.Create(ctx, productCategory, by)
}

func (s *ProductCategoryService) Get(ctx context.Context, id string) (*model.ProductCategory, error) {
	return s.repository.Get(ctx, id)
}

func (s *ProductCategoryService) GetByURL(ctx context.Context, url string) (*model.ProductCategory, error) {
	return s.repository.GetByURL(ctx, url)
}

func (s *ProductCategoryService) Update(ctx context.Context, productCategory *model.ProductCategory, by string) (*model.ProductCategory, error) {
	return s.repository.Update(ctx, productCategory, by)
}

func (s *ProductCategoryService) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.ProductCategory, error) {
	return s.repository.Patch(ctx, id, fields, by)
}

func (s *ProductCategoryService) UpdatedAt(ctx context.Context, id string) (time.Time, error) {
	return s.repository.UpdatedAt(ctx, id)
}

func (s *ProductCategoryService) TableUpdated(ctx context.Context) (time.Time, error) {
	return s.repository.TableUpdated(ctx)
}

func (s *ProductCategoryService) MaxSortOrder(ctx context.Context) (*uint32, error) {
	return s.repository.MaxSortOrder(ctx)
}

func (s *ProductCategoryService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
