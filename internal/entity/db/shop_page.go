package db

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
	"time"
)

type ShopPage struct {
	Id               uuid.UUID
	Name             uuid.UUID
	SeoTitle         uuid.UUID
	ShortDescription uuid.UUID
	Description      uuid.UUID
	Url              string
	ImageId          uuid.UUID
	HoverImageId     uuid.UUID
	Page             uint64
	PerPage          uint64
	SortOrder        int
	Active           bool
	Prime            bool
	CreatedAt        time.Time
	CreatedBy        uuid.UUID
	UpdatedAt        time.Time
	UpdatedBy        uuid.UUID
}

type ShopPageFilter struct {
	Ids  *[]uuid.UUID `json:"ids"`
	Urls *[]string    `json:"urls"`

	Active *bool   `json:"active"`
	Prime  *bool   `json:"prime"`
	Search *string `json:"search"`

	DataPagination *entity.DataPagination
	DataConfig     *entity.DataConfig
}
