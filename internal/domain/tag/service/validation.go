package service

import (
	"github.com/dmRusakov/tonoco/internal/entity"
	"github.com/dmRusakov/tonoco/pkg/common/validation"
)

var (
	NoProduct   = validation.NewValidation("Product is required", "Product is required", "j2assq", "ProductId")
	NoTagType   = validation.NewValidation("Tag Type is required", "Type is required", "jhassq", "TagTypeId")
	NoTagSelect = validation.NewValidation("Tag Select is required", "Select is required", "jhafsq", "TagSelectId")
)

func (s *Service) Validate(item *Item, fields *map[string]interface{}) []entity.Error {
	var validations []entity.Error

	// validate Item
	if item != nil {
		results := make(chan []entity.Error)
		countTests := 0

		// ProductId
		go func() {
			var validations []entity.Error
			s.validateProduct(item.ProductId, &validations)
			results <- validations
		}()
		countTests++

		// TagTypeId
		go func() {
			var validations []entity.Error
			s.validateType(item.TagTypeId, &validations)
			results <- validations
		}()
		countTests++

		// TagSelectId
		go func() {
			var validations []entity.Error
			s.validateSelect(item.TagSelectId, &validations)
			results <- validations
		}()
		countTests++

		// Collect results from the channel
		for i := 0; i < countTests; i++ {
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
				case "ProductId":
					s.validateProduct(value.(string), &v)
					break
				case "TagTypeId":
					s.validateType(value.(string), &v)
					break
				case "TagSelectId":
					s.validateSelect(value.(string), &v)
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

// ProductId
func (s *Service) validateProduct(p string, validations *[]entity.Error) {
	if p == "" {
		*validations = append(*validations, NoProduct)
	}
}

// TagTypeId
func (s *Service) validateType(t string, validations *[]entity.Error) {
	if t == "" {
		*validations = append(*validations, NoTagType)
	}
}

// TagSelectId
func (s *Service) validateSelect(t string, validations *[]entity.Error) {
	if t == "" {
		*validations = append(*validations, NoTagSelect)
	}
}
