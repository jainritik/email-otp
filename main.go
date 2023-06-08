package main

import (
	"fmt"

	"github.com/jainritik/email-otp/controllers"
)

func main() {
	emailOTPController := controllers.NewEmailOTPController()

	// Test Case 1: Valid email
	fmt.Println(emailOTPController.GenerateOTPEmail("tester1@dso.org.sg")) // STATUS_EMAIL_OK

	// Test Case 2: Invalid email
	fmt.Println(emailOTPController.GenerateOTPEmail("tester1@proc.dso.org.sg")) // STATUS_EMAIL_INVALID

	// Test Case 3: OTP validation
	userEmail := "tester1@dso.org.sg"
	userOTP := "123456"
	fmt.Println(emailOTPController.CheckOTP(userEmail, userOTP)) // STATUS_OTP_OK

	// Test Case 4: Invalid OTP
	userOTP = "654321"
	fmt.Println(emailOTPController.CheckOTP(userEmail, userOTP)) // STATUS_OTP_FAIL
}
