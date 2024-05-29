package entity

import "time"

type TagType struct {
	ID               string    `json:"id" sql:"id"`
	Name             string    `json:"name" sql:"name"`
	Url              string    `json:"url" sql:"url"`
	ShortDescription string    `json:"short_description" sql:"short_description"`
	Description      string    `json:"description" sql:"description"`
	Required         bool      `json:"required" sql:"required"`
	Active           bool      `json:"active" sql:"active"`
	Prime            bool      `json:"prime" sql:"prime"`
	ListItem         bool      `json:"list_item" sql:"list_item"`
	Filter           bool      `json:"filter" sql:"filter"`
	SortOrder        uint64    `json:"sort_order" sql:"sort_order"`
	Type             string    `json:"type" sql:"type"`
	Prefix           string    `json:"prefix" sql:"prefix"`
	Suffix           string    `json:"suffix" sql:"suffix"`
	CreatedAt        time.Time `json:"created_at" sql:"created_at"`
	CreatedBy        string    `json:"created_by" sql:"created_by"`
	UpdatedAt        time.Time `json:"updated_at" sql:"updated_at"`
	UpdatedBy        string    `json:"updated_by" sql:"updated_by"`
}

type TagTypeFilter struct {
	IDs      *[]string
	URLs     *[]string
	Required *bool
	Active   *bool
	Prime    *bool
	ListItem *bool
	Filter   *bool
	Type     *string
	Search   *string
	OrderBy  *string
	OrderDir *string
	Page     *uint64
	PerPage  *uint64
}
