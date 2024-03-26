package product

import (
	product_service "github.com/dmRusakov/tonoco/internal/domain/product/service"
	product_category_service "github.com/dmRusakov/tonoco/internal/domain/product_category/service"
	product_status_service "github.com/dmRusakov/tonoco/internal/domain/product_status/service"
	shipping_class_service "github.com/dmRusakov/tonoco/internal/domain/shipping_class/service"
	specification_service "github.com/dmRusakov/tonoco/internal/domain/specification/service"
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
	identity IdentityGenerator
	clock    Clock

	productService         *product_service.Service
	productCategoryService *product_category_service.Service
	productStatusService   *product_status_service.Service
	shippingClassService   *shipping_class_service.Service
	specificationService   *specification_service.Service
}

func NewProductUseCase(
	identity IdentityGenerator,
	clock clock.Clock,
	productService *product_service.Service,
	productCategoryService *product_category_service.Service,
	productStatusService *product_status_service.Service,
	shippingClassService *shipping_class_service.Service,
	specificationService *specification_service.Service,
) *UseCase {
	return &UseCase{
		productService:         productService,
		productCategoryService: productCategoryService,
		productStatusService:   productStatusService,
		shippingClassService:   shippingClassService,
		specificationService:   specificationService,

		identity: identity,
		clock:    clock,
	}
}
