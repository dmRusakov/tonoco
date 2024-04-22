package product

import (
	file_service "github.com/dmRusakov/tonoco/internal/domain/file/service"
	folder_service "github.com/dmRusakov/tonoco/internal/domain/folder/service"
	product_category_service "github.com/dmRusakov/tonoco/internal/domain/product_category/service"
	product_status_service "github.com/dmRusakov/tonoco/internal/domain/product_status/service"
	shipping_class_service "github.com/dmRusakov/tonoco/internal/domain/shipping_class/service"
	specification_service "github.com/dmRusakov/tonoco/internal/domain/specification/service"
	specification_type_service "github.com/dmRusakov/tonoco/internal/domain/specification_type/service"
	specification_value_service "github.com/dmRusakov/tonoco/internal/domain/specification_value/service"
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
	identity                  IdentityGenerator
	clock                     Clock
	fileService               *file_service.Service
	folderService             *folder_service.Service
	productStatusService      *product_status_service.Service
	productCategoryService    *product_category_service.Service
	shippingClassService      *shipping_class_service.Service
	specificationService      *specification_service.Service
	specificationTypeService  *specification_type_service.Service
	specificationValueService *specification_value_service.Service
	//productService         *product_service.Service

}

func NewProductUseCase(
	identity IdentityGenerator,
	clock clock.Clock,
	fileService *file_service.Service,
	folderService *folder_service.Service,
	productStatusService *product_status_service.Service,
	productCategoryService *product_category_service.Service,
	shippingClassService *shipping_class_service.Service,
	specificationService *specification_service.Service,
	specificationTypeService *specification_type_service.Service,
	specificationValueService *specification_value_service.Service,
	// productService *product_service.Service,
	//

) *UseCase {
	return &UseCase{
		fileService:               fileService,
		folderService:             folderService,
		productStatusService:      productStatusService,
		productCategoryService:    productCategoryService,
		shippingClassService:      shippingClassService,
		specificationService:      specificationService,
		specificationTypeService:  specificationTypeService,
		specificationValueService: specificationValueService,

		//productService:         productService,
		identity: identity,
		clock:    clock,
	}
}
