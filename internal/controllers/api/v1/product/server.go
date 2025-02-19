package product

import (
	productService "github.com/dmRusakov/tonoco-grpc/gen/go/proto/service/v1"
	productUsecase "github.com/dmRusakov/tonoco/internal/useCase/shop_page"
)

type Server struct {
	useCase *productUsecase.UseCase
	productService.UnimplementedProductServiceServer
}

func NewServer(
	useCase *productUsecase.UseCase,
	srv productService.UnimplementedProductServiceServer,
) *Server {
	return &Server{
		useCase:                           useCase,
		UnimplementedProductServiceServer: srv,
	}
}
