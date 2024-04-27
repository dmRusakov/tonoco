package admin_app_web_v1

import (
	"context"
	productUseCase "github.com/dmRusakov/tonoco/internal/useCase/product"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"net/http"
)

var _ Server = &server{}

type server struct {
	tmlPath          string
	productUseCase   *productUseCase.UseCase
	appCacheService  *appCacheService.Service
	userCacheService *userCacheService.Service
}

type Server interface {
	Render(w http.ResponseWriter, t string)
	Start(ctx context.Context, port string) error
}
