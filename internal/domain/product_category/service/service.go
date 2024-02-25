package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_category/model"
)

type repository interface {
	Create(ctx context.Context, productCategory *model.ProductCategory, by string) (*model.ProductCategory, error)
	Get(ctx context.Context, id string) (*model.ProductCategory, error)
	Update(ctx context.Context, productCategory *model.ProductCategory, by string) (*model.ProductCategory, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.ProductCategory, error)
	Delete(ctx context.Context, id string) error
}

type ProductCategoryService struct {
	repository model.ProductCategoryStorage
}

func NewProductCategoryService(repository model.ProductCategoryStorage) *ProductCategoryService {
	return &ProductCategoryService{repository: repository}
}

func (s *ProductCategoryService) Create(ctx context.Context, productCategory *model.ProductCategory, by string) (*model.ProductCategory, error) {
	return s.repository.Create(ctx, productCategory, by)
}

func (s *ProductCategoryService) Get(ctx context.Context, id string) (*model.ProductCategory, error) {
	return s.repository.Get(ctx, id)
}

func (s *ProductCategoryService) Update(ctx context.Context, productCategory *model.ProductCategory, by string) (*model.ProductCategory, error) {
	return s.repository.Update(ctx, productCategory, by)
}

func (s *ProductCategoryService) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.ProductCategory, error) {
	return s.repository.Patch(ctx, id, fields, by)
}

func (s *ProductCategoryService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
