package storage

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
