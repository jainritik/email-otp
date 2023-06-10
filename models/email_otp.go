package models

type OTPStorage struct {
	otpMap map[string]string
}

func NewOTPStorage() OTPStorage {
	return OTPStorage{
		otpMap: make(map[string]string),
	}
}

func (storage *OTPStorage) StoreOTP(userEmail, otp string) {
	storage.otpMap[userEmail] = otp
}

func (storage *OTPStorage) GetOTP(userEmail string) string {
	return storage.otpMap[userEmail]
}

func (storage *OTPStorage) ClearOTP(userEmail string) {
	delete(storage.otpMap, userEmail)
}
