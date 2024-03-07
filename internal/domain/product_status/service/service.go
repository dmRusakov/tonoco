package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"time"
)

type repository interface {
	All(ctx context.Context, filter *Filter) ([]*ProductStatus, error)
	Create(ctx context.Context, productStatus *ProductStatus, by string) (*ProductStatus, error)
	Get(ctx context.Context, id string) (*ProductStatus, error)
	GetByURL(ctx context.Context, url string) (*ProductStatus, error)
	Update(ctx context.Context, productStatus *ProductStatus, by string) (*ProductStatus, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*ProductStatus, error)
	UpdatedAt(ctx context.Context, id string) (time.Time, error)
	TableUpdated(ctx context.Context) (time.Time, error)
	MaxSortOrder(ctx context.Context) (*uint32, error)
	Delete(ctx context.Context, id string) error
}

type ProductStatusService struct {
	repository model.ProductStatusStorage
}

func NewProductStatusService(repository model.ProductStatusStorage) *ProductStatusService {
	return &ProductStatusService{repository: repository}
}

func (s *ProductStatusService) All(ctx context.Context, filter *Filter) ([]*ProductStatus, error) {
	return s.repository.All(ctx, filter)
}

func (s *ProductStatusService) Create(ctx context.Context, productStatus *ProductStatus, by string) (*ProductStatus, error) {
	return s.repository.Create(ctx, productStatus, by)
}

func (s *ProductStatusService) Get(ctx context.Context, id string) (*ProductStatus, error) {
	return s.repository.Get(ctx, id)
}

func (s *ProductStatusService) GetByURL(ctx context.Context, url string) (*ProductStatus, error) {
	return s.repository.GetByURL(ctx, url)
}

func (s *ProductStatusService) Update(ctx context.Context, productStatus *ProductStatus, by string) (*ProductStatus, error) {
	return s.repository.Update(ctx, productStatus, by)
}

func (s *ProductStatusService) Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*ProductStatus, error) {
	return s.repository.Patch(ctx, id, fields, by)
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
