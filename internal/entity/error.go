package entity

import "time"

type Error struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Message    string    `json:"message"`
	DevMessage string    `json:"dev_message"`
	Field      string    `json:"field"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  string    `json:"updated_by"`
}

type ErrorFilter struct {
	IDs    *[]string `json:"ids"`
	Codes  *[]string `json:"codes"`
	Active *bool     `json:"active"`
	Type   *string   `json:"type"`
	Search *string   `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`
}
