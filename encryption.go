package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

type Encryption struct {
	encoder string
	secret  string
}

// NewEncryption creates a new Encryption instance with the given encoder and secret.
func NewEncryption(encoder string, secret string) (*Encryption, error) {
	if len(encoder) != 32 {
		return nil, fmt.Errorf("encoder must be 32 characters long, got %d", len(encoder))
	}

	if len(secret) != 16 && len(secret) != 24 && len(secret) != 32 {
		return nil, fmt.Errorf("secret must be 16, 24, or 32 bytes long, got %d", len(secret))
	}

	encryption := &Encryption{
		encoder,
		secret,
	}

	return encryption, nil
}

// Encode encodes the given input string using the encoder.
func (encryption *Encryption) Encode(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	encoded := base32.NewEncoding(encryption.encoder).
		WithPadding(base32.NoPadding).
		EncodeToString([]byte(input))

	return encoded, nil
}

// Decode decodes the given input string using the encoder.
func (encryption *Encryption) Decode(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	decoded, decodeError := base32.NewEncoding(encryption.encoder).
		WithPadding(base32.NoPadding).
		DecodeString(input)

	if decodeError != nil {
		return "", fmt.Errorf("failed to decode input string: %w", decodeError)
	}

	return string(decoded), nil
}

// Encrypt encrypts the given input string using the secret.
func (encryption *Encryption) Encrypt(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	block, cipherError := aes.NewCipher([]byte(encryption.secret))
	if cipherError != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", cipherError)
	}

	gcm, gcmError := cipher.NewGCM(block)
	if gcmError != nil {
		return "", fmt.Errorf("failed to create GCM block cipher: %w", gcmError)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, readError := io.ReadFull(rand.Reader, nonce)
	if readError != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", readError)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(input), nil)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given input string using the secret.
func (encryption *Encryption) Decrypt(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}

	ciphertext, decodeError := base64.URLEncoding.DecodeString(input)
	if decodeError != nil {
		return "", fmt.Errorf("failed to decode base64 input: %w", decodeError)
	}

	block, cipherError := aes.NewCipher([]byte(encryption.secret))
	if cipherError != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", cipherError)
	}

	gcm, gcmError := cipher.NewGCM(block)
	if gcmError != nil {
		return "", fmt.Errorf("failed to create GCM block cipher: %w", gcmError)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, gcmOpenError := gcm.Open(nil, nonce, ciphertext, nil)
	if gcmOpenError != nil {
		return "", fmt.Errorf("failed to decrypt ciphertext: %w", gcmOpenError)
	}

	return string(plaintext), nil
}
