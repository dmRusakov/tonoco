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
	// get productInfos
	productInfoFilter := entity.ProductInfoFilter{
		Page:    entity.Uint64Ptr(2),
		PerPage: entity.Uint64Ptr(10),
	}
	productInfos, err := uc.productInfo.List(ctx, &productInfoFilter)
	if err != nil {
		return nil, err
	}

	// get tag_types with `list_item` type
	tagTypeFilter := entity.TagTypeFilter{
		ListItem: entity.BoolPtr(true),
	}
	_, err = uc.tagType.List(ctx, &tagTypeFilter)
	if err != nil {
		return nil, err
	}

	// get tags
	tagFilter := entity.TagFilter{
		ProductIDs: productInfoFilter.IDs,
		TagTypeIDs: tagTypeFilter.IDs,
	}
	_, err = uc.tag.List(ctx, &tagFilter)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\nproductInfos:38\n", tagFilter.TagTypeIDs)

	// dto
	var productsDto []entity.ProductListItem
	for _, product := range *productInfos {
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
