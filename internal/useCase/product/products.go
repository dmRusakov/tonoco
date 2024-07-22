package product

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dustin/go-humanize"
)

func (uc *UseCase) GetProductList(
	ctx context.Context,
	parameters *entity.ProductsPageUrlParams,
) (*map[string]entity.ProductListItem, error) {
	// get productInfos
	productInfoFilter := entity.ProductInfoFilter{
		Page:    parameters.Page,
		PerPage: parameters.PerPage,
		IsCount: entity.BoolPtr(true),
	}
	productInfos, count, err := uc.productInfo.List(ctx, &productInfoFilter, true)
	if err != nil {
		return nil, err
	}

	fmt.Println(*count, "products:24")

	//// get currencies
	//currency, err := uc.currency.GetModel(ctx, nil, parameters.Currency)
	//if err != nil {
	//	return nil, err
	//}

	// get price types

	//// get tag_types with `list_item` type
	//tagTypeFilter := entity.TagTypeFilter{
	//	ListItem: entity.BoolPtr(true),
	//}
	//_, err = uc.tagType.List(ctx, &tagTypeFilter)
	//if err != nil {
	//	return nil, err
	//}

	//// get tags
	//tagFilter := entity.TagFilter{
	//	ProductIds: productInfoFilter.Ids,
	//	TagTypeIds: tagTypeFilter.Ids,
	//}
	//_, err = uc.tag.List(ctx, &tagFilter)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// get tag selects
	//tagSelectFilter := entity.TagSelectFilter{
	//	Ids:        tagFilter.TagSelectIds,
	//	TagTypeIds: tagTypeFilter.Ids,
	//}
	//_, err = uc.tagSelect.List(ctx, &tagSelectFilter)
	//if err != nil {
	//	return nil, err
	//}

	// dto
	productsDto := make(map[string]entity.ProductListItem)
	counter := 0
	for id, product := range *productInfos {
		productsDto[id] = entity.ProductListItem{
			Id:               product.Id,
			Sku:              product.Sku,
			Quantity:         1,
			Name:             product.Name,
			ShortDescription: product.ShortDescription,
			Url:              product.Url,
			//Currency:         currency.Symbol,
			Price:          humanize.FormatFloat("#,###.##", 2195),
			IsTaxable:      product.IsTaxable,
			SeoTitle:       product.SeoTitle,
			SeoDescription: product.SeoDescription,
		}

		counter++
	}

	return &productsDto, nil
}
