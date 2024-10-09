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
) (*map[uuid.UUID]*entity.ProductListItem, error) {

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

		result, err := u.productInfo.List(ctx, &productInfoFilter)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}
		mu.Lock()
		productsInfo = result
		mu.Unlock()
		parameters.Count = productInfoFilter.Count
	}()

	// currency
	var currency *entity.Currency
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error

		// get default currency
		defaultCurrency := u.currency.GetDefault()
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
	productsDto := make(map[uuid.UUID]*entity.ProductListItem)
	counter := 0
	for id, product := range *productsInfo {
		// make product list item
		var productItem *entity.ProductListItem
		productItem = &entity.ProductListItem{
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			specialPriceTypeFilterIds := u.priceType.GetSpecialPriceTypeIds()

			specialPriceFilter := entity.PriceFilter{
				ProductIds:     &[]uuid.UUID{product.Id},
				PriceTypeIds:   &specialPriceTypeFilterIds,
				CurrencyIds:    &[]uuid.UUID{currency.Id},
				Active:         entity.BoolPtr(true),
				IsCount:        entity.BoolPtr(false),
				IsUpdateFilter: entity.BoolPtr(true),
			}

			specialPrice, err := u.price.List(ctx, &specialPriceFilter)
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}
			if specialPriceFilter.Ids != nil && len(*specialPriceFilter.Ids) > 0 {
				mu.Lock()
				productItem.SalePrice = humanize.FormatFloat("#,###.##", (*specialPrice)[(*specialPriceFilter.Ids)[0]].Price)
				mu.Unlock()
			}
		}()

		// get regularPrice
		wg.Add(1)
		go func() {
			defer wg.Done()

			regularPriceTypeIds := u.priceType.GetRegularPriceTypeIds()
			regularPriceFilter := entity.PriceFilter{
				ProductIds:     &[]uuid.UUID{product.Id},
				PriceTypeIds:   &regularPriceTypeIds,
				CurrencyIds:    &[]uuid.UUID{currency.Id},
				Active:         entity.BoolPtr(true),
				IsCount:        entity.BoolPtr(false),
				IsUpdateFilter: entity.BoolPtr(true),
			}

			regularPrice, err := u.price.List(ctx, &regularPriceFilter)
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}
			if regularPriceFilter.Ids != nil && len(*regularPriceFilter.Ids) > 0 {
				mu.Lock()
				productItem.Price = humanize.FormatFloat("#,###.##", (*regularPrice)[(*regularPriceFilter.Ids)[0]].Price)
				mu.Unlock()
			}
		}()

		//get item tags
		itemTags := make(map[uint32]entity.ProductListItemTag)
		wg.Add(1)
		go func() {
			defer wg.Done()
			defaultTagTypes := u.tagType.GetTagTypesForList(ctx)

			// get tags
			tags, err := u.tag.List(ctx, &(entity.TagFilter{
				ProductIds: &[]uuid.UUID{product.Id},
				TagTypeIds: defaultTagTypes.TagTypesIds,
				OrderBy:    entity.StringPtr("TagTypeId"),
				OrderDir:   entity.StringPtr("ASC"),
				Active:     entity.BoolPtr(true),
				IsCount:    entity.BoolPtr(false),
			}))
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

			// get tag selects
			for _, tag := range *tags {
				itemTag := entity.ProductListItemTag{
					Name: (*defaultTagTypes.TagTypes)[tag.TagTypeId].Name,
					Url:  (*defaultTagTypes.TagTypes)[tag.TagTypeId].Url,
				}

				// get tag selects if tag has tag_select_id
				if tag.TagSelectId != uuid.Nil {
					tagSelect, err := u.tagSelect.List(ctx, &(entity.TagSelectFilter{
						// omitted for brevity
					}))
					if err != nil {
						mu.Lock()
						errs = append(errs, err)
						mu.Unlock()
						return
					}

					selectName := (*tagSelect)[tag.TagSelectId].Name
					itemTag.Value = selectName
				} else {
					itemTag.Value = tag.Value
				}

				mu.Lock()
				tagOrder := *defaultTagTypes.TagOrder
				itemTags[tagOrder[tag.TagTypeId]] = itemTag
				mu.Unlock()
			}

			mu.Lock()
			productItem.Tags = itemTags
			mu.Unlock()
		}()

		// get stock quantity
		wg.Add(1)
		go func() {
			defer wg.Done()
			quantity, err := u.stockQuantity.Get(ctx, &entity.StockQuantityFilter{
				ProductIds: &[]uuid.UUID{product.Id},
				IsCount:    entity.BoolPtr(false),
			})
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

			mu.Lock()
			productItem.Quantity = quantity.Quality
			mu.Unlock()
		}()

		wg.Wait()

		// add product to dto
		productsDto[id] = productItem

		// count
		counter++

		// check errors
		if len(errs) > 0 {
			return nil, fmt.Errorf("GetProductList: %v", errs)
		}
	}

	return &productsDto, nil
}
