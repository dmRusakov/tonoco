package entity

import (
	"github.com/google/uuid"
	"time"
)

type Error struct {
	Id         uuid.UUID `json:"id"`
	Type       string    `json:"type"`
	Message    string    `json:"message"`
	DevMessage string    `json:"dev_message"`
	Field      string    `json:"field"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	CreatedBy  uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy  uuid.UUID `db:"updated_by" json:"updated_by"`
}

type ErrorFilter struct {
	Ids    *[]uuid.UUID `json:"ids"`
	Codes  *[]string    `json:"codes"`
	Active *bool        `json:"active"`
	Type   *string      `json:"type"`
	Search *string      `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	Count *uint64 `json:"count"`

	IsIdsOnly      *bool `json:"is_ids_only"`
	IsCount        *bool `json:"is_count"`
	IsUpdateFilter *bool `json:"is_update_filter"`
	IsKeepIdsOrder *bool `json:"is_keep_ids_order"`
}

type ConsoleMessage struct {
	Log   []string
	Warn  []string
	Error []string
}
