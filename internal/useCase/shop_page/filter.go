package shop_page

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
	"github.com/google/uuid"
)

func (u *UseCase) GetShopPageFilter(
	ctx context.Context,
	shopId *uuid.UUID,
) (*map[uint64]pages.ShopPageFilterItem, error) {
	// get cache
	if cache := u.getShopPageFilterCache(*shopId); cache != nil {
		return cache, nil
	}

	// get shop tag types
	filter := db.ShopTagTypeFilter{
		ShopIds: &[]uuid.UUID{*shopId},
		Active:  pointer.BoolPtr(true),
		Sources: &[]string{"filter"},
		DataConfig: &entity.DataConfig{
			IsCount:        pointer.BoolPtr(false),
			IsUpdateFilter: pointer.BoolPtr(true),
			IsKeepIdsOrder: pointer.BoolPtr(true),
		},
	}

	shopTagTypes, err := u.shopTagType.List(ctx, &filter)
	if err != nil {
		return nil, err
	}

	shopPageFilters := make(map[uint64]pages.ShopPageFilterItem)

	for _, shopTagType := range *shopTagTypes {
		shopPageFilter := pages.ShopPageFilterItem{
			Id: shopTagType.TagTypeId,
		}

		// get tag type
		tagType, err := u.tagType.Get(ctx, &db.TagTypeFilter{
			Ids: &[]uuid.UUID{shopTagType.TagTypeId},
		})
		if err != nil {
			return nil, err
		}

		shopPageFilter.Name = tagType.Name
		shopPageFilter.Url = tagType.Url

		// get tag select
		tagSelects, err := u.tagSelect.List(ctx, &db.TagSelectFilter{
			TagTypeIds: &[]uuid.UUID{shopTagType.TagTypeId},
			Active:     pointer.BoolPtr(true),
		})

		shopPageFilter.Select = make(map[uint64]pages.ShopPageFilterSelectItem)
		for _, tagSelect := range *tagSelects {
			shopPageFilterSelectItem := pages.ShopPageFilterSelectItem{
				Id:     tagSelect.Id,
				Name:   tagSelect.Name,
				Url:    tagSelect.Url,
				Active: false,
			}

			shopPageFilter.Select[tagSelect.SortOrder] = shopPageFilterSelectItem
		}

		shopPageFilters[shopTagType.SortOrder] = shopPageFilter
	}

	// set cache
	go u.setShopPageFilterCache(*shopId, &shopPageFilters)

	// return
	return &shopPageFilters, nil
}
