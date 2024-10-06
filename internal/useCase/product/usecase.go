package product

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"sync"
)

func (u *UseCase) GetProductList(
	ctx context.Context,
	parameters *entity.ProductsPageUrlParams,
	appData *entity.AppData,
) (*map[uuid.UUID]entity.ProductListItem, error) {

	var wg sync.WaitGroup
	var mu sync.Mutex
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

		productsInfo, err = u.productInfo.List(ctx, &productInfoFilter)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}
		parameters.Count = productInfoFilter.Count
	}()

	// get special prices type IDs
	var specialPriceTypeFilterIds []uuid.UUID
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		specialPriceTypeFilter := &entity.PriceTypeFilter{
			Urls:           &[]string{"special", "sale"},
			IsPublic:       entity.BoolPtr(true),
			IsIdsOnly:      entity.BoolPtr(true),
			IsCount:        entity.BoolPtr(false),
			IsUpdateFilter: entity.BoolPtr(true),
		}

		_, err = u.priceType.List(ctx, specialPriceTypeFilter)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}

		specialPriceTypeFilterIds = *specialPriceTypeFilter.Ids
	}()

	// get regular prices type IDs
	var regularPriceTypeFilterIds []uuid.UUID
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		regularPriceTypeFilter := &entity.PriceTypeFilter{
			Urls:           &[]string{"regular"},
			IsPublic:       entity.BoolPtr(true),
			IsIdsOnly:      entity.BoolPtr(true),
			IsCount:        entity.BoolPtr(false),
			IsUpdateFilter: entity.BoolPtr(true),
		}

		_, err = u.priceType.List(ctx, regularPriceTypeFilter)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}
		regularPriceTypeFilterIds = *regularPriceTypeFilter.Ids
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

		tagTypes, err = u.tagType.List(ctx, tagTypeFilter)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}

		// tag order
		for i, tagType := range *tagTypeFilter.Ids {
			tagOrder[tagType] = uint32(i)
		}
	}()

	// currency
	var currency *entity.Currency
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error

		// get default currency
		defaultCurrency := u.currency.DefaultCurrency
		if (parameters.Currency == nil) || (*parameters.Currency == "" || *parameters.Currency == defaultCurrency.Url) {
			currency = defaultCurrency
			return
		} else {
			// get currency by url
			currency, err = u.currency.Get(ctx, &entity.CurrencyFilter{
				Urls: &[]string{*parameters.Currency},
			})
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}
		}
	}()

	wg.Wait()

	// check errors
	if len(errs) > 0 {
		return nil, fmt.Errorf("GetProductList: %v", errs)
	}

	// dto
	productsDto := make(map[uuid.UUID]entity.ProductListItem)
	counter := 0
	for id, product := range *productsInfo {
		var subWg sync.WaitGroup
		var subMu sync.Mutex
		var subErrs []error

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
		subWg.Add(1)
		go func(
			productId uuid.UUID,
			currencyId uuid.UUID,
			specialPriceTypeFilterIds []uuid.UUID,
		) {
			defer subWg.Done()
			specialPriceFilter := &entity.PriceFilter{
				ProductIds:     &[]uuid.UUID{productId},
				PriceTypeIds:   &specialPriceTypeFilterIds,
				CurrencyIds:    &[]uuid.UUID{currencyId},
				Active:         entity.BoolPtr(true),
				IsCount:        entity.BoolPtr(false),
				IsUpdateFilter: entity.BoolPtr(true),
			}
			specialPrice, err := u.price.List(ctx, specialPriceFilter)
			if err != nil {
				subMu.Lock()
				subErrs = append(subErrs, err)
				subMu.Unlock()
				return
			}
			if specialPriceFilter.Ids != nil && len(*specialPriceFilter.Ids) > 0 {
				subMu.Lock()
				productItem.SalePrice = humanize.FormatFloat("#,###.##", (*specialPrice)[(*specialPriceFilter.Ids)[0]].Price)
				subMu.Unlock()
			}
		}(product.Id, currency.Id, specialPriceTypeFilterIds)

		// get regularPrice
		subWg.Add(1)
		go func(
			productId uuid.UUID,
			currencyId uuid.UUID,
			regularPriceTypeFilterIds []uuid.UUID,
		) {
			defer subWg.Done()

			regularPriceFilter := &entity.PriceFilter{
				ProductIds:     &[]uuid.UUID{productId},
				PriceTypeIds:   &regularPriceTypeFilterIds,
				CurrencyIds:    &[]uuid.UUID{currencyId},
				Active:         entity.BoolPtr(true),
				IsCount:        entity.BoolPtr(false),
				IsUpdateFilter: entity.BoolPtr(true),
			}
			regularPrice, err := u.price.List(ctx, regularPriceFilter)
			if err != nil {
				subMu.Lock()
				subErrs = append(subErrs, err)
				subMu.Unlock()
				return
			}
			if regularPriceFilter.Ids != nil && len(*regularPriceFilter.Ids) > 0 {
				subMu.Lock()
				productItem.Price = humanize.FormatFloat("#,###.##", (*regularPrice)[(*regularPriceFilter.Ids)[0]].Price)
				subMu.Unlock()
			}
		}(product.Id, currency.Id, regularPriceTypeFilterIds)

		//get item tags
		itemTags := make(map[uint32]entity.ProductListItemTag)
		subWg.Add(1)
		go func(
			productId uuid.UUID,
			tagTypeFilterIds []uuid.UUID,
			tagTypes map[uuid.UUID]entity.TagType,
			tagOrder map[uuid.UUID]uint32,
		) {
			defer subWg.Done()
			// get tags
			tags, err := u.tag.List(ctx, &(entity.TagFilter{
				ProductIds: &[]uuid.UUID{productId},
				TagTypeIds: &tagTypeFilterIds,
				OrderBy:    entity.StringPtr("TagTypeId"),
				OrderDir:   entity.StringPtr("ASC"),
				Active:     entity.BoolPtr(true),
				IsCount:    entity.BoolPtr(false),
			}))
			if err != nil {
				subMu.Lock()
				subErrs = append(subErrs, err)
				subMu.Unlock()
				return
			}

			// get tag selects
			for _, tag := range *tags {
				itemTag := entity.ProductListItemTag{
					Name: (tagTypes)[tag.TagTypeId].Name,
					Url:  (tagTypes)[tag.TagTypeId].Url,
				}

				// get tag selects if tag has tag_select_id
				if tag.TagSelectId != uuid.Nil {
					tagSelect, err := u.tagSelect.List(ctx, &(entity.TagSelectFilter{
						// omitted for brevity
					}))
					if err != nil {
						subMu.Lock()
						subErrs = append(subErrs, err)
						subMu.Unlock()
						return
					}

					selectName := (*tagSelect)[tag.TagSelectId].Name
					itemTag.Value = selectName
				} else {
					itemTag.Value = tag.Value
				}

				subMu.Lock()
				itemTags[tagOrder[tag.TagTypeId]] = itemTag
				subMu.Unlock()
			}
		}(product.Id, *tagTypeFilter.Ids, *tagTypes, tagOrder)

		// get stock quantity
		subWg.Add(1)
		go func(
			productId uuid.UUID,
		) {
			defer subWg.Done()
			quantity, err := u.stockQuantity.Get(ctx, &entity.StockQuantityFilter{
				ProductIds: &[]uuid.UUID{productId},
				IsCount:    entity.BoolPtr(false),
			})
			if err != nil {
				subMu.Lock()
				subErrs = append(subErrs, err)
				subMu.Unlock()
				return
			}

			subMu.Lock()
			productItem.Quantity = quantity.Quality
			subMu.Unlock()
		}(product.Id)

		subWg.Wait()

		// add product to dto
		productsDto[id] = productItem

		// count
		counter++

		// check errors
		if len(subErrs) > 0 {
			return nil, fmt.Errorf("GetProductList: %v", subErrs)
		}
	}

	return &productsDto, nil
}
