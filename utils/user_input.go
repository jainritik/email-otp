package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserInput struct {
	reader *bufio.Reader
}

func NewUserInput() *UserInput {
	return &UserInput{
		reader: bufio.NewReader(os.Stdin),
	}
}

// ReadOTP reads the user input for OTP
func (ui *UserInput) ReadOTP() string {
	fmt.Print("Enter OTP: ")
	otp, _ := ui.reader.ReadString('\n')
	otp = strings.TrimSpace(otp)
	return otp
}
