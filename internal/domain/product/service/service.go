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

type Service struct {
	repository model.ProductStorage
}

func NewProductService(repository model.ProductStorage) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, product *model.Product, by string) (*model.Product, error) {
	return s.repository.Create(ctx, product, by)
}

func (s *Service) Get(ctx context.Context, id string) (*model.Product, error) {
	return s.repository.Get(ctx, id)
}

func (s *Service) Update(ctx context.Context, product *model.Product, by string) (*model.Product, error) {
	return s.repository.Update(ctx, product, by)
}

func (s *Service) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.Product, error) {
	return s.repository.Patch(ctx, id, fields, by)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
