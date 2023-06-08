package controllers

import (
	"github.com/jainritik/email-otp/services"
)

type EmailOTPController struct {
	emailOTPService *services.EmailOTPService
}

func NewEmailOTPController() *EmailOTPController {
	return &EmailOTPController{
		emailOTPService: services.NewEmailOTPService(),
	}
}

func (c *EmailOTPController) GenerateOTPEmail(userEmail string) string {
	return c.emailOTPService.GenerateOTPEmail(userEmail)
}

func (c *EmailOTPController) CheckOTP(userEmail, userOTP string) string {
	return c.emailOTPService.CheckOTP(userEmail, userOTP)
}
