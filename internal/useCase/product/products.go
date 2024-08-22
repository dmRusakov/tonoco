package product

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dustin/go-humanize"
)

func (uc *UseCase) GetProductList(
	ctx context.Context,
	parameters *entity.ProductsPageUrlParams,
) (*map[string]entity.ProductListItem, *uint64, error) {
	// get productInfos
	productInfoFilter := entity.ProductInfoFilter{
		Page:    parameters.Page,
		PerPage: parameters.PerPage,
		IsCount: entity.BoolPtr(true),
	}
	productInfos, totalItemsCount, err := uc.productInfo.List(ctx, &productInfoFilter, true)
	if err != nil {
		return nil, nil, err
	}

	// get currencies
	currencyInfoFilter := entity.CurrencyFilter{
		IsCount: entity.BoolPtr(true),
		Urls:    &[]string{uc.store.DefaultStore.CurrencyUrl},
	}
	currency, err := uc.currency.Get(ctx, &currencyInfoFilter)
	if err != nil {
		return nil, nil, err
	}

	// get price types
	specialPriceTypeFilter := entity.PriceTypeFilter{
		Urls:     &[]string{"special", "sale"},
		IsPublic: entity.BoolPtr(true),
		IsCount:  entity.BoolPtr(false),
	}

	_, _, err = uc.priceType.List(ctx, &specialPriceTypeFilter, true)
	if err != nil {
		return nil, nil, err
	}

	regularPriceTypeFilter := entity.PriceTypeFilter{
		Urls:     &[]string{"regular"},
		IsPublic: entity.BoolPtr(true),
		IsCount:  entity.BoolPtr(false),
	}

	_, _, err = uc.priceType.List(ctx, &regularPriceTypeFilter, true)
	if err != nil {
		return nil, nil, err
	}

	// get tag_types with `list_item` type
	tagTypeFilter := entity.TagTypeFilter{
		OrderBy:  entity.StringPtr("SortOrder"),
		OrderDir: entity.StringPtr("ASC"),
		ListItem: entity.BoolPtr(true),
		Active:   entity.BoolPtr(true),
		IsCount:  entity.BoolPtr(false),
	}

	tagTypes, _, err := uc.tagType.List(ctx, &tagTypeFilter, true)
	if err != nil {
		return nil, nil, err
	}

	// tag order
	tagOrder := make(map[string]uint32)
	for i, tagType := range *tagTypeFilter.Ids {
		tagOrder[tagType] = uint32(i)
	}

	// dto
	productsDto := make(map[string]entity.ProductListItem)
	counter := 0
	for id, product := range *productInfos {
		// get price
		specialPriceFilter := entity.PriceFilter{
			ProductIds:   &[]string{product.Id},
			PriceTypeIds: specialPriceTypeFilter.Ids,
			CurrencyIds:  &[]string{currency.Id},
			Active:       entity.BoolPtr(true),
			IsCount:      entity.BoolPtr(false),
		}
		specialPrice, _, err := uc.price.List(ctx, &specialPriceFilter, true)
		if err != nil {
			return nil, nil, err
		}

		// get regularPrice
		regularPriceFilter := entity.PriceFilter{
			ProductIds:   &[]string{product.Id},
			PriceTypeIds: regularPriceTypeFilter.Ids,
			CurrencyIds:  &[]string{currency.Id},
			Active:       entity.BoolPtr(true),
			IsCount:      entity.BoolPtr(false),
		}

		regularPrice, _, err := uc.price.List(ctx, &regularPriceFilter, true)
		if err != nil {
			return nil, nil, err
		}

		// get item tags
		itemTags := make(map[uint32]entity.ProductListItemTag)

		// get tags
		tags, _, err := uc.tag.List(ctx, &(entity.TagFilter{
			ProductIds: &[]string{product.Id},
			TagTypeIds: tagTypeFilter.Ids,
			OrderBy:    entity.StringPtr("TagTypeId"),
			OrderDir:   entity.StringPtr("ASC"),
			Active:     entity.BoolPtr(true),
			IsCount:    entity.BoolPtr(false),
		}), true)
		if err != nil {
			return nil, nil, err
		}

		// get tag selects
		for _, tag := range *tags {
			itemTag := entity.ProductListItemTag{
				Name: (*tagTypes)[tag.TagTypeId].Name,
				Url:  (*tagTypes)[tag.TagTypeId].Url,
			}

			// get tag selects if tag has tag_select_id
			if tag.TagSelectId != "" {
				tagSelect, _, err := uc.tagSelect.List(ctx, &(entity.TagSelectFilter{
					Ids:        &[]string{tag.TagSelectId},
					TagTypeIds: tagTypeFilter.Ids,
					Active:     entity.BoolPtr(true),
					IsCount:    entity.BoolPtr(false),
				}), true)
				if err != nil {
					return nil, nil, err
				}

				selectName := (*tagSelect)[tag.TagSelectId].Name
				itemTag.Value = selectName
			} else {
				itemTag.Value = tag.Value
			}

			itemTags[tagOrder[tag.TagTypeId]] = itemTag
		}

		// make product list item
		productItem := entity.ProductListItem{
			Id:               product.Id,
			Sku:              product.Sku,
			Quantity:         1,
			Brand:            product.Brand,
			Name:             product.Name,
			ShortDescription: product.ShortDescription,
			Url:              product.Url,
			Currency:         currency.Symbol,
			SeoTitle:         product.SeoTitle,
			SeoDescription:   product.SeoDescription,
			Tags:             itemTags,
		}

		// price
		if regularPriceFilter.Ids != nil && len(*regularPriceFilter.Ids) > 0 {
			productItem.Price = humanize.FormatFloat("#,###.##", (*regularPrice)[(*regularPriceFilter.Ids)[0]].Price)
		}

		if specialPriceFilter.Ids != nil && len(*specialPriceFilter.Ids) > 0 {
			productItem.SalePrice = humanize.FormatFloat("#,###.##", (*specialPrice)[(*specialPriceFilter.Ids)[0]].Price)
		}

		// add product to dto
		productsDto[id] = productItem

		// count
		counter++
	}

	return &productsDto, totalItemsCount, nil
}
