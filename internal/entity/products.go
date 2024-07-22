package entity

type ProductListItem struct {
	No               int32    `json:"no"`
	Id               string   `json:"id"`
	Sku              string   `json:"sku"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Url              string   `json:"url"`
	Price            string   `json:"price"`
	Currency         string   `json:"currency"`
	Quantity         int32    `json:"quantity"`
	IsTaxable        bool     `json:"is_taxable"`
	SeoTitle         string   `json:"seo_title"`
	SeoDescription   string   `json:"seo_description"`
	Categories       []string `json:"categories"`
}

type ProductsPageUrlParams struct {
	Id       *string
	Category *string
	Currency *string
	Page     *uint64
	PerPage  *uint64
}
