package product

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity"
)

func (uc *UseCase) GetProductList(
	ctx context.Context,
	parameters *entity.ProductsUrlParameters,
) ([]entity.ProductListItem, error) {
	// get products
	products, err := uc.productInfo.List(ctx, &entity.ProductInfoFilter{})
	if err != nil {
		return nil, err
	}

	// dto
	var productsDto []entity.ProductListItem
	for _, product := range products {
		productsDto = append(productsDto, entity.ProductListItem{
			ID:               product.ID,
			SKU:              product.SKU,
			Name:             product.Name,
			ShortDescription: product.ShortDescription,
			Url:              product.Url,
			IsTaxable:        product.IsTaxable,
			SeoTitle:         product.SeoTitle,
			SeoDescription:   product.SeoDescription,
		})
	}

	return productsDto, nil
}