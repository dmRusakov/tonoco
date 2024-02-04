package api_v1

import (
	"context"
	productService "github.com/dmRusakov/tonoco-grpc/gen/go/proto/service/v1"
)

func (p *ProductController) GetProduct(ctx context.Context, id string) (*productService.GetProductResponse, error) {
	return p.client.GetProduct(ctx, &productService.GetProductRequest{Id: id})
}
