package entity

import "github.com/google/uuid"

type ProductPage struct {
	Name             string
	ShortDescription string
	Description      string
	Products         *map[uuid.UUID]*ProductListItem
	ProductUrl       string
	Page             *uint64
	CountItems       *uint64
	ConsoleMessage   ConsoleMessage `json:"console_message" db:"console_message"`
}
