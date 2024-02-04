package appInit

import (
	productService "github.com/dmRusakov/tonoco-grpc/gen/go/proto/service/v1"
	apiController "github.com/dmRusakov/tonoco/internal/controllers/api/v1"
)

func (a *App) ProductControllerGetterInit() (err error) {
	// if already initialized
	if a.ProductController != nil {
		return nil
	}

	// product controller
	a.ProductController = apiController.NewProductController(
		productService.UnimplementedProductServiceServer{},
	)

	return nil
}
