package currency

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/google/uuid"
)

func (u *UseCase) Get(ctx context.Context, id *uuid.UUID, url *string) (*entity.Currency, error) {
	var filter *entity.CurrencyFilter
	if id != nil {
		filter.Ids = &[]uuid.UUID{*id}
	}

	if url != nil {
		filter.Urls = &[]string{*url}
	}
	return u.currency.Get(ctx, filter)
}

func (u *UseCase) GetDefault() (*entity.Currency, error) {
	return u.currency.DefaultCurrency, nil
}
