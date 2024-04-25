package appInit

import (
	currency_model "github.com/dmRusakov/tonoco/internal/domain/currency/model"
	currency_service "github.com/dmRusakov/tonoco/internal/domain/currency/service"
	file_model "github.com/dmRusakov/tonoco/internal/domain/file/model"
	file_service "github.com/dmRusakov/tonoco/internal/domain/file/service"
	folder_model "github.com/dmRusakov/tonoco/internal/domain/folder/model"
	folder_service "github.com/dmRusakov/tonoco/internal/domain/folder/service"
	price_model "github.com/dmRusakov/tonoco/internal/domain/price/model"
	price_service "github.com/dmRusakov/tonoco/internal/domain/price/service"
	product_category_model "github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	product_category_service "github.com/dmRusakov/tonoco/internal/domain/product_category/service"
	product_info_model "github.com/dmRusakov/tonoco/internal/domain/product_info/model"
	product_info_service "github.com/dmRusakov/tonoco/internal/domain/product_info/service"
	product_status_model "github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	product_status_service "github.com/dmRusakov/tonoco/internal/domain/product_status/service"
	shipping_class_model "github.com/dmRusakov/tonoco/internal/domain/shipping_class/model"
	shipping_class_service "github.com/dmRusakov/tonoco/internal/domain/shipping_class/service"
	specification_model "github.com/dmRusakov/tonoco/internal/domain/specification/model"
	specification_service "github.com/dmRusakov/tonoco/internal/domain/specification/service"
	specification_type_model "github.com/dmRusakov/tonoco/internal/domain/specification_type/model"
	specification_type_service "github.com/dmRusakov/tonoco/internal/domain/specification_type/service"
	specification_value_model "github.com/dmRusakov/tonoco/internal/domain/specification_value/model"
	specification_value_service "github.com/dmRusakov/tonoco/internal/domain/specification_value/service"

	productPolicy "github.com/dmRusakov/tonoco/internal/domain/useCase/product"
)

func (a *App) ProductUseCaseInit() (err error) {
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

	// product status
	productStatusStorage := product_status_model.NewStorage(a.SqlDB)
	productStatusService := product_status_service.NewService(productStatusStorage)

	// product category
	productCategoryStorage := product_category_model.NewStorage(a.SqlDB)
	productCategoryService := product_category_service.NewService(productCategoryStorage)

	// shipping class
	shippingClassStorage := shipping_class_model.NewStorage(a.SqlDB)
	shippingClassService := shipping_class_service.NewService(shippingClassStorage)

	// specification
	specificationStorage := specification_model.NewStorage(a.SqlDB)
	specificationService := specification_service.NewService(specificationStorage)

	// specification type
	specificationTypeStorage := specification_type_model.NewStorage(a.SqlDB)
	specificationTypeService := specification_type_service.NewService(specificationTypeStorage)

	// specification value
	specificationValueStorage := specification_value_model.NewStorage(a.SqlDB)
	specificationValueService := specification_value_service.NewService(specificationValueStorage)

	// product
	productInfoStorage := product_info_model.NewStorage(a.SqlDB)
	productInfoService := product_info_service.NewService(productInfoStorage)

	a.ProductUseCase = productPolicy.NewProductUseCase(
		a.generator,
		a.clock,
		currencyService,
		fileService,
		folderService,
		priceService,
		productStatusService,
		productCategoryService,
		shippingClassService,
		specificationService,
		specificationTypeService,
		specificationValueService,
		productInfoService,
	)

	return nil
}
