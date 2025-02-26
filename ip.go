package utils

import (
	"errors"
	"fmt"
	"math/big"
	"net"
)

var (
	minIpNumber = big.NewInt(0)
	maxIpNumber = big.NewInt(4294967295)
)

// IpToNumber converts an IPv4 address string to its integer representation.
func IpToNumber(input string) (int64, error) {
	parsedIp := net.ParseIP(input)
	if parsedIp == nil {
		return 0, fmt.Errorf("invalid IP address format: %s", input)
	}

	v4Ip := parsedIp.To4()
	if v4Ip == nil {
		return 0, errors.New("invalid IPv4 address")
	}

	output := big.NewInt(0).SetBytes(v4Ip).Int64()

	return output, nil
}

// NumberToIp converts an integer representation of an IPv4 address back to its string form.
func NumberToIp(input int64) (string, error) {
	inputAsBigInt := big.NewInt(input)

	if inputAsBigInt.Cmp(minIpNumber) < 0 || inputAsBigInt.Cmp(maxIpNumber) > 0 {
		return "", fmt.Errorf("input out of range: %d", input)
	}

	buffer := make([]byte, 4)
	bytes := inputAsBigInt.FillBytes(buffer)

	output := net.IP(bytes).String()

	return output, nil
}
