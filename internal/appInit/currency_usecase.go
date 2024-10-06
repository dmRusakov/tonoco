package appInit

import (
	currency_usecase "github.com/dmRusakov/tonoco/internal/useCase/currency"
)

func (a *App) CurrencyUseCaseInit() error { // currency
	a.CurrencyUseCase = currency_usecase.NewUseCase(a.Services.Currency)

	return nil
}
