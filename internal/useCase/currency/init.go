package currency

import (
	currency_service "github.com/dmRusakov/tonoco/internal/domain/currency/service"
)

type UseCase struct {
	currency *currency_service.Service
}

func NewUseCase(currencyService *currency_service.Service) *UseCase {
	return &UseCase{
		currency: currencyService,
	}
}
