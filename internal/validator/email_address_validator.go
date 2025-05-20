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
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]

	mxRecords, err := net.LookupMX(domain)
	if err != nil || len(mxRecords) == 0 {
		return false
	}

	return true

}
