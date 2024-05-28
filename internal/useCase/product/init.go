package product

import (
	currency_service "github.com/dmRusakov/tonoco/internal/domain/currency/service"
	file_service "github.com/dmRusakov/tonoco/internal/domain/file/service"
	folder_service "github.com/dmRusakov/tonoco/internal/domain/folder/service"
	price_service "github.com/dmRusakov/tonoco/internal/domain/price/service"
	price_type_service "github.com/dmRusakov/tonoco/internal/domain/price_type/service"
	product_info_service "github.com/dmRusakov/tonoco/internal/domain/product_info/service"
	store_service "github.com/dmRusakov/tonoco/internal/domain/store/service"
	specification_service "github.com/dmRusakov/tonoco/internal/domain/tag/service"
	specification_value_service "github.com/dmRusakov/tonoco/internal/domain/tag_select/service"
	specification_type_service "github.com/dmRusakov/tonoco/internal/domain/tag_type/service"
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
	productInfo        *product_info_service.Service
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
	productInfoService *product_info_service.Service,
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
		productInfo:        productInfoService,
		specification:      specificationService,
		specificationType:  specificationTypeService,
		specificationValue: specificationValueService,
		store:              storeService,
		warehouse:          warehouseService,
	}
}
