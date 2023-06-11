package utils

import (
	"math/rand"
	"strconv"
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
	otp := rand.Intn(900000) + 100000
	return strconv.Itoa(otp), nil
}

// NormalizeOTP removes leading zeros from the OTP
func NormalizeOTP(otp string) string {
	// Remove leading zeros
	i, err := strconv.Atoi(otp)
	if err != nil {
		return otp
	}
	normalizedOTP := strconv.Itoa(i)
	return normalizedOTP
}
