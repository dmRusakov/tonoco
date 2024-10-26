package crypt

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}

	return string(hash), nil
}

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
