package product

import (
	currency_service "github.com/dmRusakov/tonoco/internal/domain/currency/service"
	file_service "github.com/dmRusakov/tonoco/internal/domain/file/service"
	folder_service "github.com/dmRusakov/tonoco/internal/domain/folder/service"
	price_service "github.com/dmRusakov/tonoco/internal/domain/price/service"
	price_type_service "github.com/dmRusakov/tonoco/internal/domain/price_type/service"
	product_category_service "github.com/dmRusakov/tonoco/internal/domain/product_category/service"
	product_info_service "github.com/dmRusakov/tonoco/internal/domain/product_info/service"
	product_status_service "github.com/dmRusakov/tonoco/internal/domain/product_status/service"
	shipping_class_service "github.com/dmRusakov/tonoco/internal/domain/shipping_class/service"
	specification_service "github.com/dmRusakov/tonoco/internal/domain/specification/service"
	specification_type_service "github.com/dmRusakov/tonoco/internal/domain/specification_type/service"
	specification_value_service "github.com/dmRusakov/tonoco/internal/domain/specification_value/service"
	store_service "github.com/dmRusakov/tonoco/internal/domain/store/service"
	warehouse_service "github.com/dmRusakov/tonoco/internal/domain/warehouse/service"
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

	currency           *currency_service.Service
	file               *file_service.Service
	folder             *folder_service.Service
	price              *price_service.Service
	priceType          *price_type_service.Service
	productCategory    *product_category_service.Service
	productInfo        *product_info_service.Service
	productStatus      *product_status_service.Service
	shippingClass      *shipping_class_service.Service
	specification      *specification_service.Service
	specificationType  *specification_type_service.Service
	specificationValue *specification_value_service.Service
	store              *store_service.Service
	warehouse          *warehouse_service.Service
}

func NewProductUseCase(
	identity IdentityGenerator,
	clock clock.Clock,

	currencyService *currency_service.Service,
	fileService *file_service.Service,
	folderService *folder_service.Service,
	priceService *price_service.Service,
	priceTypeService *price_type_service.Service,
	productCategoryService *product_category_service.Service,
	productInfoService *product_info_service.Service,
	productStatusService *product_status_service.Service,
	shippingClassService *shipping_class_service.Service,
	specificationService *specification_service.Service,
	specificationTypeService *specification_type_service.Service,
	specificationValueService *specification_value_service.Service,
	storeService *store_service.Service,
	warehouseService *warehouse_service.Service,
) *UseCase {
	return &UseCase{
		identity: identity,
		clock:    clock,

		currency:           currencyService,
		file:               fileService,
		folder:             folderService,
		price:              priceService,
		priceType:          priceTypeService,
		productCategory:    productCategoryService,
		productInfo:        productInfoService,
		productStatus:      productStatusService,
		shippingClass:      shippingClassService,
		specification:      specificationService,
		specificationType:  specificationTypeService,
		specificationValue: specificationValueService,
		store:              storeService,
		warehouse:          warehouseService,
	}
}
