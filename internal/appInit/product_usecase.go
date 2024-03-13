package appInit

import (
	product_model "github.com/dmRusakov/tonoco/internal/domain/product/model"
	product_service "github.com/dmRusakov/tonoco/internal/domain/product/service"
	product_category_model "github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	product_category_service "github.com/dmRusakov/tonoco/internal/domain/product_category/service"
	product_status_model "github.com/dmRusakov/tonoco/internal/domain/product_status/model"
	product_status_service "github.com/dmRusakov/tonoco/internal/domain/product_status/service"
	shipping_class_model "github.com/dmRusakov/tonoco/internal/domain/shipping_class/model"
	shipping_class_service "github.com/dmRusakov/tonoco/internal/domain/shipping_class/service"

	productPolicy "github.com/dmRusakov/tonoco/internal/domain/useCase/product"
)

func (a *App) ProductUseCaseInit() (err error) {
	// product
	productStorage := product_model.NewProductStorage(a.SqlDB)
	productService := product_service.NewProductService(productStorage)

	// product category
	productCategoryStorage := product_category_model.NewStorage(a.SqlDB)
	productCategoryService := product_category_service.NewService(productCategoryStorage)

	// product status
	productStatusStorage := product_status_model.NewStorage(a.SqlDB)
	productStatusService := product_status_service.NewService(productStatusStorage)

	// shipping class
	shippingClassStorage := shipping_class_model.NewStorage(a.SqlDB)
	shippingClassService := shipping_class_service.NewService(shippingClassStorage)

	a.ProductUseCase = productPolicy.NewProductUseCase(a.generator, a.clock, productService, productCategoryService, productStatusService, shippingClassService)

	return nil
}
