package appInit

import (
	policy_product "github.com/dmRusakov/tonoco/internal/domain/policy/product"
	"github.com/dmRusakov/tonoco/internal/domain/product/dao"
	"github.com/dmRusakov/tonoco/internal/domain/product/service"
)

// ProductPolicyInit - product policies initialization
func (a *App) productPolicyInit() (err error) {
	// check sqlDB
	if a.sqlDB == nil {
		err = a.productDBInit()
		if err != nil {
			return err
		}
	}

	// check generator
	if a.generator == nil {
		err = a.generatorInit()
		if err != nil {
			return err
		}
	}

	// check clock
	if a.clock == nil {
		err = a.clockInit()
		if err != nil {
			return err
		}
	}

	// init product storage, service and policy_product
	productStorage := dao.NewProductStorage(a.sqlDB)
	productService := service.NewProductService(productStorage)
	a.productPolicy = policy_product.NewProductPolicy(productService, a.generator, a.clock)

	return nil
}

//
//// ProductAPIInit - product API initialization
//func (a *App) ProductAPIInit() (err error) {
//	// check product policy_product
//	if a.productPolicy == nil {
//		err = a.productPolicyInit()
//		if err != nil {
//			return err
//		}
//	}
//
//	a.ProductApi = grpc_v1_product.NewServer(
//		a.productPolicy,
//		pb_prod_products.UnimplementedTonocoProductServiceServer{},
//	)
//
//	return nil
//}
