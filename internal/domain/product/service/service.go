package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product/model"
)

type repository interface {
	Create(ctx context.Context, product *model.Product, by string) (*model.Product, error)
	Get(ctx context.Context, id string) (*model.Product, error)
	Update(ctx context.Context, product *model.Product, by string) (*model.Product, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductService struct {
	repository model.ProductStorage
}

func NewProductService(repository model.ProductStorage) *ProductService {
	return &ProductService{repository: repository}
}

func (s *ProductService) Create(ctx context.Context, product *model.Product, by string) (*model.Product, error) {
	return s.repository.Create(ctx, product, by)
}

func (s *ProductService) Get(ctx context.Context, id string) (*model.Product, error) {
	return s.repository.Get(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, product *model.Product, by string) (*model.Product, error) {
	return s.repository.Update(ctx, product, by)
}

func (s *ProductService) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.Product, error) {
	return s.repository.Patch(ctx, id, fields, by)
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
