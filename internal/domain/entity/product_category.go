package entity

type ProductCategory struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	Order            int32  `json:"order"`
	Prime            bool   `json:"prime"`
	Active           bool   `json:"active"`
	CreatedAt        int64  `json:"created_at"`
	CreatedBy        string `json:"created_by"`
	UpdatedAt        int64  `json:"updated_at"`
	UpdatedBy        string `json:"updated_by"`
}
