package service

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/domain/folder/model"
	"github.com/dmRusakov/tonoco/internal/entity"
	"time"
)

type Item = entity.Folder
type Filter = entity.FolderFilter

type repository interface {
	Get(context.Context, *Filter) (*Item, []error)
	List(context.Context, *Filter, bool) (*map[string]Item, []error)
	Create(context.Context, *Item) (*string, []error)
	Update(context.Context, *Item) []error
	Patch(context.Context, *string, *map[string]interface{}) []error
	UpdatedAt(context.Context, *string) (*time.Time, []error)
	TableIndexCount(context.Context) (*uint64, []error)
	MaxSortOrder(context.Context) (*uint64, []error)
	Delete(context.Context, *string) []error
}

type Service struct {
	repository model.Storage
	itemCash   map[string]Item
	itemsCash  map[string]map[string]Item
}

func NewService(repository model.Storage) *Service {
	return &Service{
		repository: repository,
		itemCash:   make(map[string]Item),
		itemsCash:  make(map[string]map[string]Item),
	}
}
