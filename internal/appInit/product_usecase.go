package appInit

import (
	product_model "github.com/dmRusakov/tonoco/internal/domain/product/model"
	product_service "github.com/dmRusakov/tonoco/internal/domain/product/service"
	product_category_model "github.com/dmRusakov/tonoco/internal/domain/product_category/model"
	product_category_service "github.com/dmRusakov/tonoco/internal/domain/product_category/service"
	productPolicy "github.com/dmRusakov/tonoco/internal/domain/useCase/product"
)

func (a *App) ProductUseCaseInit() (err error) {
	// product
	productStorage := product_model.NewProductStorage(a.SqlDB)
	productService := product_service.NewProductService(productStorage)

	// product category
	productCategoryStorage := product_category_model.NewProductCategoryStorage(a.SqlDB)
	productCategoryService := product_category_service.NewProductCategoryService(productCategoryStorage)

	a.ProductUseCase = productPolicy.NewProductUseCase(productService, productCategoryService, a.generator, a.clock)

	return nil
}
