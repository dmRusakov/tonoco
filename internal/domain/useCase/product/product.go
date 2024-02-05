package product

import (
	"github.com/dmRusakov/tonoco/internal/domain/product/service"
	"github.com/dmRusakov/tonoco/pkg/common/core/clock"
	"time"
)

type IdentityGenerator interface {
	GenerateUUIDv4String() string
}

type Clock interface {
	Now() time.Time
}

type UseCase struct {
	productService *service.ProductService

	identity IdentityGenerator
	clock    Clock
}

func NewProductUseCase(productService *service.ProductService, identity IdentityGenerator, clock clock.Clock) *UseCase {
	return &UseCase{
		productService: productService,

		identity: identity,
		clock:    clock,
	}
}
