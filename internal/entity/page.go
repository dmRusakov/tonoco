package entity

import "github.com/google/uuid"

type ProductPage struct {
	Name             string
	ShortDescription string
	Description      string
	Products         *map[uuid.UUID]*ProductListItem
	ProductUrl       string
	CountItems       *uint64
	ConsoleMessage   ConsoleMessage `json:"console_message" db:"console_message"`
}
