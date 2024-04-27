package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoName            = validation.NewValidation("Name is required", "Name is required", "98adfd", "Name")
	NoUrl             = validation.NewValidation("Url is required", "Url is required", "00iad2", "Url")
	NoSpecificationID = validation.NewValidation("SpecificationID is required", "Specification ID is required", "98unfd", "SpecificationID")
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

		go func() {
			var validations []entity.Validation
			s.validateSpecificationIDt(item.SpecificationID, &validations)
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
				case "SpecificationID":
					s.validateSpecificationIDt(value.(string), &v)
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

func (s *Service) validateSpecificationIDt(t string, validations *[]entity.Validation) {
	if t == "" {
		*validations = append(*validations, NoSpecificationID)
	}
}
