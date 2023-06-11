package utils

import (
	"regexp"
	"strings"
)

type StatusCode string

const (
	STATUS_EMAIL_OK      StatusCode = "EMAIL_OK"
	STATUS_EMAIL_INVALID StatusCode = "EMAIL_INVALID"
	STATUS_EMAIL_FAIL    StatusCode = "EMAIL_FAIL" // Added constant for email sending failure
	STATUS_OTP_OK        StatusCode = "OTP_OK"
	STATUS_OTP_FAIL      StatusCode = "OTP_FAIL"
	STATUS_OTP_TIMEOUT   StatusCode = "OTP_TIMEOUT"
)

// ValidateEmail validates if the given email address is in a valid format
func ValidateEmail(email string) bool {
	// Check if email length exceeds 100 characters
	if len(email) > 100 {
		return false
	}

	// Disallowed special characters in the user part
	disallowedUserChars := `()[\];:@&=+$,/?#%`

	// Regular expression for email validation
	emailRegex := "^[a-zA-Z0-9.!#$%'" + "`" + "*+/=?^_{|}~-]+" +
		"@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*\\.org\\.sg$"

	// Match email against the regex
	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return false
	}

	// Check if subdomain is present
	parts := strings.Split(email, "@")
	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) != 3 || domainParts[len(domainParts)-3] != "dso" || domainParts[len(domainParts)-2] != "org" || domainParts[len(domainParts)-1] != "sg" {
		return false
	}

	// Check if email contains any disallowed characters in the user part
	for _, char := range disallowedUserChars {
		if strings.Contains(parts[0], string(char)) {
			return false
		}
	}

	return true
}

// IsEmailDomainAllowed checks if the domain of the given email address is allowed
func IsEmailDomainAllowed(email string) bool {
	allowedDomain := "@dso.org.sg"
	return strings.HasSuffix(email, allowedDomain)
}
