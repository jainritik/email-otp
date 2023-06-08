package models

import (
	"fmt"
	"math/rand"
	"time"
)

type EmailOTP struct {
	otpMap map[string]string
}

func NewEmailOTP() *EmailOTP {
	return &EmailOTP{
		otpMap: make(map[string]string),
	}
}

func (eo *EmailOTP) GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	return otp
}

func (eo *EmailOTP) ValidateOTP(email, otp string) bool {
	storedOTP, ok := eo.otpMap[email]
	if !ok {
		return false
	}
	delete(eo.otpMap, email)
	return storedOTP == otp
}
