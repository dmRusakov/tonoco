package entity

import "github.com/google/uuid"

type ProductPage struct {
	Name             string
	ShortDescription string
	Description      string
	Items            *map[uuid.UUID]*ProductListItem
	Url              string
	Page             uint64
	PerPage          uint64
	TotalPages       uint64
	TotalItems       uint64
	Pagination       map[uint64]string
	ConsoleMessage   ConsoleMessage
}
