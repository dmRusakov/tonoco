package product

import (
	pb_prod_products "github.com/dmRusakov/tonoco-grpc/gen/go/proto/tonoco/product/v1"
	"github.com/dmRusakov/tonoco/internal/domain/policy/product"
)

type Server struct {
	policy *product.Policy
	pb_prod_products.UnimplementedTonocoProductServiceServer
}

func NewServer(policy *product.Policy, srv pb_prod_products.UnimplementedTonocoProductServiceServer) *Server {
	return &Server{
		policy:                                  policy,
		UnimplementedTonocoProductServiceServer: srv,
	}
}
