package entity

import "github.com/google/uuid"

type ProductPage struct {
	Name             string
	ShortDescription string
	Description      string
	Items            *map[uuid.UUID]*ProductListItem
	Url              string
	Page             int
	PerPage          int
	TotalPages       int
	TotalItems       int
	ConsoleMessage   ConsoleMessage
}
