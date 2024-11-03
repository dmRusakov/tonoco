package product

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"sync"
)

func (u *UseCase) GetProductList(
	ctx context.Context,
	parameters *pages.ProductsPageUrlParams,
) (*map[uuid.UUID]*pages.ProductListItem, error) {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	// get product ids
	var productIds *[]uuid.UUID
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		productIdsFilter := db.ProductInfoFilter{
			Page:    parameters.Page,
			PerPage: parameters.PerPage,
			IsCount: pointer.BoolPtr(true),
		}
		productIds, err = u.productInfo.Ids(ctx, &productIdsFilter)
		if err != nil {
			errs = append(errs, err)
			return
		}

		parameters.Count = productIdsFilter.Count
	}()

	// currency
	var currency *db.Currency
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
			currency, err = u.currency.Get(ctx, &db.CurrencyFilter{
				Urls: &[]string{*parameters.Currency},
			})
			if err != nil {
				errs = append(errs, err)
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
	productsDto := make(map[uuid.UUID]*pages.ProductListItem)
	counter := 0
	for _, productId := range *productIds {
		item := &pages.ProductListItem{}

		/* product info */
		wg.Add(1)
		go func() {
			defer wg.Done()
			product, err := u.productInfo.Get(ctx, &db.ProductInfoFilter{
				Ids: &[]uuid.UUID{productId},
			})
			if err != nil {
				errs = append(errs, err)
				return
			}

			item.Id = product.Id
			item.Sku = product.Sku
			item.Brand = product.Brand
			item.Name = product.Name
			item.ShortDescription = product.ShortDescription
			item.Url = product.Url
			item.SeoTitle = product.SeoTitle
			item.SeoDescription = product.SeoDescription
		}()

		// /* make product list item */
		// item := &pages.ProductListItem{
		// 	Id:               product.Id,
		// 	Sku:              product.Sku,
		// 	Brand:            product.Brand,
		// 	Name:             product.Name,
		// 	ShortDescription: product.ShortDescription,
		// 	Url:              product.Url,
		// 	Currency:         currency.Symbol,
		// 	SeoTitle:         product.SeoTitle,
		// 	SeoDescription:   product.SeoDescription,
		// }
		//
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
			filter := db.PriceFilter{
				ProductIds:     &[]uuid.UUID{productId},
				PriceTypeIds:   typeIds,
				CurrencyIds:    &[]uuid.UUID{currency.Id},
				Active:         pointer.BoolPtr(true),
				IsCount:        pointer.BoolPtr(false),
				IsUpdateFilter: pointer.BoolPtr(true),
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
			filter := db.PriceFilter{
				ProductIds:     &[]uuid.UUID{productId},
				PriceTypeIds:   typeIds,
				CurrencyIds:    &[]uuid.UUID{currency.Id},
				Active:         pointer.BoolPtr(true),
				IsCount:        pointer.BoolPtr(false),
				IsUpdateFilter: pointer.BoolPtr(true),
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
		itemTags := make(map[uint64]pages.ProductListItemTag)
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
			tags, err := u.tag.List(ctx, &(db.TagFilter{
				ProductIds: &[]uuid.UUID{productId},
				TagTypeIds: defaultTagTypes.TagTypesIds,
				OrderBy:    pointer.StringToPtr("TagTypeId"),
				OrderDir:   pointer.StringToPtr("ASC"),
				Active:     pointer.BoolPtr(true),
				IsCount:    pointer.BoolPtr(false),
			}))

			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}

			// get tag selects
			for _, tag := range *tags {
				itemTag := pages.ProductListItemTag{
					Name: (*defaultTagTypes.TagTypes)[tag.TagTypeId].Name,
					Url:  (*defaultTagTypes.TagTypes)[tag.TagTypeId].Url,
				}

				// get tag selects if tag has tag_select_id
				if tag.TagSelectId != uuid.Nil {
					tagSelect, err := u.tagSelect.List(ctx, &(db.TagSelectFilter{
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
			quantity, err := u.stockQuantity.Get(ctx, &db.StockQuantityFilter{
				ProductIds: &[]uuid.UUID{productId},
				IsCount:    pointer.BoolPtr(false),
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
			imageInfo, _ := u.productImage.Get(ctx, &db.ProductImageFilter{
				ProductIds: &[]uuid.UUID{productId},
				IsCount:    pointer.BoolPtr(false),
				Type:       &[]string{"main"},
			})

			if imageInfo == nil {
				return
			}

			image, err := u.image.Get(ctx, &db.ImageFilter{
				Ids: &[]uuid.UUID{imageInfo.ImageId},
			})

			if err != nil {
				return
			}

			mu.Lock()
			item.MainImage = image
			mu.Unlock()

			// compress image if not compressed
			if !image.IsCompressed {
				err := u.image.Compression(ctx, &db.ImageCompression{
					Ids:         &[]uuid.UUID{imageInfo.ImageId},
					Compression: pointer.UintPtr(80),
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
			imageInfo, _ := u.productImage.Get(ctx, &db.ProductImageFilter{
				ProductIds: &[]uuid.UUID{productId},
				IsCount:    pointer.BoolPtr(false),
				Type:       &[]string{"hover"},
			})

			if imageInfo == nil {
				return
			}

			image, _ := u.image.Get(ctx, &db.ImageFilter{
				Ids: &[]uuid.UUID{imageInfo.ImageId},
			})

			if image == nil {
				return
			}

			mu.Lock()
			item.HoverImage = image
			mu.Unlock()

			// compress image if not compressed
			if !image.IsCompressed && image.Id != uuid.Nil {
				err := u.image.Compression(ctx, &db.ImageCompression{
					Ids:         &[]uuid.UUID{image.Id},
					Compression: pointer.UintPtr(80),
				})
				if err != nil {
					mu.Lock()
					errs = append(errs, err)
					mu.Unlock()
				}
			}
		}()

		wg.Wait()

		// add product to dto
		productsDto[productId] = item

		// count
		counter++

		// check errors
		if len(errs) > 0 {
			return nil, fmt.Errorf("GetProductList: %v", errs)
		}
	}

	return &productsDto, nil
}
