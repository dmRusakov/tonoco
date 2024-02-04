package api_v1

import (
	productService "github.com/dmRusakov/tonoco-grpc/gen/go/proto/service/v1"
)

type ProductController struct {
	client productService.UnimplementedProductServiceServer
}

func NewProductController(c productService.UnimplementedProductServiceServer) *ProductController {
	return &ProductController{client: c}
}
