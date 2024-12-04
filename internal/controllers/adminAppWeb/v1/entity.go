package admin_app_web_v1

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
	"github.com/dmRusakov/tonoco/internal/useCase/shop_page"
	"github.com/dmRusakov/tonoco/pkg/appCacheService"
	"github.com/dmRusakov/tonoco/pkg/userCacheService"
	"net/http"
)

var _ Server = &Controller{}

type Controller struct {
	cfg  *config.Config
	rout []struct {
		path    string
		handler func(http.ResponseWriter, *http.Request)
	}
	tmlPath          string
	shopPageUseCase  *shop_page.UseCase
	appCacheService  *appCacheService.Service
	userCacheService *userCacheService.Service
}

type Server interface {
	Render(w http.ResponseWriter, t string)
	Start(ctx context.Context) error

	RenderShopPage(context.Context, http.ResponseWriter, *http.Request, string)
}
