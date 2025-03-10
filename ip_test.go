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
		{
			ip:     "192.168.0.1",
			number: 3232235521,
		},
		{
			ip:     "0.0.0.0",
			number: 0,
		},
		{
			ip:     "255.255.255.255",
			number: 4294967295,
		},
	}
	invalidIpTestCases = []ipTestCase{
		{
			ip:     "invalid_ip",
			number: -1,
		},
		{
			ip:     "::1",
			number: 4294967296,
		},
	}
)

func TestIpToNumber(test *testing.T) {
	for _, testCase := range validIpTestCases {
		output, convertError := IpToNumber(testCase.ip)
		if convertError != nil {
			test.Fatalf("failed to convert ip to number: %v", convertError)
		}

		if output != testCase.number {
			test.Errorf("expected %d, got %d", testCase.number, output)
		}
	}

	for _, testCase := range invalidIpTestCases {
		_, convertError := IpToNumber(testCase.ip)
		if convertError == nil {
			test.Errorf("expected error for invalid ip, got nil")
		}
	}
}

func TestNumberToIp(test *testing.T) {
	for _, testCase := range validIpTestCases {
		output, convertError := NumberToIp(testCase.number)
		if convertError != nil {
			test.Fatalf("failed to convert number to ip: %v", convertError)
		}

		if output != testCase.ip {
			test.Errorf("expected %s, got %s", testCase.ip, output)
		}
	}

	for _, testCase := range invalidIpTestCases {
		_, convertError := NumberToIp(testCase.number)
		if convertError == nil {
			test.Fatalf("expected error for invalid number, got nil")
		}
	}
}
