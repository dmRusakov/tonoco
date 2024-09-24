package entity

import (
	"github.com/google/uuid"
	"time"
)

type Store struct {
	Id           uuid.UUID `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Url          string    `json:"url" db:"url"`
	Abbreviation string    `json:"abbreviation" db:"abbreviation"`
	SortOrder    uint64    `json:"sort_order" db:"sort_order"`
	Active       bool      `json:"active" db:"active"`
	AddressLine1 string    `json:"address_line1" db:"address_line_1"`
	AddressLine2 string    `json:"address_line2" db:"address_line_2"`
	City         string    `json:"city" db:"city"`
	State        string    `json:"state" db:"state"`
	ZipCode      string    `json:"zip" db:"zip"`
	Country      string    `json:"country" db:"country"`
	WebSite      string    `json:"web_site" db:"web_site"`
	Phone        string    `json:"phone" db:"phone"`
	Email        string    `json:"email" db:"email"`
	CurrencyUrl  string    `json:"currency_url" db:"currency_url"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	CreatedBy    uuid.UUID `db:"created_by" json:"created_by"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy    uuid.UUID `db:"updated_by" json:"updated_by"`
}

type StoreFilter struct {
	Ids           *[]uuid.UUID `json:"ids"`
	Urls          *[]string    `json:"urls"`
	Abbreviations *[]string    `json:"abbreviations"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	Count *uint64 `json:"count"`

	IsIdsOnly          *bool `json:"is_ids_only"`
	IsCount            *bool `json:"is_count"`
	IsUpdateFilter     *bool `json:"is_update_filter"`
	IsRemoveDuplicates *bool `json:"is_remove_duplicates"`
}
