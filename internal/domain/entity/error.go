package entity

type Error struct {
	Message    string   `json:"message"`
	DevMessage string   `json:"dev_message"`
	Field      []string `json:"field"`
	Code       string   `json:"code"`
}
