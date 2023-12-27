package appInit

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/config"
)

func NewAppInit(ctx *context.Context, cfg *config.Config) *App {
	return &App{
		Ctx: ctx,
		Cfg: cfg,
	}
}
