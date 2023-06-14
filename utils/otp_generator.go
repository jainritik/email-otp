package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type OTPGenerator struct{}

// NewOTPGenerator creates a new instance of OTPGenerator
func NewOTPGenerator() OTPGenerator {
	return OTPGenerator{}
}

type OTPGeneratorInterface interface {
	GenerateOTP() (string, error)
}

// GenerateOTP generates a 6-digit OTP and returns it along with any error
func (og OTPGenerator) GenerateOTP() (string, error) {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number from 0 to 999999 (inclusive)
	otp := rand.Intn(1000000)

	// Format the OTP as a 6-digit string with leading zeros
	otpString := fmt.Sprintf("%06d", otp)

	// Check if the generated OTP is "000000"
	if otpString == "000000" {
		// Generate a new OTP recursively if it is invalid
		return og.GenerateOTP()
	}

	return otpString, nil
}

// NormalizeOTP removes leading zeros from the OTP
//func NormalizeOTP(otp string) string {
//	// Remove leading zeros
//	i, err := strconv.Atoi(otp)
//	if err != nil {
//		return otp
//	}
//	normalizedOTP := strconv.Itoa(i)
//	return normalizedOTP
//}
