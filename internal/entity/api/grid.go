package api

type GridParam struct {
	Id  string `json:"id"`
	Sku string `json:"sku"`
}

type GridItem struct {
	Id               string                 `json:"id"`
	No               int32                  `json:"no"`
	Sku              string                 `json:"sku"`
	Brand            string                 `json:"brand"`
	Name             string                 `json:"name"`
	ShortDescription string                 `json:"short_description"`
	Url              string                 `json:"url"`
	SalePrice        string                 `json:"sale_price"`
	Price            string                 `json:"price"`
	Currency         string                 `json:"currency" `
	Quantity         int64                  `json:"quantity"`
	Status           string                 `json:"status"`
	IsTaxable        bool                   `json:"is_taxable"`
	SeoTitle         string                 `json:"seo_title"`
	SeoDescription   string                 `json:"seo_description"`
	Categories       map[string]interface{} `json:"categories" db:"categories"`
	Tags             map[string]interface{} `json:"tags" db:"tags"`
	MainImage        map[string]interface{} `json:"main_image" db:"main_image"`
	HoverImage       map[string]interface{} `json:"hover_image" db:"hover_image"`
}
