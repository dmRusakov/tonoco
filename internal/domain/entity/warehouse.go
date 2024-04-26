package entity

import "time"

type Warehouse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	SortOrder    uint64    `json:"sort_order"`
	Active       bool      `json:"active"`
	AddressLine1 string    `json:"address_line1"`
	AddressLine2 string    `json:"address_line2"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	ZipCode      string    `json:"zip_code"`
	Country      string    `json:"country"`
	WebSite      string    `json:"web_site"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
}

type WarehouseFilter struct {
	IDs           *[]string `json:"ids"`
	Abbreviations *[]string `json:"abbreviations"`
	Active        *bool     `json:"active"`
	Prime         *bool     `json:"prime"`
	Search        *string   `json:"search"`
	OrderBy       *string   `json:"order_by"`
	OrderDir      *string   `json:"order_dir"`
	Page          *uint64   `json:"page"`
	PerPage       *uint64   `json:"per_page"`
}
