package shop_page

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/dmRusakov/tonoco/pkg/utils/crypt"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"sync"
)

func (u *UseCase) GetProductList(
	ctx context.Context,
	parameters *pages.ProductsPageUrlParams,
) (*map[uuid.UUID]*pages.ProductGridItem, *[]uuid.UUID, error) {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errs []error

	// get product ids
	var itemIds *[]uuid.UUID
	wg.Add(1)
	go func() {
		defer wg.Done()
		itemIds = u.fetchProductIds(ctx, parameters, &errs)
	}()

	// currency
	var currency *db.Currency
	wg.Add(1)
	go func() {
		defer wg.Done()
		currency = u.getCurrency(ctx, parameters, &errs)
	}()

	wg.Wait()

	// check errors
	if len(errs) > 0 {
		return nil, nil, fmt.Errorf("GetProductList: %v", errs)
	}

	// dto
	productsDto := make(map[uuid.UUID]*pages.ProductGridItem)

	for _, itemId := range *itemIds {
		wg.Add(1)
		go func(itemId uuid.UUID) {
			defer wg.Done()
			item := u.fetchProductDetails(ctx, itemId, currency, &errs, &mu)
			mu.Lock()
			productsDto[itemId] = item
			mu.Unlock()

			// check errors
			if len(errs) > 0 {
				return
			}
		}(itemId)
	}

	wg.Wait()

	// check errors
	if len(errs) > 0 {
		return nil, nil, fmt.Errorf("GetProductList: %v", errs)
	}

	return &productsDto, itemIds, nil
}

func (u *UseCase) fetchProductIds(ctx context.Context, parameters *pages.ProductsPageUrlParams, errs *[]error) *[]uuid.UUID {
	// hash parameters
	hash := crypt.HashFilter(parameters)
	if itemIds, count := u.getItemIdsCache(hash); itemIds != nil {
		parameters.Count = count
		return itemIds
	}

	// get product ids
	var productIds *[]uuid.UUID
	productIdsFilter := db.ProductInfoFilter{
		DataPagination: &entity.DataPagination{
			Page:    parameters.Page,
			PerPage: parameters.PerPage,
		},
		DataConfig: &entity.DataConfig{
			IsCount: pointer.BoolPtr(true),
		},
	}
	productIds, err := u.productInfo.Ids(ctx, &productIdsFilter)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	parameters.Count = productIdsFilter.DataConfig.Count

	// save to cache
	go u.setItemIdsCache(hash, productIds, parameters.Count)

	return productIds
}

func (u *UseCase) fetchProductDetails(ctx context.Context, itemId uuid.UUID, currency *db.Currency, errs *[]error, mu *sync.Mutex) *pages.ProductGridItem {
	// Check cache
	if item := u.getGridItemCache(itemId); item != nil {
		return item
	}

	// Create new item
	item := &pages.ProductGridItem{}
	var wg sync.WaitGroup
	isOk := true

	// Fetch product details
	wg.Add(1)
	go func() {
		defer wg.Done()
		product, err := u.productInfo.Get(ctx, &db.ProductInfoFilter{
			Ids: &[]uuid.UUID{itemId},
		})
		if err != nil {
			isOk = false
			*errs = append(*errs, err)
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

	// Fetch product special prices
	wg.Add(1)
	go func() {
		defer wg.Done()
		typeIds, err := u.priceType.GetDefaultIds("special")
		if err != nil {
			isOk = false
			*errs = append(*errs, err)
			return
		}

		filter := db.PriceFilter{
			ProductIds:     &[]uuid.UUID{itemId},
			PriceTypeIds:   typeIds,
			CurrencyIds:    &[]uuid.UUID{currency.Id},
			Active:         pointer.BoolPtr(true),
			IsCount:        pointer.BoolPtr(false),
			IsUpdateFilter: pointer.BoolPtr(true),
		}

		price, err := u.price.List(ctx, &filter)
		if err != nil {
			isOk = false
			*errs = append(*errs, err)
			return
		}

		if filter.Ids != nil && len(*filter.Ids) > 0 {
			item.SalePrice = humanize.FormatFloat("#,###.##", (*price)[(*filter.Ids)[0]].Price)
		}
	}()

	// Fetch product regular prices
	wg.Add(1)
	go func() {
		defer wg.Done()
		typeIds, err := u.priceType.GetDefaultIds("regular")
		if err != nil {
			isOk = false
			*errs = append(*errs, err)
			return
		}

		filter := db.PriceFilter{
			ProductIds:     &[]uuid.UUID{itemId},
			PriceTypeIds:   typeIds,
			CurrencyIds:    &[]uuid.UUID{currency.Id},
			Active:         pointer.BoolPtr(true),
			IsCount:        pointer.BoolPtr(false),
			IsUpdateFilter: pointer.BoolPtr(true),
		}

		price, err := u.price.List(ctx, &filter)
		if err != nil {
			isOk = false
			*errs = append(*errs, err)
			return
		}

		if filter.Ids != nil && len(*filter.Ids) > 0 {
			item.Price = humanize.FormatFloat("#,###.##", (*price)[(*filter.Ids)[0]].Price)
		}
	}()

	// Fetch product tags
	itemTags := make(map[uint64]pages.ProductListItemTag)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defaultTagTypes, err := u.tagType.GetDefaultIds("list")
		if err != nil {
			isOk = false
			*errs = append(*errs, err)
			return
		}

		tags, err := u.tag.List(ctx, &(db.TagFilter{
			ProductIds: &[]uuid.UUID{itemId},
			TagTypeIds: defaultTagTypes.TagTypesIds,
			OrderBy:    pointer.StringToPtr("TagTypeId"),
			OrderDir:   pointer.StringToPtr("ASC"),
			Active:     pointer.BoolPtr(true),
			IsCount:    pointer.BoolPtr(false),
		}))

		if err != nil {
			isOk = false
			*errs = append(*errs, err)
			return
		}

		for _, tag := range *tags {
			itemTag := pages.ProductListItemTag{
				Name: (*defaultTagTypes.TagTypes)[tag.TagTypeId].Name,
				Url:  (*defaultTagTypes.TagTypes)[tag.TagTypeId].Url,
			}

			if tag.TagSelectId != uuid.Nil {
				tagSelect, err := u.tagSelect.List(ctx, &(db.TagSelectFilter{
					// omitted for brevity
				}))
				if err != nil {
					isOk = false
					*errs = append(*errs, err)
					return
				}

				selectName := (*tagSelect)[tag.TagSelectId].Name
				itemTag.Value = selectName
			} else {
				itemTag.Value = tag.Value
			}

			tagOrder := *defaultTagTypes.TagOrder
			itemTags[tagOrder[tag.TagTypeId]] = itemTag
		}

		mu.Lock()
		item.Tags = itemTags
		mu.Unlock()
	}()

	// Fetch product stock quantity
	wg.Add(1)
	go func() {
		defer wg.Done()
		quantity, err := u.stockQuantity.Get(ctx, &db.StockQuantityFilter{
			ProductIds: &[]uuid.UUID{itemId},
			IsCount:    pointer.BoolPtr(false),
		})
		if err != nil {
			isOk = false
			*errs = append(*errs, err)
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

	// Fetch product main image
	wg.Add(1)
	go func() {
		defer wg.Done()
		imageInfo, _ := u.productImage.Get(ctx, &db.ProductImageFilter{
			ProductIds: &[]uuid.UUID{itemId},
			IsCount:    pointer.BoolPtr(false),
			Type:       &[]string{"main"},
		})

		if imageInfo == nil {
			isOk = false
			return
		}

		image, err := u.image.Get(ctx, &db.ImageFilter{
			Ids: &[]uuid.UUID{imageInfo.ImageId},
		})

		if err != nil {
			isOk = false
			return
		}

		item.MainImage = image

		if !image.IsCompressed {
			err := u.image.Compression(ctx, &db.ImageCompression{
				Ids:         &[]uuid.UUID{imageInfo.ImageId},
				Compression: pointer.UintPtr(80),
			})
			if err != nil {
				isOk = false
				*errs = append(*errs, err)
			}
		}
	}()

	// Fetch product hover image
	wg.Add(1)
	go func() {
		defer wg.Done()
		imageInfo, _ := u.productImage.Get(ctx, &db.ProductImageFilter{
			ProductIds: &[]uuid.UUID{itemId},
			IsCount:    pointer.BoolPtr(false),
			Type:       &[]string{"hover"},
		})

		if imageInfo == nil {
			isOk = false
			return
		}

		image, _ := u.image.Get(ctx, &db.ImageFilter{
			Ids: &[]uuid.UUID{imageInfo.ImageId},
		})

		if image == nil {
			isOk = false
			return
		}

		item.HoverImage = image

		if !image.IsCompressed && image.Id != uuid.Nil {
			err := u.image.Compression(ctx, &db.ImageCompression{
				Ids:         &[]uuid.UUID{image.Id},
				Compression: pointer.UintPtr(80),
			})
			if err != nil {
				isOk = false
				*errs = append(*errs, err)
			}
		}
	}()

	wg.Wait()

	// Save to cache
	if isOk {
		go u.setGridItemCache(itemId, item)
	}

	// Return item
	return item
}
