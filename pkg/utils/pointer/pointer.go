package pointer

func Pointer[T any](v T) *T {
	return &v
}

func UintPtr(v uint) *uint {
	return &v
}

func Uint32Ptr(v uint32) *uint32 {
	return &v
}

func Uint64Ptr(v uint64) *uint64 {
	return &v
}
