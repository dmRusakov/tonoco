package product

import (
	"context"
	productService "github.com/dmRusakov/tonoco-grpc/gen/go/proto/service/v1"
)

func (p *Server) GetProduct(ctx context.Context, id string) (*productService.GetProductResponse, error) {
	return nil, nil
}
