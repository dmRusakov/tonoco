package admin_app_web_v1

import (
	"github.com/dmRusakov/tonoco/internal/config"
	productUseCase "github.com/dmRusakov/tonoco/internal/useCase/shop_page"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
)

func NewWebServer(
	cfg *config.Config,
	productUseCase *productUseCase.UseCase,
	appCacheService *appCacheService.Service,
	userCacheService *userCacheService.Service,
) (*Controller, error) {
	return &Controller{
		cfg:              cfg,
		tmlPath:          "./assets/templates/",
		shopPageUseCase:  productUseCase,
		appCacheService:  appCacheService,
		userCacheService: userCacheService,
	}, nil
}
