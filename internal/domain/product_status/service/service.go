package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
)

type repository interface {
	All(ctx context.Context, filter *model.Filter) ([]*model.ProductStatus, error)
	Create(ctx context.Context, productStatus *model.ProductStatus, by string) (*model.ProductStatus, error)
	Get(ctx context.Context, id string) (*model.ProductStatus, error)
	Update(ctx context.Context, productStatus *model.ProductStatus, by string) (*model.ProductStatus, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.ProductStatus, error)
	Delete(ctx context.Context, id string) error
}

type ProductStatusService struct {
	repository model.ProductStatusStorage
}

func NewProductStatusService(repository model.ProductStatusStorage) *ProductStatusService {
	return &ProductStatusService{repository: repository}
}

func (s *ProductStatusService) All(ctx context.Context, filter *model.Filter) ([]*model.ProductStatus, error) {
	return s.repository.All(ctx, filter)
}

func (s *ProductStatusService) Create(ctx context.Context, productStatus *model.ProductStatus, by string) (*model.ProductStatus, error) {
	return s.repository.Create(ctx, productStatus, by)
}

func (s *ProductStatusService) Get(ctx context.Context, id string) (*model.ProductStatus, error) {
	return s.repository.Get(ctx, id)
}

func (s *ProductStatusService) Update(ctx context.Context, productStatus *model.ProductStatus, by string) (*model.ProductStatus, error) {
	return s.repository.Update(ctx, productStatus, by)
}

func (s *ProductStatusService) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.ProductStatus, error) {
	return s.repository.Patch(ctx, id, fields, by)
}

func (s *ProductStatusService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
