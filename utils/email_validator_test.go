package utils_test

import (
	"github.com/jainritik/email-otp/utils"
	"strings"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	// Valid email address
	validEmail := "john.doe@dso.org.sg"
	validResult := utils.ValidateEmail(validEmail)
	if !validResult {
		t.Errorf("Expected ValidateEmail('%s') to return true, but got false", validEmail)
	}

	// Email address exceeding 100 characters
	invalidEmail := "a" + strings.Repeat("b", 99) + "@example.org.sg"
	invalidResult := utils.ValidateEmail(invalidEmail)
	if invalidResult {
		t.Errorf("Expected ValidateEmail('%s') to return false, but got true", invalidEmail)
	}

	// Invalid email address format
	invalidFormatEmail := "john.doe@example.org"
	invalidFormatResult := utils.ValidateEmail(invalidFormatEmail)
	if invalidFormatResult {
		t.Errorf("Expected ValidateEmail('%s') to return false, but got true", invalidFormatEmail)
	}

	// Email address without 'dso.org.sg' domain
	invalidDomainEmail := "john.doe@example.com"
	invalidDomainResult := utils.ValidateEmail(invalidDomainEmail)
	if invalidDomainResult {
		t.Errorf("Expected ValidateEmail('%s') to return false, but got true", invalidDomainEmail)
	}

	// Email address with disallowed characters
	disallowedCharsEmail := "joh[n]@example.org.sg"
	disallowedCharsResult := utils.ValidateEmail(disallowedCharsEmail)
	if disallowedCharsResult {
		t.Errorf("Expected ValidateEmail('%s') to return false, but got true", disallowedCharsEmail)
	}
}
