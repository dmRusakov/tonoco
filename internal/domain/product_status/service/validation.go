package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoName = validation.NewValidation("Name is required", "Name is required", "5kd223", "name")
	NoUrl  = validation.NewValidation("Url is required", "Url is required", "98ndaa", "url")
)

func (s *Service) Validate(item *Item, fields *map[string]interface{}) []entity.Validation {
	var validations []entity.Validation

	// validate Item
	if item != nil {
		results := make(chan []entity.Validation)

		// Start goroutines for validateName and validateUrl
		go func() {
			var validations []entity.Validation
			s.validateName(item.Name, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Validation
			s.validateUrl(item.Url, &validations)
			results <- validations
		}()

		// Collect results from the channel
		for i := 0; i < 2; i++ {
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
				case "Name":
					s.validateName(value.(string), &v)
					break
				case "Url":
					s.validateUrl(value.(string), &v)
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

func (s *Service) validateName(name string, validations *[]entity.Validation) {
	if name == "" {
		*validations = append(*validations, NoName)
	}
}

func (s *Service) validateUrl(url string, validations *[]entity.Validation) {
	if url == "" {
		*validations = append(*validations, NoUrl)
	}
}
