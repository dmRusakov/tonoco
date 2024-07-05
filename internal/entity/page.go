package entity

type ProductPage struct {
	Name             string
	ShortDescription string
	Description      string
	Products         *map[string]ProductListItem
	ProductUrl       string
}
