package entity

import (
	"github.com/dmRusakov/tonoco/pkg/common/errors"
)

var (
	ErrCacheNotFound = errors.New("cache not found")
	ErrFilterIsNil   = errors.New("filter is nil")
)
