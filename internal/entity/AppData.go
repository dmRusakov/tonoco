package entity

type AppData struct {
	IsProd         bool           `json:"is_prod" db:"is_prod"`
	IsDebug        bool           `json:"is_debug" db:"is_debug"`
	ConsoleMessage ConsoleMessage `json:"console_message" db:"console_message"`
}
