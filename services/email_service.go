package services

import (
	"github.com/jainritik/email-otp/models"
	"github.com/jainritik/email-otp/utils"
)

type EmailOTPService struct {
	emailOTPModel *models.EmailOTP
	emailService  *utils.EmailService
}

func NewEmailOTPService() *EmailOTPService {
	return &EmailOTPService{
		emailOTPModel: models.NewEmailOTP(),
		emailService:  utils.NewEmailService(),
	}
}

func (s *EmailOTPService) GenerateOTPEmail(userEmail string) string {
	return s.emailOTPModel.GenerateOTPEmail(userEmail, s.emailService)
}

func (s *EmailOTPService) CheckOTP(userEmail, userOTP string) string {
	return s.emailOTPModel.CheckOTP(userEmail, userOTP)
}
