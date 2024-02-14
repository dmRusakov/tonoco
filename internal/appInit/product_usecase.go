package appInit

import (
	"github.com/dmRusakov/tonoco/internal/domain/product/model"
	"github.com/dmRusakov/tonoco/internal/domain/product/service"
	productPolicy "github.com/dmRusakov/tonoco/internal/domain/useCase/product"
)

func (a *App) ProductUseCaseInit() (err error) {
	productStorage := model.NewProductStorage(a.SqlDB)
	productService := service.NewProductService(productStorage)
	a.ProductUseCase = productPolicy.NewProductUseCase(productService, a.generator, a.clock)

	return nil
}
