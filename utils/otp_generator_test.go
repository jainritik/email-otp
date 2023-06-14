package utils_test

import (
	"strconv"
	"testing"

	"github.com/jainritik/email-otp/utils"
)

func TestGenerateOTP(t *testing.T) {
	otpGenerator := utils.NewOTPGenerator()

	// Test generating multiple OTPs and ensure they are unique
	otp1, _ := otpGenerator.GenerateOTP()
	otp2, _ := otpGenerator.GenerateOTP()
	if otp1 == otp2 {
		t.Errorf("Generated OTPs are not unique: %s, %s", otp1, otp2)
	}

	// Test generating OTPs in a loop and verify their properties
	for i := 0; i < 100; i++ {
		otp, err := otpGenerator.GenerateOTP()
		if err != nil {
			t.Errorf("Unexpected error while generating OTP: %v", err)
		}

		// Assert OTP length is 6
		if len(otp) != 6 {
			t.Errorf("Expected OTP length 6, got %d", len(otp))
		}

		// Assert OTP is a valid 6-digit number
		if _, err := strconv.Atoi(otp); err != nil {
			t.Errorf("OTP is not a valid number: %s", otp)
		}

		// Assert OTP is not equal to "000000"
		if otp == "000000" {
			t.Errorf("OTP should not be '000000'")
		}

		// Assert OTP consists of digits only
		for _, digit := range otp {
			if digit < '0' || digit > '9' {
				t.Errorf("OTP contains invalid characters: %s", otp)
				break
			}
		}
	}
}
