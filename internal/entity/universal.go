package entity

import (
	"github.com/dmRusakov/tonoco/pkg/utils/pointer"
)

type DataPagination struct {
	OrderBy  *string
	OrderDir *string
	Page     *uint64
	PerPage  *uint64
}

func CheckDataPagination(d *DataPagination) {
	// OrderBy
	if d.OrderBy == nil {
		d.OrderBy = pointer.StringToPtr("SortOrder")
	}

	// OrderDir
	if d.OrderDir == nil {
		d.OrderDir = pointer.StringToPtr("ASC")
	}

	// PerPage
	if d.PerPage == nil {
		if d.Page == nil {
			d.PerPage = pointer.UintTo64Ptr(999999999999999999)
		} else {
			d.PerPage = pointer.UintTo64Ptr(18)
		}
	}

	// Page
	if d.Page == nil {
		d.Page = pointer.UintTo64Ptr(1)
	}
}

type DataConfig struct {
	Count *uint64

	IsIdsOnly      *bool
	IsCount        *bool
	IsUpdateFilter *bool
	IsKeepIdsOrder *bool
}
