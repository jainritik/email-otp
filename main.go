package main

import (
	"fmt"
	"github.com/jainritik/email-otp/controllers"
	"github.com/jainritik/email-otp/utils"
	"os"
	"time"
)

func main() {
	emailOTPController := controllers.NewEmailOTPController()

	// Read user email address
	var userEmail string
	fmt.Println("Please enter your email address:")
	fmt.Scanln(&userEmail)

	// Generate OTP and send email
	status := emailOTPController.GenerateOTP(userEmail)
	switch status {
	case utils.STATUS_EMAIL_OK:
		fmt.Println("OTP has been sent to the email address.")
	case utils.STATUS_EMAIL_INVALID:
		fmt.Println("Email address is invalid.")
		return
	case utils.STATUS_EMAIL_FAIL:
		fmt.Println("Email address does not exist or sending to the email has failed.")
		return
	}

	// Create a channel to receive user input
	inputCh := make(chan string)

	// Create a channel to signal when the maximum tries are exceeded
	maxTriesCh := make(chan struct{})

	// Start a goroutine to read user input
	go func() {
		fmt.Println("Please enter the OTP received in your email:")
		for i := 0; i < 10; i++ {
			var input string
			fmt.Scanln(&input)
			inputCh <- input
		}
		close(maxTriesCh)
	}()

	// Start a timer for user input timeout
	timeoutCh := time.After(1 * time.Minute)

	// Wait for user input, max tries exceeded, or timeout
	for {
		select {
		case enteredOTP := <-inputCh:
			status = emailOTPController.CheckOTP(userEmail, enteredOTP)
			switch status {
			case utils.STATUS_OTP_OK:
				fmt.Println("OTP verification successful.")
				goto exit
			case utils.STATUS_OTP_FAIL:
				fmt.Println("Incorrect OTP. Please try again.")
			case utils.STATUS_OTP_TIMEOUT:
				fmt.Println("OTP verification timeout.")
				goto exit
			}
		case <-maxTriesCh:
			fmt.Println("Maximum tries exceeded.")
			goto exit
		case <-timeoutCh:
			fmt.Println("Timeout for user input.")
			goto exit
		}
	}

exit:
	os.Exit(0)
}
