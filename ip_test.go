package utils

import (
	"testing"
)

type ipTestCase struct {
	ip     string
	number int64
}

var (
	validIpTestCases = []ipTestCase{
		{"192.168.0.1", 3232235521},
		{"0.0.0.0", 0},
		{"255.255.255.255", 4294967295},
	}
	invalidIpTestCases = []ipTestCase{
		{"invalid_ip", -1},
		{"::1", 4294967296},
	}
)

func TestIpToNumber(test *testing.T) {
	for _, testCase := range validIpTestCases {
		output, err := IpToNumber(testCase.ip)
		if err != nil {
			test.Fatalf("failed to convert ip to number: %v", err)
		}

		if output != testCase.number {
			test.Errorf("expected %d, got %d", testCase.number, output)
		}
	}

	for _, testCase := range invalidIpTestCases {
		_, err := IpToNumber(testCase.ip)
		if err == nil {
			test.Errorf("expected error for invalid ip, got nil")
		}
	}
}

func TestNumberToIp(test *testing.T) {
	for _, testCase := range validIpTestCases {
		output, err := NumberToIp(testCase.number)
		if err != nil {
			test.Fatalf("failed to convert number to ip: %v", err)
		}

		if output != testCase.ip {
			test.Errorf("expected %s, got %s", testCase.ip, output)
		}
	}

	for _, testCase := range invalidIpTestCases {
		_, err := NumberToIp(testCase.number)
		if err == nil {
			test.Fatalf("expected error for invalid number, got nil")
		}
	}
}
