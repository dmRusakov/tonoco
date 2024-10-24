package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/dmRusakov/tonoco/pkg/common/errors"
	"github.com/google/uuid"
)

func Hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

func HashFilter(filter interface{}) (string, error) {
	// Convert the filter to a JSON string
	jsonFilter, err := json.Marshal(filter)
	if err != nil {
		return "", err
	}

	// Compute the SHA256 hash of the JSON string
	hash := sha256.Sum256(jsonFilter)

	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hash[:]), nil
}

func PtrString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
func IntPtr(i int) *int {
	return &i
}

func UintPtr(i uint) *uint {
	return &i
}

func Uint64Ptr(i uint64) *uint64 {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}

func StringPtr(s string) *string {
	return &s
}

func StringToUUID(s string) (u uuid.UUID, err error) {
	u, err = uuid.Parse(s)
	if err != nil {
		return u, errors.AddCode(err, "837732")
	}

	return u, nil
}

// Remove duplicates from a slice of any
func RemoveDuplicates[T comparable](slice []T, isKeepOrder bool) []T {
	uniqueMap := make(map[T]bool)
	uniqueSlice := []T{}

	for _, item := range slice {
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = true
			if isKeepOrder {
				uniqueSlice = append(uniqueSlice, item)
			}
		}
	}

	if !isKeepOrder {
		for item := range uniqueMap {
			uniqueSlice = append(uniqueSlice, item)
		}
	}

	return uniqueSlice
}
