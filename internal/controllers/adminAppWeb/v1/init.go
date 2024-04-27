package admin_app_web_v1

import (
	productUseCase "github.com/dmRusakov/tonoco/internal/useCase/product"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
)

func NewWebServer(
	productUseCase *productUseCase.UseCase,
	appCacheService *appCacheService.Service,
	userCacheService *userCacheService.Service,
) (*server, error) {
	return &server{
		tmlPath:          "./assets/templates/",
		productUseCase:   productUseCase,
		appCacheService:  appCacheService,
		userCacheService: userCacheService,
	}, nil
}
