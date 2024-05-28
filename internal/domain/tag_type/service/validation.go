package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoName      = validation.NewValidation("Name is required", "Name is required", "ls22wa", "Name")
	NoUrl       = validation.NewValidation("Url is required", "Url is required", "00gfgf", "Url")
	NoType      = validation.NewValidation("Type is required", "Type is required", "00ghgf", "Type")
	InvalidType = validation.NewValidation("Type is invalid", "Type is invalid", "02ghgf", "Type")
)

func (s *Service) Validate(item *Item, fields *map[string]interface{}) []entity.Error {
	var errors []entity.Error

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
			s.validateType(item.Type, &validations)
			results <- validations
		}()

		// Collect results from the channel
		for i := 0; i < 3; i++ {
			v := <-results
			errors = append(errors, v...)
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
				}

				results <- v
			}(field, value)
		}

		// Collect results from the channel
		for range *fields {
			v := <-results
			errors = append(errors, v...)
		}
	}

	return errors
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
		*validations = append(*validations, NoType)
	} else {
		switch t {
		case "select":
			break
		case "text":
			break
		default:
			*validations = append(*validations, InvalidType)
		}
	}
}
