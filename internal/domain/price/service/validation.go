package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoProductID   = validation.NewValidation("Product Id is required", "Product Id is required", "934722", "ProductIds")
	NoPriceTypeID = validation.NewValidation("Price Type Id is required", "Price Type Id is required", "934723", "PriceTypeIds")
	NoCurrencyID  = validation.NewValidation("Currency Id is required", "Currency Id is required", "934724", "CurrencyIds")
	NoPrice       = validation.NewValidation("Price is required", "Price is required", "934725", "Price")
)

func (s *Service) Validate(item *Item, fields *map[string]interface{}) []entity.Error {
	var validations []entity.Error

	// validate Item
	if item != nil {
		results := make(chan []entity.Error)

		// Start goroutines for validateName and validateUrl
		go func() {
			var validations []entity.Error
			s.validateProductID(item.ProductID, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Error
			s.validatePriceTypeID(item.PriceTypeID, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Error
			s.validateCurrencyID(item.CurrencyID, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Error
			s.validatePrice(item.Price, &validations)
			results <- validations
		}()

		// Collect results from the channel
		for i := 0; i < 4; i++ {
			v := <-results
			validations = append(validations, v...)
		}
	}

	// validate fields using goroutines and switch
	if fields != nil {
		results := make(chan []entity.Error)

		for field, value := range *fields {
			go func(field string, value interface{}) {
				var v []entity.Error
				switch field {
				case "ProductIds":
					s.validateProductID(value.(string), &v)
					break
				case "PriceTypeIds":
					s.validatePriceTypeID(value.(string), &v)
					break
				case "CurrencyIds":
					s.validateCurrencyID(value.(string), &v)
					break
				case "Price":
					s.validatePrice(value.(float64), &v)
					break
				}
				results <- v
			}(field, value)
		}

		// Collect results from the channel
		for range *fields {
			v := <-results
			validations = append(validations, v...)
		}
	}

	return validations
}

func (s *Service) validateProductID(productID string, validations *[]entity.Error) {
	if productID == "" {
		*validations = append(*validations, NoProductID)
	}
}

func (s *Service) validatePriceTypeID(priceTypeID string, validations *[]entity.Error) {
	if priceTypeID == "" {
		*validations = append(*validations, NoPriceTypeID)
	}
}

func (s *Service) validateCurrencyID(currencyID string, validations *[]entity.Error) {
	if currencyID == "" {
		*validations = append(*validations, NoCurrencyID)
	}
}

func (s *Service) validatePrice(price float64, validations *[]entity.Error) {
	if price == 0 {
		*validations = append(*validations, NoPrice)
	}
}
