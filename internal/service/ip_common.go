package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gmcriptobox/otus-go-final-project/internal/entity"
)

var ErrNetworkAlreadyExists = errors.New("network already exists in the list")

// GetNetworkPrefixBinary generates the binary network prefix based on the given IP address and subnet mask.
func GetNetworkPrefixBinary(ip, mask string) (string, error) {
	binaryIP, err := IPAddressToBinary(ip)
	if err != nil {
		return "", err
	}

	maskInt, err := strconv.Atoi(mask)
	if err != nil {
		return "", err
	}
	return binaryIP[:maskInt], nil
}

// IPAddressToBinary converts an IP address from decimal format to binary format.
func IPAddressToBinary(ip string) (string, error) {
	ipParts := strings.Split(ip, ".")
	stringBuilder := strings.Builder{}
	for _, part := range ipParts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return "", err
		}
		stringBuilder.WriteString(fmt.Sprintf("%08b", val))
	}
	return stringBuilder.String(), nil
}

// GetNetwork retrieves a network entity based on the input string.
func GetNetwork(network string) (entity.Network, error) {
	parts := strings.Split(network, "/")
	binaryPrefix, err := GetNetworkPrefixBinary(parts[0], parts[1])
	if err != nil {
		return entity.Network{}, err
	}
	return entity.Network{IP: parts[0], Mask: parts[1], BinaryPrefix: binaryPrefix}, nil
}
