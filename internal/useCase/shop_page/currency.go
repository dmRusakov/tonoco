package shop_page

import (
	"context"
	"github.com/dmRusakov/tonoco/internal/entity/db"
	"github.com/dmRusakov/tonoco/internal/entity/pages"
)

func (u *UseCase) getCurrency(ctx context.Context, parameters *pages.ProductsPageUrlParams, errs *[]error) *db.Currency {
	var currency *db.Currency
	var err error

	// get default currency
	defaultCurrency := u.currency.GetDefault()
	if (parameters.Currency == nil) || (*parameters.Currency == "" || *parameters.Currency == defaultCurrency.Url) {
		currency = defaultCurrency
	} else {
		// get currency by url
		currency, err = u.currency.Get(ctx, &db.CurrencyFilter{
			Urls: &[]string{*parameters.Currency},
		})
		if err != nil {
			*errs = append(*errs, err)
		}
	}
	return currency
}
