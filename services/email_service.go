package services

import (
	"fmt"
)

type EmailService struct {
	// Email sending service implementation
}

func NewEmailService() *EmailService {
	return &EmailService{}
}

func (es *EmailService) SendEmail(email, otp string) error {
	// Implement the email sending functionality here
	fmt.Printf("Sending email to %s with OTP: %s\n", email, otp)
	return nil
}
