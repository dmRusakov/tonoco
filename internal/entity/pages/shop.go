package pages

import (
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/google/uuid"
	"time"
)

type Shop struct {
	Id               uuid.UUID
	Name             string
	SeoTitle         string
	ShortDescription string
	Description      string
	Url              string
	Image            *db.Image
	HoverImage       *db.Image
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
	TagUrlMap *map[string]uuid.UUID                     // map[url]tagTypeId
	TagTypes  *map[uuid.UUID]db.TagType                 // map[tagTypeId]tagType
	TagOrder  *map[uint64]uuid.UUID                     // map[sortOrder]tagTypeId
	TagSelect *map[uuid.UUID]map[uuid.UUID]db.TagSelect // map[tagTypeId]map[tagSelectId]tagSelect
}
