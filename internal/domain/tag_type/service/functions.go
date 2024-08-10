package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	"time"
)

// clear cashes
func (s *Service) clearCashes() {
	s.itemCash = make(map[string]Item)
	s.itemsCash = make(map[string]map[string]Item)
	s.countCash = make(map[string]uint64)
}

// get item cash
func (s *Service) getItemCash(cacheKey string) *Item {
	if item, ok := s.itemCash[cacheKey]; ok {
		return &item
	}
	return nil
}

// set item cash
func (s *Service) setItemCash(cacheKey string, item *Item) {
	if item == nil {
		return
	}
	s.itemCash[cacheKey] = *item
}

// get items cash
func (s *Service) getItemsCash(cacheKey string) (*map[string]Item, *uint64, error) {
	items := s.itemsCash[cacheKey]
	count := s.countCash[cacheKey]
	if items != nil && count != 0 {
		return &items, &count, nil
	}

	return nil, nil, errors.AddCode(entity.ErrCacheNotFound, "bsfd33")
}

// set items cash
func (s *Service) setItemsCash(cacheKey string, items *map[string]Item, count *uint64) {
	if items == nil {
		items = new(map[string]Item)
	}
	s.itemsCash[cacheKey] = *items
	if count == nil {
		count = new(uint64)
	}
	s.countCash[cacheKey] = *count
}

func (s *Service) Get(ctx context.Context, filter *Filter) (*Item, error) {
	// Generate a itemCash key based on id and url
	cacheKey, err := entity.HashFilter(filter)
	if err != nil {
		return nil, err
	}

	// Try to get the item from the itemCash
	if item := s.getItemCash(cacheKey); item != nil {
		return item, nil
	}

	// If the item is not in the itemCash, get it from the repository
	item, err := s.repository.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Store the item in the itemCash
	s.setItemCash(cacheKey, item)

	return item, nil
}

func (s *Service) List(ctx context.Context, filter *Filter, isUpdateFilter bool) (*map[string]Item, *uint64, error) {
	// Generate a itemCash key based on the filter
	cacheKey, err := entity.HashFilter(filter)
	if err != nil {
		return nil, nil, err
	}

	// Try to get the items from the itemsCash
	items, count, err := s.getItemsCash(cacheKey)
	if err == nil {
		return items, count, nil
	}

	// If the items are not in the itemsCash, get them from the repository
	items, count, err = s.repository.List(ctx, filter, isUpdateFilter)
	if err != nil {
		return nil, nil, err
	}

	// Store the items in the itemsCash
	s.setItemsCash(cacheKey, items, count)

	// Return the items
	return items, count, nil
}

func (s *Service) Create(ctx context.Context, item *Item) (*string, error) {
	s.clearCashes()
	return s.repository.Create(ctx, item)
}

func (s *Service) Update(ctx context.Context, item *Item) error {
	s.clearCashes()
	return s.repository.Update(ctx, item)
}

func (s *Service) Patch(ctx context.Context, id *string, fields *map[string]interface{}) error {
	s.clearCashes()
	return s.repository.Patch(ctx, id, fields)
}

func (s *Service) UpdatedAt(ctx context.Context, id *string) (*time.Time, error) {
	return s.repository.UpdatedAt(ctx, id)
}

func (s *Service) TableIndexCount(ctx context.Context) (*uint64, error) {
	return s.repository.TableIndexCount(ctx)
}

func (s *Service) MaxSortOrder(ctx context.Context) (*uint64, error) {
	return s.repository.MaxSortOrder(ctx)
}

func (s *Service) Delete(ctx context.Context, id *string) error {
	s.clearCashes()
	return s.repository.Delete(ctx, id)
}
