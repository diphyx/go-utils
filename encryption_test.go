package utils

import (
	"testing"
)

var (
	encryptionTestCase       = "Sample"
	encryptionValidEncoder   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef"
	encryptionValidSecret    = "0123456789123456"
	encryptionInvalidEncoder = "SHORTENCODER"
	encryptionInvalidSecret  = "shortsecret"
)

func TestNewEncryption(test *testing.T) {
	_, encryptionError := NewEncryption(encryptionValidEncoder, encryptionValidSecret)
	if encryptionError != nil {
		test.Errorf("expected no error, got %v", encryptionError)
	}

	_, encryptionError = NewEncryption(encryptionInvalidEncoder, encryptionValidSecret)
	if encryptionError == nil {
		test.Errorf("expected error for short encoder, got nil")
	}

	_, encryptionError = NewEncryption(encryptionValidEncoder, encryptionInvalidSecret)
	if encryptionError == nil {
		test.Errorf("expected error for short secret, got nil")
	}
}

func TestEncodeDecode(test *testing.T) {
	encryption, encryptionError := NewEncryption(encryptionValidEncoder, encryptionValidSecret)
	if encryptionError != nil {
		test.Fatalf("failed to create encryption: %v", encryptionError)
	}

	encoded, encodeError := encryption.Encode(encryptionTestCase)
	if encodeError != nil {
		test.Fatalf("failed to encode: %v", encodeError)
	}

	decoded, decodeError := encryption.Decode(encoded)
	if decodeError != nil {
		test.Fatalf("failed to decode: %v", decodeError)
	}

	if decoded != encryptionTestCase {
		test.Errorf("expected %s, got %s", encryptionTestCase, decoded)
	}
}

func TestEncryptDecrypt(test *testing.T) {
	encryption, encryptionError := NewEncryption(encryptionValidEncoder, encryptionValidSecret)
	if encryptionError != nil {
		test.Fatalf("failed to create encryption: %v", encryptionError)
	}

	encrypted, encryptError := encryption.Encrypt(encryptionTestCase)
	if encryptError != nil {
		test.Fatalf("failed to encrypt: %v", encryptError)
	}

	decrypted, decryptError := encryption.Decrypt(encrypted)
	if decryptError != nil {
		test.Fatalf("failed to decrypt: %v", decryptError)
	}

	if decrypted != encryptionTestCase {
		test.Errorf("expected %s, got %s", encryptionTestCase, decrypted)
	}
}

func TestEmptyInput(test *testing.T) {
	encryption, encryptionError := NewEncryption(encryptionValidEncoder, encryptionValidSecret)
	if encryptionError != nil {
		test.Fatalf("failed to create encryption: %v", encryptionError)
	}

	_, encodeError := encryption.Encode("")
	if encodeError == nil {
		test.Errorf("expected error for empty input, got nil")
	}

	_, decodeError := encryption.Decode("")
	if decodeError == nil {
		test.Errorf("expected error for empty input, got nil")
	}

	_, encryptError := encryption.Encrypt("")
	if encryptError == nil {
		test.Errorf("expected error for empty input, got nil")
	}

	_, decryptError := encryption.Decrypt("")
	if decryptError == nil {
		test.Errorf("expected error for empty input, got nil")
	}
}
