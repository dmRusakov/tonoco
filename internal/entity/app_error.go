package entity

import (
	"github.com/dmRusakov/tonoco/pkg/common/errors"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrFilterIsNil = errors.New("filter is nil")
)
