package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yourusername/email-otp/utils"
)

type EmailOTP struct {
	GeneratedOTPs map[string]string
}

func NewEmailOTP() *EmailOTP {
	return &EmailOTP{
		GeneratedOTPs: make(map[string]string),
	}
}

func (e *EmailOTP) GenerateOTPEmail(userEmail string, emailService *utils.EmailService) string {
	if !emailService.IsValidEmail(userEmail) {
		return "STATUS_EMAIL_INVALID"
	}

	otp := e.generateOTP()
	e.GeneratedOTPs[userEmail] = otp

	emailBody := fmt.Sprintf("Your OTP Code is %s. The code is valid for 1 minute.", otp)
	emailService.SendEmail(userEmail, emailBody)

	return "STATUS_EMAIL_OK"
}

func (e *EmailOTP) CheckOTP(userEmail, userOTP string) string {
	if otp, ok := e.GeneratedOTPs[userEmail]; ok {
		if userOTP == otp {
			delete(e.GeneratedOTPs, userEmail)
			return "STATUS_OTP_OK"
		}
	}

	return "STATUS_OTP_FAIL"
}

func (e *EmailOTP) generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"

	otp := make([]byte, 6)
	for i := 0; i < 6; i++ {
		otp[i] = digits[rand.Intn(len(digits))]
	}

	return string(otp)
}
