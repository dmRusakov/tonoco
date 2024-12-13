package shop_page

import (
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/google/uuid"
)

func (u *UseCase) getGridItemCache(itemID uuid.UUID) *pages.ProductGridItem {
	return u.gridItemCache[itemID]
}

func (u *UseCase) setGridItemCache(key uuid.UUID, value *pages.ProductGridItem) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.gridItemCache[key] = value
}

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
