package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/entity"
	productUseCase "github.com/dmRusakov/tonoco/internal/useCase/product"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"net/http"
)

var _ Server = &Controller{}

type Controller struct {
	cfg              *config.Config
	tmlPath          string
	productUseCase   *productUseCase.UseCase
	appCacheService  *appCacheService.Service
	userCacheService *userCacheService.Service
}

type Server interface {
	Render(w http.ResponseWriter, t string, appData entity.AppData)
	Start(ctx context.Context, cfg *config.Config) error

	RenderProducts(ctx context.Context, w http.ResponseWriter, r *http.Request, appData entity.AppData)
}
