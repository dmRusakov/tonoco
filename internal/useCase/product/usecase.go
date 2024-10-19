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
		/* make product list item */
		item := &entity.ProductListItem{
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

		/* get special price */
		wg.Add(1)
		go func() {
			defer wg.Done()

			// get special price type typeIds
			typeIds, err := u.priceType.GetDefaultIds("special")
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

			// make special price filter
			filter := entity.PriceFilter{
				ProductIds:     &[]uuid.UUID{product.Id},
				PriceTypeIds:   typeIds,
				CurrencyIds:    &[]uuid.UUID{currency.Id},
				Active:         entity.BoolPtr(true),
				IsCount:        entity.BoolPtr(false),
				IsUpdateFilter: entity.BoolPtr(true),
			}

			//	get special price
			price, err := u.price.List(ctx, &filter)
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

			// format special price
			if filter.Ids != nil && len(*filter.Ids) > 0 {
				mu.Lock()
				item.SalePrice = humanize.FormatFloat("#,###.##", (*price)[(*filter.Ids)[0]].Price)
				mu.Unlock()
			}
		}()

		/* get regularPrice */
		wg.Add(1)
		go func() {
			defer wg.Done()

			// get regular price type ids
			typeIds, err := u.priceType.GetDefaultIds("regular")
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

			// make regular price filter
			filter := entity.PriceFilter{
				ProductIds:     &[]uuid.UUID{product.Id},
				PriceTypeIds:   typeIds,
				CurrencyIds:    &[]uuid.UUID{currency.Id},
				Active:         entity.BoolPtr(true),
				IsCount:        entity.BoolPtr(false),
				IsUpdateFilter: entity.BoolPtr(true),
			}

			// get regular price
			price, err := u.price.List(ctx, &filter)
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

			// format regular price
			if filter.Ids != nil && len(*filter.Ids) > 0 {
				mu.Lock()
				item.Price = humanize.FormatFloat("#,###.##", (*price)[(*filter.Ids)[0]].Price)
				mu.Unlock()
			}
		}()

		/* get item tags */
		itemTags := make(map[uint32]entity.ProductListItemTag)
		wg.Add(1)
		go func() {
			defer wg.Done()

			// get default tag types
			defaultTagTypes, err := u.tagType.GetDefaultIds("list")
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

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
			item.Tags = itemTags
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
			item.Quantity = quantity.Quality
			if item.Quantity > 0 {
				item.Status = "in_stock"
			} else if item.Quantity < -50 {
				item.Status = "discontinued"
			} else {
				item.Status = "pre_order"
			}
			mu.Unlock()
		}()

		// product main image
		wg.Add(1)
		go func() {
			defer wg.Done()
			imageInfo, _ := u.productImage.Get(ctx, &entity.ProductImageFilter{
				ProductIds: &[]uuid.UUID{product.Id},
				IsCount:    entity.BoolPtr(false),
				Type:       &[]string{"main"},
			})

			if imageInfo == nil {
				return
			}

			image, _ := u.image.Get(ctx, &entity.ImageFilter{
				Ids: &[]uuid.UUID{imageInfo.ImageId},
			})

			if image != nil {
				mu.Lock()
				item.MainImage = image
				mu.Unlock()
			}

			// compress image if not compressed
			if !image.IsCompressed {
				err := u.image.Compression(ctx, &entity.ImageCompression{
					Ids:         &[]uuid.UUID{imageInfo.ImageId},
					Compression: entity.UintPtr(80),
				})
				if err != nil {
					mu.Lock()
					errs = append(errs, err)
					mu.Unlock()
				}
			}
		}()

		// hover image
		wg.Add(1)
		go func() {
			defer wg.Done()
			imageInfo, _ := u.productImage.Get(ctx, &entity.ProductImageFilter{
				ProductIds: &[]uuid.UUID{product.Id},
				IsCount:    entity.BoolPtr(false),
				Type:       &[]string{"hover"},
			})

			if imageInfo == nil {
				return
			}

			image, _ := u.image.Get(ctx, &entity.ImageFilter{
				Ids: &[]uuid.UUID{imageInfo.ImageId},
			})

			if image != nil {
				mu.Lock()
				item.HoverImage = image
				mu.Unlock()
			}
		}()

		wg.Wait()

		// add product to dto
		productsDto[id] = item

		// count
		counter++

		// check errors
		if len(errs) > 0 {
			return nil, fmt.Errorf("GetProductList: %v", errs)
		}
	}

	return &productsDto, nil
}
