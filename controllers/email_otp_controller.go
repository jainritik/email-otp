package controllers

import (
	"fmt"
	"github.com/jainritik/email-otp/models"
	"github.com/jainritik/email-otp/services"
	"github.com/jainritik/email-otp/utils"
	"log"
	"time"
)

type EmailOTPController struct {
	emailService services.EmailServiceInterface
	otpGenerator utils.OTPGeneratorInterface
	otpStorage   models.OTPStorageInterface
	tries        map[string]int
	timer        *time.Timer
}

func NewEmailOTPController(emailService services.EmailServiceInterface, otpGenerator utils.OTPGeneratorInterface, otpStorage models.OTPStorageInterface) *EmailOTPController {
	return &EmailOTPController{
		emailService: emailService,
		otpGenerator: otpGenerator,
		otpStorage:   otpStorage,
		tries:        make(map[string]int),
	}
}

type EmailOTPControllerInterface interface {
	GenerateOTP(userEmail string) utils.StatusCode
	CheckOTP(userEmail string, enteredOTP string) utils.StatusCode
}

func (controller *EmailOTPController) GenerateOTP(userEmail string) utils.StatusCode {
	// Validate email address
	if !utils.ValidateEmail(userEmail) {
		return utils.STATUS_EMAIL_INVALID
	}

	// Generate OTP
	otp, err := controller.otpGenerator.GenerateOTP()
	if err != nil {
		log.Printf("Error generating OTP: %v", err)
		return utils.STATUS_EMAIL_FAIL
	}

	// Store OTP with user's email and current time
	currentTime := time.Now()
	err = controller.otpStorage.StoreOTP(userEmail, otp, currentTime)
	if err != nil {
		log.Printf("Error storing OTP: %v", err)
		return utils.STATUS_EMAIL_FAIL
	}

	// Send OTP email
	emailBody := fmt.Sprintf("Your OTP code is %s. The code is valid for 1 minute.", otp)
	err = controller.emailService.SendEmail(userEmail, emailBody)
	if err != nil {
		// Clear stored OTP if email sending fails
		controller.otpStorage.ClearOTP(userEmail)
		log.Printf("Error sending email: %v", err)
		return utils.STATUS_EMAIL_FAIL
	}

	// Reset the tries counter and start the timer
	controller.tries[userEmail] = 0
	controller.startTimer(userEmail)

	return utils.STATUS_EMAIL_OK
}

func (controller *EmailOTPController) CheckOTP(userEmail, enteredOTP string) utils.StatusCode {
	// Validate email address
	if !utils.ValidateEmail(userEmail) {
		return utils.STATUS_OTP_FAIL
	}

	// Increment the tries counter
	controller.tries[userEmail]++

	// Check if maximum tries exceeded
	if controller.tries[userEmail] >= 10 {
		err := controller.otpStorage.ClearOTP(userEmail)
		if err != nil {
			log.Printf("Error clearing OTP after maximum tries: %v", err)
			return utils.STATUS_OTP_FAIL
		}
		controller.stopTimer(userEmail)
		return utils.STATUS_OTP_FAIL
	}

	// Retrieve OTP, its creation time, and existence status for the user's email
	storedOTP, createTime, exists, err := controller.otpStorage.GetOTP(userEmail)
	if err != nil {
		log.Printf("Error retrieving OTP: %v", err)
		return utils.STATUS_OTP_FAIL
	}

	// Check if OTP exists
	if !exists {
		return utils.STATUS_OTP_FAIL
	}

	// Check if OTP has expired
	elapsed := time.Since(createTime)
	if elapsed > time.Minute {
		err = controller.otpStorage.ClearOTP(userEmail)
		if err != nil {
			log.Printf("Error clearing expired OTP: %v", err)
			return utils.STATUS_OTP_FAIL
		}
		controller.stopTimer(userEmail)
		return utils.STATUS_OTP_TIMEOUT
	}

	// Compare entered OTP with stored OTP
	if storedOTP == enteredOTP {
		// Clear stored OTP after successful verification
		err = controller.otpStorage.ClearOTP(userEmail)
		if err != nil {
			log.Printf("Error clearing OTP: %v", err)
			return utils.STATUS_OTP_FAIL
		}
		controller.stopTimer(userEmail)
		return utils.STATUS_OTP_OK
	}

	// Reset the timer for every action taken
	controller.startTimer(userEmail)

	return utils.STATUS_OTP_FAIL
}

func (controller *EmailOTPController) startTimer(userEmail string) {
	// Stop the timer if it's already running
	if controller.timer != nil {
		controller.timer.Stop()
	}

	// Start a new timer for 1 minute duration
	controller.timer = time.AfterFunc(time.Minute, func() {
		err := controller.otpStorage.ClearOTP(userEmail)
		if err != nil {
			log.Printf("Error clearing OTP after timeout: %v", err)
		}
	})

}

func (controller *EmailOTPController) stopTimer(userEmail string) {
	// Stop the timer if it's running
	if controller.timer != nil {
		controller.timer.Stop()
	}

	// Reset the tries counter
	delete(controller.tries, userEmail)
}
