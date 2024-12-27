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
	shopId *uuid.UUID,
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

	// tags on the grid item
	var gridTagTypes pages.ProductGridTagTypes
	wg.Add(1)
	go func() {
		defer wg.Done()
		gridTagTypes = *u.fetchGridTags(ctx, shopId, &errs)
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
			item := u.fetchProductDetails(ctx, itemId, currency, &errs, &mu, gridTagTypes)
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

func (u *UseCase) fetchGridTags(
	ctx context.Context,
	shopId *uuid.UUID,
	errs *[]error,
) *pages.ProductGridTagTypes {
	// get cache
	if cache := u.getGridTagTypesCache(*shopId); cache != nil {
		return cache
	}

	// get tag types
	var gridTagTypes pages.ProductGridTagTypes
	tagOrder := make(map[uuid.UUID]uint64)
	gridTagTypes.TagOrder = &tagOrder

	// get tag types ids
	filter := db.ShopTagTypeFilter{
		ShopIds: &[]uuid.UUID{*shopId},
		Active:  pointer.BoolPtr(true),
		Sources: &[]string{"grid-tag"},
		DataConfig: &entity.DataConfig{
			IsCount:        pointer.BoolPtr(false),
			IsUpdateFilter: pointer.BoolPtr(true),
			IsKeepIdsOrder: pointer.BoolPtr(true),
		},
	}
	shopTagTypes, err := u.shopTagType.List(ctx, &filter)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	if shopTagTypes == nil {
		*errs = append(*errs, fmt.Errorf("shopTagTypes is nil"))
		return nil
	}

	// get tag types
	gridTagTypes.TagTypesIds = filter.TagTypeIds
	for _, item := range *shopTagTypes {
		(*gridTagTypes.TagOrder)[item.TagTypeId] = item.SortOrder
	}

	gridTagTypes.TagTypes, err = u.tagType.List(ctx, &(db.TagTypeFilter{
		Ids:    filter.TagTypeIds,
		Active: pointer.BoolPtr(true),
	}))
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	if gridTagTypes.TagTypes == nil {
		*errs = append(*errs, fmt.Errorf("gridTagTypes.TagTypes is nil"))
		return nil
	}

	// save to cache
	go u.setGridTagTypesCache(*shopId, &gridTagTypes)

	// return
	return &gridTagTypes
}

func (u *UseCase) fetchProductIds(
	ctx context.Context,
	parameters *pages.ProductsPageUrlParams,
	errs *[]error,
) *[]uuid.UUID {
	// hash parameters
	hash := crypt.HashFilter(parameters)

	// check cache
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

func (u *UseCase) fetchProductDetails(
	ctx context.Context,
	itemId uuid.UUID,
	currency *db.Currency,
	errs *[]error,
	mu *sync.Mutex,
	gridTagTypes pages.ProductGridTagTypes,
) *pages.ProductGridItem {
	// Check cache
	if item := u.getGridItemCache(itemId); item != nil {
		return item
	}

	// Create new item
	item := &pages.ProductGridItem{}
	var wg sync.WaitGroup

	// Fetch product details
	wg.Add(1)
	go func() {
		defer wg.Done()
		product, err := u.productInfo.Get(ctx, &db.ProductInfoFilter{
			Ids: &[]uuid.UUID{itemId},
		})
		if err != nil {
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

		tags, err := u.tag.List(ctx, &(db.TagFilter{
			ProductIds: &[]uuid.UUID{itemId},
			TagTypeIds: gridTagTypes.TagTypesIds,
			OrderBy:    pointer.StringToPtr("TagTypeId"),
			OrderDir:   pointer.StringToPtr("ASC"),
			Active:     pointer.BoolPtr(true),
			IsCount:    pointer.BoolPtr(false),
		}))

		if err != nil {

			*errs = append(*errs, err)
			return
		}

		for _, tag := range *tags {
			itemTag := pages.ProductListItemTag{
				Name: (*gridTagTypes.TagTypes)[tag.TagTypeId].Name,
				Url:  (*gridTagTypes.TagTypes)[tag.TagTypeId].Url,
			}

			if tag.TagSelectId != uuid.Nil {
				tagSelect, e := u.tagSelect.Get(ctx, &(db.TagSelectFilter{
					Ids: &[]uuid.UUID{tag.TagSelectId},
				}))
				if e != nil {
					*errs = append(*errs, e)
					return
				}

				itemTag.Value = tagSelect.Name
			} else {
				itemTag.Value = tag.Value
			}

			tagOrder := *gridTagTypes.TagOrder
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

			return
		}

		image, err := u.image.Get(ctx, &db.ImageFilter{
			Ids: &[]uuid.UUID{imageInfo.ImageId},
		})

		if err != nil {

			return
		}

		item.MainImage = image

		//if !image.IsCompressed {
		//	err := u.image.Compression(ctx, &db.ImageCompression{
		//		Ids:         &[]uuid.UUID{imageInfo.ImageId},
		//		Compression: pointer.UintPtr(80),
		//	})
		//	if err != nil {
		//		*errs = append(*errs, err)
		//	}
		//}
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

			return
		}

		image, _ := u.image.Get(ctx, &db.ImageFilter{
			Ids: &[]uuid.UUID{imageInfo.ImageId},
		})

		if image == nil {

			return
		}

		item.HoverImage = image

		//if !image.IsCompressed && image.Id != uuid.Nil {
		//	err := u.image.Compression(ctx, &db.ImageCompression{
		//		Ids:         &[]uuid.UUID{image.Id},
		//		Compression: pointer.UintPtr(80),
		//	})
		//	if err != nil {
		//		*errs = append(*errs, err)
		//	}
		//}
	}()

	wg.Wait()

	// error
	if len(*errs) > 0 {
		return nil
	}

	// Save to cache

	go u.setGridItemCache(itemId, item)

	// Return item
	return item
}
