package product

import (
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/google/uuid"
)

func (u *UseCase) getGridItemCache(itemID uuid.UUID) *pages.ProductGridItem {
	return u.gridItemCache[itemID]
}

func (u *UseCase) setGridItemCache(itemID uuid.UUID, item *pages.ProductGridItem) {
	u.gridItemCache[itemID] = item
}

func (u *UseCase) getItemIdsCache(id uuid.UUID) (*[]uuid.UUID, *uint64) {
	cache, ok := u.itemIdsCache[id]
	if !ok {
		return nil, nil
	}

	return cache.ids, cache.count
}

func (u *UseCase) setItemIdsCache(id uuid.UUID, item *[]uuid.UUID, count *uint64) {
	u.itemIdsCache[id] = struct {
		ids   *[]uuid.UUID
		count *uint64
	}{
		ids:   item,
		count: count,
	}
}
