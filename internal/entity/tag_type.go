package entity

import "time"

type TagType struct {
	Id               string    `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Url              string    `json:"url" db:"url"`
	ShortDescription string    `json:"short_description" db:"short_description"`
	Description      string    `json:"description" db:"description"`
	Required         bool      `json:"required" db:"required"`
	Active           bool      `json:"active" db:"active"`
	Prime            bool      `json:"prime" db:"prime"`
	ListItem         bool      `json:"list_item" db:"list_item"`
	Filter           bool      `json:"filter" db:"filter"`
	SortOrder        uint64    `json:"sort_order" db:"sort_order"`
	Type             string    `json:"type" db:"type"`
	Prefix           string    `json:"prefix" db:"prefix"`
	Suffix           string    `json:"suffix" db:"suffix"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	CreatedBy        string    `json:"created_by" db:"created_by"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy        string    `json:"updated_by" db:"updated_by"`
}

type TagTypeFilter struct {
	Ids  *[]string
	Urls *[]string

	Required *bool
	Active   *bool
	Prime    *bool
	ListItem *bool
	Filter   *bool
	Type     *string

	Search *string

	OrderBy  *string
	OrderDir *string
	Page     *uint64
	PerPage  *uint64

	IsCount *bool `json:"is_count"`
}
