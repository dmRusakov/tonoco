package appInit

import (
	productPolicy "github.com/dmRusakov/tonoco/internal/domain/policy/product"
	"github.com/dmRusakov/tonoco/internal/domain/product/model"
	"github.com/dmRusakov/tonoco/internal/domain/product/service"
)

func (a *App) ProductPolicyInit() (err error) {
	productStorage := model.NewProductStorage(a.sqlDB)
	productService := service.NewProductService(productStorage)
	a.ProductPolicy = productPolicy.NewProductPolicy(productService, a.generator, a.clock)

	return nil
}
