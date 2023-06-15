package services

import (
	"fmt"
)

type EmailService struct{}

// NewEmailService creates a new instance of EmailService
func NewEmailService() EmailService {
	return EmailService{}
}

type EmailServiceInterface interface {
	SendEmail(email string, body string) error
}

// SendEmail sends an email to the given address
func (es EmailService) SendEmail(email, body string) error {
	fmt.Println("Sending email to:", email)
	fmt.Println("Email body:", body)
	return nil
}
