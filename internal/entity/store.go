package entity

import "time"

type Store struct {
	Id           string    `json:"id" db:"id"`
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
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	CreatedBy    string    `json:"created_by" db:"created_by"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy    string    `json:"updated_by" db:"updated_by"`
}

type StoreFilter struct {
	Ids           *[]string `json:"ids"`
	Urls          *[]string `json:"urls"`
	Abbreviations *[]string `json:"abbreviations"`

	Active *bool   `json:"active"`
	Search *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
