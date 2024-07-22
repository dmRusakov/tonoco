package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

func Uint64Ptr(i uint64) *uint64 {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}

func StringPtr(s string) *string {
	return &s
}