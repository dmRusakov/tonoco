package shop_page

import (
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/google/uuid"
)

// grid item cache
func (u *UseCase) getGridItemCache(itemID uuid.UUID) *pages.ProductGridItem {
	cache, ok := u.gridItemCache[itemID]
	if !ok {
		return nil
	}
	return cache
}

func (u *UseCase) setGridItemCache(key uuid.UUID, value *pages.ProductGridItem) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.gridItemCache[key] = value
}

// shop page cache
func (u *UseCase) getShopPageCache(key string) *pages.Shop {
	cache, ok := u.shopPageCache[key]
	if !ok {
		return nil
	}
	return cache
}

func (u *UseCase) setShopPageCache(key string, value *pages.Shop) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.shopPageCache[key] = value
}

// shop page filter cache
func (u *UseCase) getShopPageFilterCache(id uuid.UUID) *map[uint64]pages.ShopPageFilterItem {
	cache, ok := u.shopPageFilterCache[id]
	if !ok {
		return nil
	}
	return cache
}

func (u *UseCase) setShopPageFilterCache(id uuid.UUID, value *map[uint64]pages.ShopPageFilterItem) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.shopPageFilterCache[id] = value
}

// grid tag types cache
func (u *UseCase) getGridTagTypesCache(id uuid.UUID) *pages.ProductGridTagTypes {
	cache, ok := u.gridTagTypesCache[id]
	if !ok {
		return nil
	}
	return cache
}

func (u *UseCase) setGridTagTypesCache(id uuid.UUID, value *pages.ProductGridTagTypes) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.gridTagTypesCache[id] = value
}

// item ids cache
func (u *UseCase) getItemIdsCache(id uuid.UUID) (*[]uuid.UUID, *uint64) {
	cache, ok := u.itemIdsCache[id]
	if !ok {
		return nil, nil
	}

	return cache.ids, cache.count
}

func (u *UseCase) setItemIdsCache(id uuid.UUID, item *[]uuid.UUID, count *uint64) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.itemIdsCache[id] = struct {
		ids   *[]uuid.UUID
		count *uint64
	}{
		ids:   item,
		count: count,
	}

}
