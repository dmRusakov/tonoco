package appInit

import (
	"fmt"
	productPolicy "github.com/dmRusakov/tonoco/internal/domain/policy/product"
	"github.com/dmRusakov/tonoco/internal/domain/product/dao"
	"github.com/dmRusakov/tonoco/internal/domain/product/service"
)

func (a *App) ProductPolicyInit() (err error) {
	productStorage := dao.NewProductStorage(a.sqlDB)
	productService := service.NewProductService(productStorage)
	a.ProductPolicy = productPolicy.NewProductPolicy(productService, a.generator, a.clock)

	fmt.Println("2")
	return nil
}
