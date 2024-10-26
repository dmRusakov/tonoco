package validation

import (
	"github.com/dmRusakov/tonoco/internal/entity/db"
)

func NewValidation(msg string, adminMsg string, code string, field string) db.Error {
	return db.Error{
		Message:    msg,
		DevMessage: adminMsg,
		Code:       code,
		Field:      field,
	}
}
