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
	page := uint64(2)
	perPage := uint64(10)
	productFilter := entity.ProductInfoFilter{
		Page:    &page,
		PerPage: &perPage,
	}
	products, err := uc.productInfo.List(ctx, &productFilter)
	if err != nil {
		fmt.Println(err, "products:17")
		return nil, err
	}

	fmt.Println(products, "products:20")

	// get tag_types with `list_item` type
	listItem := true
	_, tagTypeIds, err := uc.tagType.List(ctx, &entity.TagTypeFilter{
		ListItem: &listItem,
	})
	if err != nil {
		return nil, err
	}

	// get tags
	tagFilter := entity.TagFilter{
		ProductIDs: productFilter.IDs,
		TagTypeIDs: tagTypeIds,
	}
	_, err = uc.tag.List(ctx, &tagFilter)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\nproducts:38\n", tagFilter.TagTypeIDs)

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
