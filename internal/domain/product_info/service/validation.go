package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoName             = validation.NewValidation("Name is required", "Name is required", "94jfaa", "name")
	NoUrl              = validation.NewValidation("Url is required", "Url is required", "kyfqzz", "url")
	NoShortDescription = validation.NewValidation("Short description is required", "Short description is required", "yutyi6", "short_description")
	NoDescription      = validation.NewValidation("Description is required", "Description is required", "0934jr", "description")
	NoStatusID         = validation.NewValidation("Status ID is required", "Status ID is required", "9fj3j3", "status_id")
	NoRegularPrice     = validation.NewValidation("Regular price is required", "Regular price is required", "9fj3j5", "regular_price")
	NoShippingClassID  = validation.NewValidation("Shipping class ID is required", "Shipping class ID is required", "9fj3j7", "shipping_class_id")
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

		go func() {
			var validations []entity.Validation
			s.validateStatusID(item.StatusID, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Validation
			s.validateRegularPrice(item.RegularPrice, &validations)
			results <- validations
		}()

		go func() {
			var validations []entity.Validation
			s.validateShippingClassID(item.ShippingClassID, &validations)
			results <- validations
		}()

		// Collect results from the channel
		for i := 0; i < 7; i++ {
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
				case "StatusID":
					s.validateStatusID(value.(string), &v)
					break
				case "RegularPrice":
					s.validateRegularPrice(value.(float32), &v)
					break
				case "ShippingClassID":
					s.validateShippingClassID(value.(string), &v)
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

func (s *Service) validateStatusID(statusID string, validations *[]entity.Validation) {
	if statusID == "" {
		*validations = append(*validations, NoStatusID)
	}
}

func (s *Service) validateRegularPrice(regularPrice float32, validations *[]entity.Validation) {
	if regularPrice < 0 {
		*validations = append(*validations, NoRegularPrice)
	}
}

func (s *Service) validateShippingClassID(shippingClassID string, validations *[]entity.Validation) {
	if shippingClassID == "" {
		*validations = append(*validations, NoShippingClassID)
	}
}
