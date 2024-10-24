package entity

import "github.com/google/uuid"

type ProductPage struct {
	Name             string
	ShortDescription string
	Description      string
	Items            *map[uuid.UUID]*ProductListItem
	Url              string
	Page             uint32
	PerPage          uint32
	TotalPages       uint32
	TotalItems       uint32
	Pagination       []uint32
	ConsoleMessage   ConsoleMessage
}
