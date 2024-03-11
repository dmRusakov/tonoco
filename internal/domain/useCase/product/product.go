package product

import (
	product_service "github.com/dmRusakov/tonoco/internal/domain/product/service"
	product_category_service "github.com/dmRusakov/tonoco/internal/domain/product_category/service"
	"github.com/dmRusakov/tonoco/pkg/common/core/clock"
	"time"
)

type IdentityGenerator interface {
	GenerateUUIDv4String() string
}

type Clock interface {
	Now() time.Time
}

type UseCase struct {
	productService         *product_service.ProductService
	productCategoryService *product_category_service.ProductCategoryService

	identity IdentityGenerator
	clock    Clock
}

func NewProductUseCase(
	productService *product_service.ProductService,
	productCategoryService *product_category_service.ProductCategoryService,
	identity IdentityGenerator,
	clock clock.Clock,
) *UseCase {
	return &UseCase{
		productService:         productService,
		productCategoryService: productCategoryService,

		identity: identity,
		clock:    clock,
	}
}
