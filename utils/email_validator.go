package utils

import (
	"regexp"
)

type EmailValidator struct {
	// Email validation logic
}

func NewEmailValidator() *EmailValidator {
	return &EmailValidator{}
}

func (ev *EmailValidator) ValidateEmail(email string) bool {
	// Implement email validation logic using regex or any other method
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return match
}
