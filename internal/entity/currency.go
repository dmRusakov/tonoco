package entity

import (
	"github.com/google/uuid"
	"time"
)

type Currency struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Symbol    string    `json:"symbol" db:"symbol"`
	Url       string    `json:"url" db:"url"`
	SortOrder uint64    `json:"sort_order" db:"sort_order"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	CreatedBy uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy uuid.UUID `db:"updated_by" json:"updated_by"`
}

type CurrencyFilter struct {
	Ids  *[]uuid.UUID `json:"ids"`
	Urls *[]string    `json:"urls"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
