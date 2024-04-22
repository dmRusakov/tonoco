package service

import (
	"github.com/dmRusakov/tonoco/internal/domain/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoName             = validation.NewValidation("Name is required", "Name is required", "jdsy3a", "name")
	NoUrl              = validation.NewValidation("Url is required", "Url is required", "pote04", "url")
	NoShortDescription = validation.NewValidation("Short description is required", "Short description is required", "jnci22", "short_description")
	NoDescription      = validation.NewValidation("Description is required", "Description is required", "0a9834", "description")
)

func (s *Service) Validate(item *Item, fields *map[string]interface{}) []entity.Validation {
	var validations []entity.Validation

	// validate Item
	if item != nil {
		results := make(chan []entity.Validation)

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
			s.validateShortDescription(item.ShortDescription, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Validation
			s.validateDescription(item.Description, &validations)
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
				case "Name":
					s.validateName(value.(string), &v)
					break
				case "Url":
					s.validateUrl(value.(string), &v)
					break
				case "ShortDescription":
					s.validateShortDescription(value.(string), &v)
					break
				case "Description":
					s.validateDescription(value.(string), &v)
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

func (s *Service) validateShortDescription(shortDescription string, validations *[]entity.Validation) {
	if shortDescription == "" {
		*validations = append(*validations, NoShortDescription)
	}
}

func (s *Service) validateDescription(description string, validations *[]entity.Validation) {
	if description == "" {
		*validations = append(*validations, NoDescription)
	}
}
