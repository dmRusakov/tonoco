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

	currencyService           *currency_service.Service
	fileService               *file_service.Service
	folderService             *folder_service.Service
	priceService              *price_service.Service
	priceTypeService          *price_type_service.Service
	productStatusService      *product_status_service.Service
	productCategoryService    *product_category_service.Service
	shippingClassService      *shipping_class_service.Service
	specificationService      *specification_service.Service
	specificationTypeService  *specification_type_service.Service
	specificationValueService *specification_value_service.Service
	productInfoService        *product_info_service.Service
	warehouseService          *warehouse_service.Service
	storeService              *store_service.Service
}

func NewProductUseCase(
	identity IdentityGenerator,
	clock clock.Clock,
	currencyService *currency_service.Service,
	fileService *file_service.Service,
	folderService *folder_service.Service,
	priceService *price_service.Service,
	priceTypeService *price_type_service.Service,
	productStatusService *product_status_service.Service,
	productCategoryService *product_category_service.Service,
	shippingClassService *shipping_class_service.Service,
	specificationService *specification_service.Service,
	specificationTypeService *specification_type_service.Service,
	specificationValueService *specification_value_service.Service,
	productInfoService *product_info_service.Service,
	warehouseService *warehouse_service.Service,
	storeService *store_service.Service,
) *UseCase {
	return &UseCase{
		currencyService:           currencyService,
		fileService:               fileService,
		folderService:             folderService,
		priceService:              priceService,
		priceTypeService:          priceTypeService,
		productStatusService:      productStatusService,
		productCategoryService:    productCategoryService,
		shippingClassService:      shippingClassService,
		specificationService:      specificationService,
		specificationTypeService:  specificationTypeService,
		specificationValueService: specificationValueService,
		productInfoService:        productInfoService,
		warehouseService:          warehouseService,
		storeService:              storeService,

		identity: identity,
		clock:    clock,
	}
}
