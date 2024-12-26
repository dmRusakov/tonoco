package appInit

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
)

type AppInit interface {
	AppCacheServiceInit() (err error)
	UserCacheServiceInit() (err error)
	ProductAPIInit() (err error)
}

func NewAppInit(ctx context.Context, cfg *config.Config) *App {
	return &App{
		Ctx: ctx,
		Cfg: cfg,
	}
}
