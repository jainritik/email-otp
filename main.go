package main

import (
	"fmt"
	"github.com/jainritik/email-otp/controllers"
	"github.com/jainritik/email-otp/utils"
)

func main() {
	emailOTPController := controllers.NewEmailOTPController()

	// Generate OTP and send email
	userEmail := "tester1@dso.org.sg"
	status := emailOTPController.GenerateOTP(userEmail)
	switch status {
	case utils.STATUS_EMAIL_OK:
		fmt.Println("OTP has been sent to the email address.")
	case utils.STATUS_EMAIL_INVALID:
		fmt.Println("Invalid email address.")
	case utils.STATUS_EMAIL_FAIL:
		fmt.Println("Failed to send email. Please try again later.")
	}

	// Check OTP entered by the user
	fmt.Println("Please enter the OTP received in your email:")
	enteredOTP := utils.ReadOTP()
	status = emailOTPController.CheckOTP(userEmail, enteredOTP)
	switch status {
	case utils.STATUS_OTP_OK:
		fmt.Println("OTP verification successful.")
	case utils.STATUS_OTP_FAIL:
		fmt.Println("Incorrect OTP. Please try again.")
	case utils.STATUS_OTP_TIMEOUT:
		fmt.Println("OTP verification timeout.")
	}
}
