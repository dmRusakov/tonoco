package product

import (
	"context"
	"fmt"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
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
