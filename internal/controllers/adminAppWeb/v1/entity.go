package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	productUseCase "github.com/dmRusakov/tonoco/internal/useCase/product"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"net/http"
)

var _ Server = &Service{}

type Service struct {
	cfg              *config.Config
	tmlPath          string
	productUseCase   *productUseCase.UseCase
	appCacheService  *appCacheService.Service
	userCacheService *userCacheService.Service
}

type Server interface {
	Render(w http.ResponseWriter, t string)
	Start(ctx context.Context) error
}
