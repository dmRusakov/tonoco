package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoProductID   = validation.NewValidation("Product ID is required", "Product ID is required", "934722", "ProductID")
	NoPriceTypeID = validation.NewValidation("Price Type ID is required", "Price Type ID is required", "934723", "PriceTypeID")
	NoCurrencyID  = validation.NewValidation("Currency ID is required", "Currency ID is required", "934724", "CurrencyID")
	NoPrice       = validation.NewValidation("Price is required", "Price is required", "934725", "Price")
)

func (s *Service) Validate(item *Item, fields *map[string]interface{}) []entity.Validation {
	var validations []entity.Validation

	// validate Item
	if item != nil {
		results := make(chan []entity.Validation)

		// Start goroutines for validateName and validateUrl
		go func() {
			var validations []entity.Validation
			s.validateProductID(item.ProductID, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Validation
			s.validatePriceTypeID(item.PriceTypeID, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Validation
			s.validateCurrencyID(item.CurrencyID, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Validation
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
		results := make(chan []entity.Validation)

		for field, value := range *fields {
			go func(field string, value interface{}) {
				var v []entity.Validation
				switch field {
				case "ProductID":
					s.validateProductID(value.(string), &v)
					break
				case "PriceTypeID":
					s.validatePriceTypeID(value.(string), &v)
					break
				case "CurrencyID":
					s.validateCurrencyID(value.(string), &v)
					break
				case "Price":
					s.validatePrice(value.(uint64), &v)
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

func (s *Service) validateProductID(productID string, validations *[]entity.Validation) {
	if productID == "" {
		*validations = append(*validations, NoProductID)
	}
}

func (s *Service) validatePriceTypeID(priceTypeID string, validations *[]entity.Validation) {
	if priceTypeID == "" {
		*validations = append(*validations, NoPriceTypeID)
	}
}

func (s *Service) validateCurrencyID(currencyID string, validations *[]entity.Validation) {
	if currencyID == "" {
		*validations = append(*validations, NoCurrencyID)
	}
}

func (s *Service) validatePrice(price uint64, validations *[]entity.Validation) {
	if price == 0 {
		*validations = append(*validations, NoPrice)
	}
}
