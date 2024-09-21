package appInit

import (
	"github.com/dmRusakov/tonoco/internal/config"
	currency_model "github.com/dmRusakov/tonoco/internal/domain/currency/model"
	currency_service "github.com/dmRusakov/tonoco/internal/domain/currency/service"
	file_model "github.com/dmRusakov/tonoco/internal/domain/file/model"
	file_service "github.com/dmRusakov/tonoco/internal/domain/file/service"
	folder_model "github.com/dmRusakov/tonoco/internal/domain/folder/model"
	folder_service "github.com/dmRusakov/tonoco/internal/domain/folder/service"
	image_model "github.com/dmRusakov/tonoco/internal/domain/image/model"
	image_service "github.com/dmRusakov/tonoco/internal/domain/image/service"
	price_model "github.com/dmRusakov/tonoco/internal/domain/price/model"
	price_service "github.com/dmRusakov/tonoco/internal/domain/price/service"
	price_type_model "github.com/dmRusakov/tonoco/internal/domain/price_type/model"
	price_type_service "github.com/dmRusakov/tonoco/internal/domain/price_type/service"
	product_image_model "github.com/dmRusakov/tonoco/internal/domain/product_image/model"
	product_image_service "github.com/dmRusakov/tonoco/internal/domain/product_image/service"
	product_info_model "github.com/dmRusakov/tonoco/internal/domain/product_info/model"
	product_info_service "github.com/dmRusakov/tonoco/internal/domain/product_info/service"
	stock_quantity_model "github.com/dmRusakov/tonoco/internal/domain/stock_quantity/model"
	stock_quantity_service "github.com/dmRusakov/tonoco/internal/domain/stock_quantity/service"
	store_model "github.com/dmRusakov/tonoco/internal/domain/store/model"
	store_service "github.com/dmRusakov/tonoco/internal/domain/store/service"
	specification_model "github.com/dmRusakov/tonoco/internal/domain/tag/model"
	specification_service "github.com/dmRusakov/tonoco/internal/domain/tag/service"
	specification_value_model "github.com/dmRusakov/tonoco/internal/domain/tag_select/model"
	specification_value_service "github.com/dmRusakov/tonoco/internal/domain/tag_select/service"
	specification_type_model "github.com/dmRusakov/tonoco/internal/domain/tag_type/model"
	specification_type_service "github.com/dmRusakov/tonoco/internal/domain/tag_type/service"
	warehouse_model "github.com/dmRusakov/tonoco/internal/domain/warehouse/model"
	warehouse_service "github.com/dmRusakov/tonoco/internal/domain/warehouse/service"

	productPolicy "github.com/dmRusakov/tonoco/internal/useCase/product"
)

func (a *App) ProductUseCaseInit(cfg *config.Config) error {
	// currency
	currencyStorage := currency_model.NewStorage(a.SqlDB)
	currencyService := currency_service.NewService(currencyStorage)

	// file
	fileStorage := file_model.NewStorage(a.SqlDB)
	fileService := file_service.NewService(fileStorage)

	// folder
	folderStorage := folder_model.NewStorage(a.SqlDB)
	folderService := folder_service.NewService(folderStorage)

	// price
	priceStorage := price_model.NewStorage(a.SqlDB)
	priceService := price_service.NewService(priceStorage)

	// price type
	priceTypeStorage := price_type_model.NewStorage(a.SqlDB)
	priceTypeService := price_type_service.NewService(priceTypeStorage)

	// tag
	specificationStorage := specification_model.NewStorage(a.SqlDB)
	specificationService := specification_service.NewService(specificationStorage)

	// tag type
	specificationTypeStorage := specification_type_model.NewStorage(a.SqlDB)
	specificationTypeService := specification_type_service.NewService(specificationTypeStorage)

	// tag value
	specificationValueStorage := specification_value_model.NewStorage(a.SqlDB)
	specificationValueService := specification_value_service.NewService(specificationValueStorage)

	// product
	productInfoStorage := product_info_model.NewStorage(a.SqlDB)
	productInfoService := product_info_service.NewService(productInfoStorage)
	// warehouse
	warehouseStorage := warehouse_model.NewStorage(a.SqlDB)
	warehouseService := warehouse_service.NewService(warehouseStorage)

	// store
	storeStorage := store_model.NewStorage(a.SqlDB)
	storeService, err := store_service.NewService(storeStorage, cfg)
	if err != nil {
		return err
	}

	// stoke quantity
	stockQuantityStorage := stock_quantity_model.NewStorage(a.SqlDB)
	stockQuantityService := stock_quantity_service.NewService(stockQuantityStorage)

	// image
	imageStorage := image_model.NewStorage(a.SqlDB)
	imageService := image_service.NewService(imageStorage)

	// product_image
	productImageStorage := product_image_model.NewStorage(a.SqlDB)
	productImageService := product_image_service.NewService(productImageStorage)

	a.ProductUseCase = productPolicy.NewProductUseCase(
		a.generator,
		a.clock,

		currencyService,
		fileService,
		folderService,
		priceService,
		priceTypeService,
		productInfoService,
		specificationService,
		specificationTypeService,
		specificationValueService,
		stockQuantityService,
		storeService,
		warehouseService,
		imageService,
		productImageService,
	)

	return nil
}
