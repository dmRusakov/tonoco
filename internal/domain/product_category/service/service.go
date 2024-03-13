package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	"github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	"time"
)

type repository interface {
	All(ctx context.Context, filter *entity.ProductCategoryFilter) ([]*entity.ProductCategory, error)
	Create(ctx context.Context, productCategory *entity.ProductCategory) (*entity.ProductCategory, error)
	Get(ctx context.Context, id string) (*entity.ProductCategory, error)
	GetByURL(ctx context.Context, url string) (*entity.ProductCategory, error)
	Update(ctx context.Context, productCategory *entity.ProductCategory) (*entity.ProductCategory, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}) (*entity.ProductCategory, error)
	UpdatedAt(ctx context.Context, id string) (time.Time, error)
	TableUpdated(ctx context.Context) (time.Time, error)
	MaxSortOrder(ctx context.Context) (*uint32, error)
	Delete(ctx context.Context, id string) error
}

type Service struct {
	repository model.Storage
}

func NewService(repository model.Storage) *Service {
	return &Service{repository: repository}
}
