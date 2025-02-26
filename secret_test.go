package utils

import (
	"strings"
	"testing"
)

type secretTestCase struct {
	prefix string
}

var secretTestCases = []secretTestCase{
	{"test"},
	{""},
}

func TestNewSecret(test *testing.T) {
	for _, testCase := range secretTestCases {
		secret, secretError := NewSecret(testCase.prefix)
		if secretError != nil {
			test.Fatalf("failed to create secret: %v", secretError)
		}

		if !strings.HasPrefix(secret, testCase.prefix) {
			test.Errorf("expected secret to start with %v, got %v", testCase.prefix, secret)
		}
	}
}
