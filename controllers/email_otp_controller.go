package controllers

import (
	"fmt"
	"github.com/jainritik/email-otp/models"
	"github.com/jainritik/email-otp/services"
	"github.com/jainritik/email-otp/utils"
)

type EmailOTPController struct {
	emailService services.EmailService
	otpGenerator utils.OTPGenerator
	otpStorage   models.OTPStorage
}

func NewEmailOTPController() *EmailOTPController {
	return &EmailOTPController{
		emailService: services.NewEmailService(),
		otpGenerator: utils.NewOTPGenerator(),
		otpStorage:   models.NewOTPStorage(),
	}
}

func (controller *EmailOTPController) GenerateOTP(userEmail string) utils.StatusCode {
	// Validate email address
	if !utils.ValidateEmail(userEmail) {
		return utils.STATUS_EMAIL_INVALID
	}

	// Check if email domain is allowed
	if !utils.IsEmailDomainAllowed(userEmail) {
		return utils.STATUS_EMAIL_INVALID
	}

	// Generate OTP
	otp := controller.otpGenerator.GenerateOTP()

	// Store OTP with user's email
	controller.otpStorage.StoreOTP(userEmail, otp)

	// Send OTP email
	emailBody := fmt.Sprintf("Your OTP code is %s. The code is valid for 1 minute.", otp)
	err := controller.emailService.SendEmail(userEmail, emailBody)
	if err != nil {
		return utils.STATUS_EMAIL_FAIL
	}

	return utils.STATUS_EMAIL_OK
}

func (controller *EmailOTPController) CheckOTP(userEmail, enteredOTP string) utils.StatusCode {
	// Validate email address
	if !utils.ValidateEmail(userEmail) {
		return utils.STATUS_OTP_FAIL
	}

	// Check if email domain is allowed
	if !utils.IsEmailDomainAllowed(userEmail) {
		return utils.STATUS_OTP_FAIL
	}

	// Retrieve OTP for the user's email
	storedOTP := controller.otpStorage.GetOTP(userEmail)

	// Compare entered OTP with stored OTP
	if storedOTP == enteredOTP {
		// Clear stored OTP after successful verification
		controller.otpStorage.ClearOTP(userEmail)
		return utils.STATUS_OTP_OK
	}

	return utils.STATUS_OTP_FAIL
}
