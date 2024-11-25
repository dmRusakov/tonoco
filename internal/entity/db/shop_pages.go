package db

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
	"time"
)

type StorePage struct {
	Id               uuid.UUID
	Name             uuid.UUID
	SeoTitle         uuid.UUID
	ShortDescription uuid.UUID
	Description      uuid.UUID
	Url              string
	ImageId          uuid.UUID
	HoverImageId     uuid.UUID
	Page             int
	PerPage          int
	SortOrder        int
	Active           bool
	CreatedAt        time.Time
	CreatedBy        uuid.UUID
	UpdatedAt        time.Time
	UpdatedBy        uuid.UUID
}

type StorePageFilter struct {
	Ids  *[]uuid.UUID `json:"ids"`
	Urls *[]string    `json:"urls"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	DataPagination *entity.DataPagination
	DataConfig     *entity.DataConfig
}
