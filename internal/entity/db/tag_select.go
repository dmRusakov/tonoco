package db

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
	"time"
)

type TagSelect struct {
	Id        uuid.UUID
	TagTypeId uuid.UUID
	Name      string
	Url       string
	Active    bool
	SortOrder uint64
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
}

type TagSelectFilter struct {
	Ids        *[]uuid.UUID
	Urls       *[]string
	TagTypeIds *[]uuid.UUID
	Active     *bool
	Search     *string

	DataPagination *entity.DataPagination
	DataConfig     *entity.DataConfig
}
