package product

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/entity"
)

func (uc *UseCase) GetProductList(
	ctx context.Context,
	parameters *entity.ProductsUrlParameters,
) ([]entity.ProductListItem, error) {
	// get products
	products, productIds, err := uc.productInfo.List(ctx, &entity.ProductInfoFilter{})
	if err != nil {
		return nil, err
	}

	// get tag_types with `list_item` type
	listItem := true
	tagTypes, tagTypeIds, err := uc.tagType.List(ctx, &entity.TagTypeFilter{ListItem: &listItem})
	if err != nil {
		return nil, err
	}

	// get tags
	tag, err := uc.tag.List(ctx, &entity.TagFilter{
		ProductIDs: productIds,
		TagTypeIDs: tagTypeIds,
	})

	fmt.Println(tagTypes, "products:26")
	fmt.Println(tag, "products:27")
	fmt.Println(productIds, "products:28")

	// dto
	var productsDto []entity.ProductListItem
	for _, product := range *products {
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
