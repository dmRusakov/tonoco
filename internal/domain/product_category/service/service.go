package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	"time"
)

type repository interface {
	All(ctx context.Context, filter *Filter) ([]*ProductCategory, error)
	Create(ctx context.Context, productCategory *ProductCategory, by string) (*model.ProductCategory, error)
	Get(ctx context.Context, id string) (*ProductCategory, error)
	GetByURL(ctx context.Context, url string) (*ProductCategory, error)
	Update(ctx context.Context, productCategory *ProductCategory, by string) (*ProductCategory, error)
	Patch(ctx context.Context, id string, fields map[string]interface{}, by string) (*ProductCategory, error)
	UpdatedAt(ctx context.Context, id string) (time.Time, error)
	TableUpdated(ctx context.Context) (time.Time, error)
	MaxSortOrder(ctx context.Context) (*uint32, error)
	Delete(ctx context.Context, id string) error
}

type ProductCategoryService struct {
	repository model.ProductCategoryStorage
}

func NewProductCategoryService(repository model.ProductCategoryStorage) *ProductCategoryService {
	return &ProductCategoryService{repository: repository}
}
