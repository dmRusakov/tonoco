package validation

import "github.com/dmRusakov/tonoco/internal/domain/entity"

func NewValidation(msg string, adminMsg string, code string, field string) entity.Validation {
	return entity.Validation{
		Msg:      msg,
		AdminMsg: adminMsg,
		Code:     code,
		Field:    field,
	}
}
