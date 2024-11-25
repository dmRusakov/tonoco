package db

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
	"time"
)

type Text struct {
	Id       uuid.UUID
	Language string

	Source   string
	SourceId uuid.UUID

	Text   string
	Active bool

	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
}

type TextFilter struct {
	Ids      *[]uuid.UUID
	Language *string

	Source   *[]string
	SourceId *[]uuid.UUID

	Active *bool
	Search *string

	DataPagination *entity.DataPagination
	DataConfig     *entity.DataConfig
}
