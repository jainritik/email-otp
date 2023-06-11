package utils_test

import (
	"testing"

	"github.com/jainritik/email-otp/utils"
)

func TestValidateEmail(t *testing.T) {
	validEmails := []string{
		"tester1@dso.org.sg",
		"admin.hr@dso.org.sg",
		"Hr@dso.org.sg",
		"jon@dso.org.sg",
		"Jon.it@dso.org.sg",
	}

	invalidEmails := []string{
		"tester1@proc.dso.org.sg",
		"tester2@yahoo.com",
		"tester2@a.dso.org.sg.yahoo.com",
		"tester2@dso.org.sg.yahoo.com",
		"123@dso.org",
		"123@it.dso.org",
		"123@dso.com",
		"123@12@dso.com",
	}

	for _, email := range validEmails {
		if !utils.ValidateEmail(email) {
			t.Errorf("Expected email '%s' to be valid, but it was invalid", email)
		}
	}

	for _, email := range invalidEmails {
		if utils.ValidateEmail(email) {
			t.Errorf("Expected email '%s' to be invalid, but it was valid", email)
		}
	}
}

func TestIsEmailDomainAllowed(t *testing.T) {
	allowedEmails := []string{
		"tester1@dso.org.sg",
		"admin.hr@dso.org.sg",
		"Hr@dso.org.sg",
		"jon@dso.org.sg",
		"Jon.it@dso.org.sg",
	}

	disallowedEmails := []string{
		"tester1@yahoo.com",
		"admin.hr@gmail.com",
		"Hr@abc.org.sg",
		"jon@xyz.org.sg",
		"Jon.it@xyz.org.sg",
	}

	for _, email := range allowedEmails {
		if !utils.IsEmailDomainAllowed(email) {
			t.Errorf("Expected email '%s' to have allowed domain, but it was disallowed", email)
		}
	}

	for _, email := range disallowedEmails {
		if utils.IsEmailDomainAllowed(email) {
			t.Errorf("Expected email '%s' to have disallowed domain, but it was allowed", email)
		}
	}
}
