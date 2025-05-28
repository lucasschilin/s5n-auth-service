package validator

import (
	"net"
	"regexp"
	"strings"
)

func IsValidEmailAddress(address string) bool {
	addressRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !addressRegex.MatchString(address) {
		return false
	}

	parts := strings.Split(address, "@")
	username := parts[0]
	domain := parts[1]

	rfc5321FullAddressMaxLength := 254
	if len(address) > rfc5321FullAddressMaxLength {
		return false
	}
	rfc5321UsernameAddressMaxLength := 64
	if len(username) > rfc5321UsernameAddressMaxLength {
		return false
	}

	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false
	}

	return true

}
