package utils

import (
	"bufio"
	"fmt"
	"os"
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
	// Regular expression for email validation
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`

	// Match email against the regex
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

// IsEmailDomainAllowed checks if the domain of the given email address is allowed
func IsEmailDomainAllowed(email string) bool {
	allowedDomain := "@dso.org.sg"
	return strings.HasSuffix(email, allowedDomain)
}
func ReadOTP() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter OTP: ")
	otp, _ := reader.ReadString('\n')
	otp = strings.TrimSpace(otp)
	return otp
}
