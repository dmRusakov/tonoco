package entity

type SpecificationValue struct {
	ID              string `json:"id"`
	SpecificationID string `json:"specification_id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Active          bool   `json:"active"`
	Order           int32  `json:"order"`
	CreatedAt       int64  `json:"created_at"`
	CreatedBy       string `json:"created_by"`
	UpdatedAt       int64  `json:"updated_at"`
	UpdatedBy       string `json:"updated_by"`
}