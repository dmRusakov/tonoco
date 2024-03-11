package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	"github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	"time"
)

type repository interface {
	All(ctx context.Context, filter *entity.ProductStatusFilter) ([]*entity.ProductStatus, error)
	Create(ctx context.Context, productStatus *entity.ProductStatus, by string) (*entity.ProductStatus, error)
	Get(ctx context.Context, id string) (*entity.ProductStatus, error)
	GetByURL(ctx context.Context, url string) (*entity.ProductStatus, error)
	Update(ctx context.Context, productStatus *entity.ProductStatus) (*entity.ProductStatus, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*entity.ProductStatus, error)
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
