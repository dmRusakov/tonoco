package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"time"
)

type repository interface {
	All(ctx context.Context, filter *Filter) ([]*model.ProductStatus, error)
	Create(ctx context.Context, productStatus *model.ProductStatus, by string) (*model.ProductStatus, error)
	Get(ctx context.Context, id string) (*model.ProductStatus, error)
	GetByURL(ctx context.Context, url string) (*model.ProductStatus, error)
	Update(ctx context.Context, productStatus *model.ProductStatus) (*model.ProductStatus, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*model.ProductStatus, error)
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
