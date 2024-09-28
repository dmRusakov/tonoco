package product

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"sync"
)

func (uc *UseCase) GetProductList(
	ctx context.Context,
	parameters *entity.ProductsPageParams,
	appData *entity.AppData,
) (*map[uuid.UUID]entity.ProductListItem, error) {

	var wg sync.WaitGroup
	var errs []error

	// get productsInfo
	var productsInfo *map[uuid.UUID]entity.ProductInfo
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		productInfoFilter := entity.ProductInfoFilter{
			Page:           parameters.Page,
			PerPage:        parameters.PerPage,
			IsCount:        entity.BoolPtr(true),
			IsUpdateFilter: entity.BoolPtr(true),
		}

		productsInfo, err = uc.productInfo.List(ctx, &productInfoFilter)
		if err != nil {
			errs = append(errs, err)
			return
		}
		parameters.Count = productInfoFilter.Count
	}()

	// get currencies
	var currency *entity.Currency
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		currency, err = uc.currency.Get(ctx, &entity.CurrencyFilter{
			IsCount: entity.BoolPtr(true),
			Urls:    &[]string{uc.store.DefaultStore.CurrencyUrl},
		})
		if err != nil {
			errs = append(errs, err)
			return
		}
	}()

	// get special prices type IDs
	var specialPriceTypeFilter *entity.PriceTypeFilter
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		specialPriceTypeFilter = &entity.PriceTypeFilter{
			Urls:           &[]string{"special", "sale"},
			IsPublic:       entity.BoolPtr(true),
			IsIdsOnly:      entity.BoolPtr(true),
			IsCount:        entity.BoolPtr(false),
			IsUpdateFilter: entity.BoolPtr(true),
		}

		_, err = uc.priceType.List(ctx, specialPriceTypeFilter)
		if err != nil {
			errs = append(errs, err)
			return
		}
	}()

	// get regular prices type IDs
	var regularPriceTypeFilter *entity.PriceTypeFilter
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		regularPriceTypeFilter = &entity.PriceTypeFilter{
			Urls:           &[]string{"regular"},
			IsPublic:       entity.BoolPtr(true),
			IsIdsOnly:      entity.BoolPtr(true),
			IsCount:        entity.BoolPtr(false),
			IsUpdateFilter: entity.BoolPtr(true),
		}

		_, err = uc.priceType.List(ctx, regularPriceTypeFilter)
		if err != nil {
			errs = append(errs, err)
			return
		}
	}()

	// get tag_types with `list_item` type
	var tagTypes *map[uuid.UUID]entity.TagType
	var tagTypeFilter *entity.TagTypeFilter
	tagOrder := make(map[uuid.UUID]uint32)
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		tagTypeFilter = &entity.TagTypeFilter{
			OrderBy:        entity.StringPtr("SortOrder"),
			OrderDir:       entity.StringPtr("ASC"),
			ListItem:       entity.BoolPtr(true),
			Active:         entity.BoolPtr(true),
			IsCount:        entity.BoolPtr(false),
			IsUpdateFilter: entity.BoolPtr(true),
		}

		tagTypes, err = uc.tagType.List(ctx, tagTypeFilter)
		if err != nil {
			errs = append(errs, err)
			return
		}

		// tag order
		for i, tagType := range *tagTypeFilter.Ids {
			tagOrder[tagType] = uint32(i)
		}
	}()
	wg.Wait()

	// check errors
	if len(errs) > 0 {
		// TODO log errors
		return nil, fmt.Errorf("GetProductList: %v", errs)
	}

	// dto
	productsDto := make(map[uuid.UUID]entity.ProductListItem)
	counter := 0
	for id, product := range *productsInfo {
		var group sync.WaitGroup

		// make product list item
		productItem := entity.ProductListItem{
			Id:               product.Id,
			Sku:              product.Sku,
			Brand:            product.Brand,
			Name:             product.Name,
			ShortDescription: product.ShortDescription,
			Url:              product.Url,
			Currency:         currency.Symbol,
			SeoTitle:         product.SeoTitle,
			SeoDescription:   product.SeoDescription,
		}

		// get special price
		group.Add(1)
		go func() {
			defer group.Done()
			var err error
			specialPriceFilter := &entity.PriceFilter{
				ProductIds:     &[]uuid.UUID{product.Id},
				PriceTypeIds:   specialPriceTypeFilter.Ids,
				CurrencyIds:    &[]uuid.UUID{currency.Id},
				Active:         entity.BoolPtr(true),
				IsCount:        entity.BoolPtr(false),
				IsUpdateFilter: entity.BoolPtr(true),
			}
			specialPrice, err := uc.price.List(ctx, specialPriceFilter)
			if err != nil {
				errs = append(errs, err)
				return
			}
			if specialPriceFilter.Ids != nil && len(*specialPriceFilter.Ids) > 0 {
				productItem.SalePrice = humanize.FormatFloat("#,###.##", (*specialPrice)[(*specialPriceFilter.Ids)[0]].Price)
			}
		}()

		// get regularPrice
		group.Add(1)
		go func() {
			defer group.Done()
			var err error
			regularPriceFilter := &entity.PriceFilter{
				ProductIds:     &[]uuid.UUID{product.Id},
				PriceTypeIds:   regularPriceTypeFilter.Ids,
				CurrencyIds:    &[]uuid.UUID{currency.Id},
				Active:         entity.BoolPtr(true),
				IsCount:        entity.BoolPtr(false),
				IsUpdateFilter: entity.BoolPtr(true),
			}
			regularPrice, err := uc.price.List(ctx, regularPriceFilter)
			if err != nil {
				errs = append(errs, err)
				return
			}
			if regularPriceFilter.Ids != nil && len(*regularPriceFilter.Ids) > 0 {
				productItem.Price = humanize.FormatFloat("#,###.##", (*regularPrice)[(*regularPriceFilter.Ids)[0]].Price)
			}
		}()

		// get item tags
		itemTags := make(map[uint32]entity.ProductListItemTag)
		group.Add(1)
		go func() {
			defer group.Done()
			// get tags
			tags, err := uc.tag.List(ctx, &(entity.TagFilter{
				ProductIds: &[]uuid.UUID{product.Id},
				TagTypeIds: tagTypeFilter.Ids,
				OrderBy:    entity.StringPtr("TagTypeId"),
				OrderDir:   entity.StringPtr("ASC"),
				Active:     entity.BoolPtr(true),
				IsCount:    entity.BoolPtr(false),
			}))
			if err != nil {
				errs = append(errs, err)
				return
			}

			// get tag selects
			for _, tag := range *tags {
				itemTag := entity.ProductListItemTag{
					Name: (*tagTypes)[tag.TagTypeId].Name,
					Url:  (*tagTypes)[tag.TagTypeId].Url,
				}

				// get tag selects if tag has tag_select_id
				if tag.TagSelectId != uuid.Nil {
					tagSelect, err := uc.tagSelect.List(ctx, &(entity.TagSelectFilter{
						Ids:        &[]uuid.UUID{tag.TagSelectId},
						TagTypeIds: tagTypeFilter.Ids,
						Active:     entity.BoolPtr(true),
						IsCount:    entity.BoolPtr(false),
					}))
					if err != nil {
						errs = append(errs, err)
					}

					selectName := (*tagSelect)[tag.TagSelectId].Name
					itemTag.Value = selectName
				} else {
					itemTag.Value = tag.Value
				}

				itemTags[tagOrder[tag.TagTypeId]] = itemTag
			}
		}()

		// get stock quantity
		group.Add(1)
		go func() {
			defer group.Done()
			quantity, err := uc.stockQuantity.Get(ctx, &entity.StockQuantityFilter{
				ProductIds: &[]uuid.UUID{product.Id},
				IsCount:    entity.BoolPtr(false),
			})
			if err != nil {
				errs = append(errs, err)
				return
			}

			productItem.Quantity = quantity.Quality
		}()

		group.Wait()

		// got product images IDs (main, hover)
		productImagesFilter := entity.ProductImageFilter{
			ProductIds:     &[]uuid.UUID{product.Id},
			Type:           &[]string{"main", "hover"},
			IsCount:        entity.BoolPtr(false),
			IsUpdateFilter: entity.BoolPtr(true),
		}
		productImages, err := uc.productImageService.List(ctx, &productImagesFilter)
		if err != nil {
			return nil, err
		}

		// get product images info
		imageFilter := entity.ImageFilter{
			Ids: productImagesFilter.ImageIds,
		}
		images, err := uc.imageService.List(ctx, &imageFilter)

		//

		appData.ConsoleMessage.Log = append(appData.ConsoleMessage.Log, fmt.Sprintf("products:177 %v", productImages))
		appData.ConsoleMessage.Log = append(appData.ConsoleMessage.Log, fmt.Sprintf("products:178 %v", productImagesFilter.ImageIds))
		appData.ConsoleMessage.Log = append(appData.ConsoleMessage.Log, fmt.Sprintf("products:213 %v", images))

		// add product to dto
		productsDto[id] = productItem

		// count
		counter++
	}

	return &productsDto, nil
}
