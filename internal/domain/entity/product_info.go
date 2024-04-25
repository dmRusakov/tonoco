package entity

type Product struct {
	Product        ProductInfo
	Status         ProductStatus
	Categories     []ProductCategory
	Specifications []Specification
}
