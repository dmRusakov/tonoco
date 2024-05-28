package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoName    = validation.NewValidation("Name is required", "Name is required", "jhasqq", "Name")
	NoUrl     = validation.NewValidation("Url is required", "Url is required", "ofyjfd", "Url")
	NoTagType = validation.NewValidation("Tag Type is required", "Type is required", "jhassq", "Type")
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
			s.validateType(item.TagTypeId, &validations)
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
				case "Type":
					s.validateType(value.(string), &v)
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

func (s *Service) validateType(t string, validations *[]entity.Error) {
	if t == "" {
		*validations = append(*validations, NoTagType)
	}
}
