package models

import (
	"errors"
	"time"
)

type OTPStorage struct {
	otpMap map[string]otpData
}

type otpData struct {
	otp       string
	timestamp time.Time
}

func NewOTPStorage() OTPStorage {
	return OTPStorage{
		otpMap: make(map[string]otpData),
	}
}

func (storage *OTPStorage) StoreOTP(userEmail, otp string, timestamp time.Time) error {
	storage.otpMap[userEmail] = otpData{
		otp:       otp,
		timestamp: timestamp,
	}
	return nil
}

func (storage *OTPStorage) GetOTP(userEmail string) (string, time.Time, bool, error) {
	data, ok := storage.otpMap[userEmail]
	if ok {
		return data.otp, data.timestamp, true, nil
	}
	return "", time.Time{}, false, errors.New("OTP not found")
}

func (storage *OTPStorage) ClearOTP(userEmail string) error {
	if _, ok := storage.otpMap[userEmail]; ok {
		delete(storage.otpMap, userEmail)
		return nil
	}
	return errors.New("OTP not found")
}
