package product

import (
	"github.com/dmRusakov/tonoco/internal/domain/product/service"
	"github.com/dmRusakov/tonoco/pkg/common/core/clock"
	"time"
)

type IdentityGenerator interface {
}

type Clock interface {
	Now() time.Time
}

type Policy struct {
	productService *service.ProductService

	identity IdentityGenerator
	clock    Clock
}

func NewProductPolicy(productService *service.ProductService, identity IdentityGenerator, clock clock.Clock) *Policy {
	return &Policy{
		productService: productService,

		identity: identity,
		clock:    clock,
	}
}
