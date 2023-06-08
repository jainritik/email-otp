package controllers

import (
	"github.com/jainritik/email-otp/models"
	"github.com/jainritik/email-otp/services"
	"github.com/jainritik/email-otp/utils"
)

const (
	STATUS_EMAIL_OK      = 1
	STATUS_EMAIL_FAIL    = 2
	STATUS_EMAIL_INVALID = 3
	STATUS_OTP_OK        = 1
	STATUS_OTP_FAIL      = 2
	STATUS_OTP_TIMEOUT   = 3
)

type EmailOTPController struct {
	emailOTPModel  *models.EmailOTP
	emailService   *services.EmailService
	emailValidator *utils.EmailValidator
}

func NewEmailOTPController() *EmailOTPController {
	return &EmailOTPController{
		emailOTPModel:  models.NewEmailOTP(),
		emailService:   services.NewEmailService(),
		emailValidator: utils.NewEmailValidator(),
	}
}

func (ec *EmailOTPController) GenerateOTP(email string) int {
	if !ec.emailValidator.ValidateEmail(email) {
		return STATUS_EMAIL_INVALID
	}

	otp := ec.emailOTPModel.GenerateOTP()

	err := ec.emailService.SendEmail(email, otp)
	if err != nil {
		return STATUS_EMAIL_FAIL
	}

	return STATUS_EMAIL_OK
}

func (ec *EmailOTPController) ValidateOTP(email, otp string) int {
	if !ec.emailValidator.ValidateEmail(email) {
		return STATUS_EMAIL_INVALID
	}

	valid := ec.emailOTPModel.ValidateOTP(email, otp)
	if !valid {
		return STATUS_OTP_FAIL
	}

	return STATUS_OTP_OK
}
