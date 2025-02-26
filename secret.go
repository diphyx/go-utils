package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// NewSecret generates a new secret string with the given prefix.
func NewSecret(prefix string) (string, error) {
	seed := make([]byte, 64)

	_, readError := rand.Read(seed)
	if readError != nil {
		return "", fmt.Errorf("failed to read random seed: %w", readError)
	}

	hash := sha512.New()

	_, writeError := hash.Write(seed)
	if writeError != nil {
		return "", fmt.Errorf("failed to write seed to hash: %w", writeError)
	}

	secret := prefix + hex.EncodeToString(hash.Sum(nil))

	return secret, nil
}
