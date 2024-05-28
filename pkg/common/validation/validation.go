package validation

import (
	"github.com/dmRusakov/tonoco/internal/entity"
)

func NewValidation(msg string, adminMsg string, code string, field string) entity.Error {
	return entity.Error{
		Message:    msg,
		DevMessage: adminMsg,
		Code:       code,
		Field:      field,
	}
}
