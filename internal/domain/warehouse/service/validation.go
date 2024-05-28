package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoName         = validation.NewValidation("Name is required", "Name is required", "182341", "Name")
	NoAbbreviation = validation.NewValidation("Abbreviation is required", "Abbreviation is required", "182342", "Abbreviation")
	NoAddressLine1 = validation.NewValidation("Address Line 1 is required", "Address Line 1 is required", "182344", "AddressLine1")
	NoCity         = validation.NewValidation("City is required", "City is required", "182345", "City")
	NoState        = validation.NewValidation("State is required", "State is required", "182346", "State")
	NoZipCode      = validation.NewValidation("Zip Code is required", "Zip Code is required", "182347", "ZipCode")
	NoCountry      = validation.NewValidation("Country is required", "Country is required", "182348", "Country")
	NoEmail        = validation.NewValidation("Email is required", "Email is required", "182349", "Email")
)

func (s *Service) Validate(item *Item, fields *map[string]interface{}) []entity.Error {
	var validations []entity.Error

	// validate Item
	if item != nil {
		results := make(chan []entity.Error)

		// Name
		go func() {
			var validations []entity.Error
			s.validateName(item.Name, &validations)
			results <- validations
		}()

		// Abbreviation
		go func() {
			var validations []entity.Error
			s.validateAbbreviation(item.Abbreviation, &validations)
			results <- validations
		}()

		// AddressLine1
		go func() {
			var validations []entity.Error
			s.validateAddressLine1(item.AddressLine1, &validations)
			results <- validations
		}()

		// City
		go func() {
			var validations []entity.Error
			s.validateCity(item.City, &validations)
			results <- validations
		}()

		// State
		go func() {
			var validations []entity.Error
			s.validateState(item.State, &validations)
			results <- validations
		}()

		// ZipCode
		go func() {
			var validations []entity.Error
			s.validateZipCode(item.ZipCode, &validations)
			results <- validations
		}()

		// Country
		go func() {
			var validations []entity.Error
			s.validateCountry(item.Country, &validations)
			results <- validations
		}()

		// Email
		go func() {
			var validations []entity.Error
			s.validateEmail(item.Email, &validations)
			results <- validations
		}()

		// Collect results from the channel
		for i := 0; i < 8; i++ {
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
				case "Abbreviation":
					s.validateAbbreviation(value.(string), &v)
					break
				case "AddressLine1":
					s.validateAddressLine1(value.(string), &v)
					break
				case "City":
					s.validateCity(value.(string), &v)
					break
				case "State":
					s.validateState(value.(string), &v)
					break
				case "ZipCode":
					s.validateZipCode(value.(string), &v)
					break
				case "Country":
					s.validateCountry(value.(string), &v)
					break
				case "Email":
					s.validateEmail(value.(string), &v)
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

func (s *Service) validateAbbreviation(abbreviation string, validations *[]entity.Error) {
	if abbreviation == "" {
		*validations = append(*validations, NoAbbreviation)
	}
}

func (s *Service) validateAddressLine1(addressLine1 string, validations *[]entity.Error) {
	if addressLine1 == "" {
		*validations = append(*validations, NoAddressLine1)
	}
}

func (s *Service) validateCity(city string, validations *[]entity.Error) {
	if city == "" {
		*validations = append(*validations, NoCity)
	}
}

func (s *Service) validateState(state string, validations *[]entity.Error) {
	if state == "" {
		*validations = append(*validations, NoState)
	}
}

func (s *Service) validateZipCode(zipCode string, validations *[]entity.Error) {
	if zipCode == "" {
		*validations = append(*validations, NoZipCode)
	}
}

func (s *Service) validateCountry(country string, validations *[]entity.Error) {
	if country == "" {
		*validations = append(*validations, NoCountry)
	}
}

func (s *Service) validateEmail(email string, validations *[]entity.Error) {
	if email == "" {
		*validations = append(*validations, NoEmail)
	}
}
