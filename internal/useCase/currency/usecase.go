package currency

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/google/uuid"
)

func (u *UseCase) Get(ctx context.Context, id *uuid.UUID, url *string) (*db.Currency, error) {
	var filter *db.CurrencyFilter
	if id != nil {
		filter.Ids = &[]uuid.UUID{*id}
	}

	if url != nil {
		filter.Urls = &[]string{*url}
	}
	return u.currency.Get(ctx, filter)
}

func (u *UseCase) GetDefault() (*db.Currency, error) {
	return u.currency.GetDefault(), nil
}
