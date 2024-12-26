package db

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
	"time"
)

// ShopTagType represents the shop_tag_type table in the database.
type ShopTagType struct {
	Id        uuid.UUID
	ShopId    uuid.UUID
	TagTypeId uuid.UUID
	Source    string
	SortOrder uint64
	Active    bool
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
}

// ShopTagTypeFilter represents the filter criteria for querying shop_tag_type records.
type ShopTagTypeFilter struct {
	Ids        *[]uuid.UUID
	ShopIds    *[]uuid.UUID
	TagTypeIds *[]uuid.UUID
	Sources    *[]string
	SortOrders *[]uint64
	Active     *bool
	Search     *string

	DataPagination *entity.DataPagination
	DataConfig     *entity.DataConfig
}
