package validator

import (
	"regexp"

	"github.com/gmcriptobox/otus-go-final-project/internal/entity/request"
)

var (
	ipRegex      = regexp.MustCompile(`^([01]?\d\d?|2[0-4]\d|25[0-5])(?:\.(?:[01]?\d\d?|2[0-4]\d|25[0-5])){3}$`)
	networkRegex = regexp.MustCompile(
		`^([01]?\d\d?|2[0-4]\d|25[0-5])(?:\.(?:[01]?\d\d?|2[0-4]\d|25[0-5])){3}(?:/[0-2]\d|/3[0-2])?$`,
	)
)

func ValidateAuthRequest(request *request.AuthRequest) bool {
	if request == nil {
		return false
	}

	if request.Login == "" || request.Password == "" || request.IP == "" {
		return false
	}

	return ipRegex.MatchString(request.IP)
}

func ValidateBucketResetRequest(request *request.BucketResetRequest) bool {
	if request == nil {
		return false
	}

	if request.Login == "" && request.IP == "" {
		return false
	}

	return request.IP == "" || ipRegex.MatchString(request.IP)
}

func ValidateNetworkRequest(request *request.NetworkRequest) bool {
	return ValidateNetwork(request.Network)
}

func ValidateNetwork(network string) bool {
	return networkRegex.MatchString(network)
}
