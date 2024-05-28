package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoName   = validation.NewValidation("Name is required", "Name is required", "0038qq", "name")
	NoUrl    = validation.NewValidation("Url is required", "Url is required", "76eda2", "url")
	NoSymbol = validation.NewValidation("Symbol is required", "Symbol is required", "76eda3", "symbol")
)

func (s *Service) Validate(item *Item, fields *map[string]interface{}) []entity.Error {
	var validations []entity.Error

	// validate Item
	if item != nil {
		results := make(chan []entity.Error)

		// Start goroutines for validateName and validateUrl
		go func() {
			var validations []entity.Error
			s.validateName(item.Name, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Error
			s.validateUrl(item.Url, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Error
			s.validateSymbol(item.Symbol, &validations)
			results <- validations
		}()

		// Collect results from the channel
		for i := 0; i < 3; i++ {
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
				case "Name":
					s.validateName(value.(string), &v)
					break
				case "Url":
					s.validateUrl(value.(string), &v)
					break
				case "Symbol":
					s.validateSymbol(value.(string), &v)
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

func (s *Service) validateName(name string, validations *[]entity.Error) {
	if name == "" {
		*validations = append(*validations, NoName)
	}
}

func (s *Service) validateUrl(url string, validations *[]entity.Error) {
	if url == "" {
		*validations = append(*validations, NoUrl)
	}
}

func (s *Service) validateSymbol(symbol string, validations *[]entity.Error) {
	if symbol == "" {
		*validations = append(*validations, NoSymbol)
	}
}
