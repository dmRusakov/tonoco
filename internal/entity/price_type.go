package entity

import "time"

type PriceType struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Url       string    `json:"url" db:"url"`
	SortOrder uint64    `json:"sort_order" db:"sort_order"`
	IsPublic  bool      `json:"is_public" db:"is_public"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy string    `json:"updated_by" db:"updated_by"`
}

type PriceTypeFilter struct {
	Ids  *[]string `json:"ids"`
	Urls *[]string `json:"urls"`

	Active   *bool   `json:"active"`
	IsPublic *bool   `json:"prime"`
	Search   *string `json:"search"`

	OrderBy  *string `json:"order_by"`
	OrderDir *string `json:"order_dir"`
	Page     *uint64 `json:"page"`
	PerPage  *uint64 `json:"per_page"`

	IsCount *bool `json:"is_count"`
}
