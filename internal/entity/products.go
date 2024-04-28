package entity

type ProductListItem struct {
	ID               string `json:"id"`
	SKU              string `json:"sku"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Url              string `json:"url"`
	IsTaxable        bool   `json:"is_taxable"`
	SeoTitle         string `json:"seo_title"`
	SeoDescription   string `json:"seo_description"`
}

type ProductsUrlParameters struct {
	ID       string
	Category string
}
