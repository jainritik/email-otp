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

// GenerateOTP generates a 6-digit OTP
func (og OTPGenerator) GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(900000) + 100000
	return strconv.Itoa(otp)
}
