package main

import (
	"fmt"
	"github.com/jainritik/email-otp/controllers"
)

func main() {
	// Create an instance of the EmailOTPController
	emailOTPController := controllers.NewEmailOTPController()

	// Generate OTP and validate email
	email := "user@example.com"
	status := emailOTPController.GenerateOTP(email)

	if status == controllers.STATUS_EMAIL_OK {
		fmt.Println("OTP generated and sent successfully")
	} else {
		fmt.Println("Failed to generate OTP")
	}

	// Validate OTP
	otp := "123456"
	status = emailOTPController.ValidateOTP(email, otp)

	if status == controllers.STATUS_OTP_OK {
		fmt.Println("OTP validation successful")
	} else if status == controllers.STATUS_OTP_FAIL {
		fmt.Println("OTP validation failed")
	} else if status == controllers.STATUS_OTP_TIMEOUT {
		fmt.Println("OTP validation timeout")
	} else {
		fmt.Println("Unknown OTP validation status")
	}
}
