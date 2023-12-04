package v1

var _ ProductStatus = &productServer{}

type productServer struct {
}

type ProductStatus interface {
}

func NewProductServer() *productServer {
	return &productServer{}
}
