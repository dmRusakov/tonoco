package entity

type ShippingClass struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Order     uint32 `json:"order"`
	Active    bool   `json:"active"`
	CreatedAt uint32 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt uint32 `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}
