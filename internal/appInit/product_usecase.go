package appInit

import (
	product_usecase "github.com/dmRusakov/tonoco/internal/useCase/product"
)

func (a *App) ProductUseCaseInit() error {
	a.ProductUseCase = product_usecase.NewUseCase(
		a.Services.Currency,
		a.Services.File,
		a.Services.Folder,
		a.Services.Price,
		a.Services.PriceType,
		a.Services.ProductInfo,
		a.Services.Tag,
		a.Services.TagType,
		a.Services.TagSelect,
		a.Services.StockQuantity,
		a.Services.Store,
		a.Services.Warehouse,
		a.Services.ImageService,
		a.Services.ProductImageService,
	)

	return nil
}
