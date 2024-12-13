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
	parameters *pages.ProductsPageUrlParams,
	shopId *uuid.UUID,
) (*pages.ShopPageFilter, error) {
	var shopPageFilter pages.ShopPageFilter
	urlMapping := make(map[string]uuid.UUID)
	shopPageFilter.TagUrlMap = &urlMapping
	tagOrder := make(map[uint64]uuid.UUID)
	shopPageFilter.TagOrder = &tagOrder
	tagSelects := make(map[uuid.UUID]map[uuid.UUID]db.TagSelect)
	shopPageFilter.TagSelect = &tagSelects

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
	for _, item := range *shopTagTypes {
		(*shopPageFilter.TagOrder)[item.SortOrder] = item.TagTypeId
	}

	// get tag types
	shopPageFilter.TagTypes, err = u.tagType.List(ctx, &(db.TagTypeFilter{
		Ids:    filter.TagTypeIds,
		Active: pointer.BoolPtr(true),
	}))
	for _, item := range *shopPageFilter.TagTypes {
		(*shopPageFilter.TagUrlMap)[item.Url] = item.Id
	}

	if err != nil {
		return nil, err
	}

	// get tag select
	for _, tagTypesId := range *shopPageFilter.TagOrder {
		(*shopPageFilter.TagSelect)[tagTypesId] = make(map[uuid.UUID]db.TagSelect)
		data, e := u.tagSelect.List(ctx, &(db.TagSelectFilter{
			TagTypeIds: &[]uuid.UUID{tagTypesId},
			Active:     pointer.BoolPtr(true),
		}))
		if e != nil {
			return nil, err
		}

		for _, item := range *data {
			(*shopPageFilter.TagSelect)[tagTypesId][item.Id] = item
		}
	}

	return &shopPageFilter, nil
}
