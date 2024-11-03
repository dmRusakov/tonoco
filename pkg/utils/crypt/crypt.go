package crypt

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}

	return string(hash), nil
}

func Hash(s string) uuid.UUID {
	hash := sha256.Sum256([]byte(s))
	bytes, _ := uuid.FromBytes(hash[:16])
	return bytes
}

func HashFilter(filter interface{}) uuid.UUID {
	jsonFilter, _ := json.Marshal(filter)
	return Hash(string(jsonFilter))
}
