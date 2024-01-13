package product

import (
	"context"
	pb_prod_products "github.com/dmRusakov/tonoco-grpc/gen/go/proto/tonoco/product/v1"
	"github.com/dmRusakov/tonoco/internal/domain/policy/product"
)

type Server struct {
	policy *product.Policy
	pb_prod_products.UnimplementedTonocoProductServiceServer
}

func NewServer(
	policy *product.Policy,
	srv pb_prod_products.UnimplementedTonocoProductServiceServer) *Server {
	return &Server{
		policy:                                  policy,
		UnimplementedTonocoProductServiceServer: srv,
	}
}

func (s *Server) GetProduct(
	ctx context.Context,
	request *pb_prod_products.GetProductRequest,
) (*pb_prod_products.GetProductResponse, error) {
	// TODO: implement

	return nil, nil
}

func (s *Server) ListProducts(ctx context.Context, request *pb_prod_products.ListProductsRequest) (*pb_prod_products.ListProductsResponse, error) {
	// TODO: implement

	return nil, nil
}
func (s *Server) CreateProduct(ctx context.Context, request *pb_prod_products.CreateProductRequest) (*pb_prod_products.CreateProductResponse, error) {
	// TODO: implement

	return nil, nil
}
func (s *Server) UpdateProduct(ctx context.Context, request *pb_prod_products.UpdateProductRequest) (*pb_prod_products.UpdateProductResponse, error) {
	// TODO: implement

	return nil, nil
}
func (s *Server) PatchProduct(ctx context.Context, request *pb_prod_products.PatchProductRequest) (*pb_prod_products.PatchProductResponse, error) {
	// TODO: implement

	return nil, nil
}
func (s *Server) DeleteProduct(ctx context.Context, request *pb_prod_products.DeleteProductRequest) (*pb_prod_products.DeleteProductResponse, error) {
	// TODO: implement

	return nil, nil
}
