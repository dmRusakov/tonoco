package entity

import "github.com/dmRusakov/tonoco/internal/entity/pages"

type AppData struct {
	IsProd         bool                 `json:"is_prod" db:"is_prod"`
	IsDebug        bool                 `json:"is_debug" db:"is_debug"`
	ConsoleMessage pages.ConsoleMessage `json:"console_message" db:"console_message"`
}
