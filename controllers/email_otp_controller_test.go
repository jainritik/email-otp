package controllers

import (
	"errors"
	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/jainritik/email-otp/mocks"
	"github.com/jainritik/email-otp/utils"
	"testing"
	"time"
)

func TestEmailOTPController_GenerateOTP_ValidEmail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "test@dso.org.sg"
	otp := "123456"

	mockOTPGenerator.EXPECT().GenerateOTP().Return(otp, nil).Times(1)
	mockOTPStorage.EXPECT().StoreOTP(userEmail, gomock.Any(), gomock.Any()).Return(nil).Times(1)
	mockEmailService.EXPECT().SendEmail(userEmail, gomock.Any()).Return(nil).Times(1)

	status := controller.GenerateOTP(userEmail)
	if status != utils.STATUS_EMAIL_OK {
		t.Errorf("Expected STATUS_EMAIL_OK, got: %v", status)
	}
}

func TestEmailOTPController_GenerateOTP_InvalidEmail_InvalidStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "invalid_email"

	status := controller.GenerateOTP(userEmail)
	if status != utils.STATUS_EMAIL_INVALID {
		t.Errorf("Expected STATUS_EMAIL_INVALID, got: %v", status)
	}
}

func TestEmailOTPController_GenerateOTP_StorageError_FailStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "test@dso.org.sg"
	otp := "123456"
	storeError := errors.New("storage error")

	mockOTPGenerator.EXPECT().GenerateOTP().Return(otp, nil).Times(1)
	mockOTPStorage.EXPECT().StoreOTP(userEmail, gomock.Any(), gomock.Any()).Return(storeError).Times(1)

	status := controller.GenerateOTP(userEmail)
	if status != utils.STATUS_EMAIL_FAIL {
		t.Errorf("Expected STATUS_EMAIL_FAIL, got: %v", status)
	}
}

func TestEmailOTPController_CheckOTP_ValidEmail_ValidOTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "test@dso.org.sg"
	otp := "123456"
	storedOTP := "123456"
	createTime := time.Now()
	exists := true

	mockOTPStorage.EXPECT().GetOTP(userEmail).Return(storedOTP, createTime, exists, nil).Times(1)
	mockOTPStorage.EXPECT().ClearOTP(userEmail).Return(nil).Times(1)

	status := controller.CheckOTP(userEmail, otp)
	if status != utils.STATUS_OTP_OK {
		t.Errorf("Expected STATUS_OTP_OK, got: %v", status)
	}
}

func TestEmailOTPController_CheckOTP_InvalidEmail_InvalidStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "invalid_email"
	otp := "123456"

	status := controller.CheckOTP(userEmail, otp)
	if status != utils.STATUS_OTP_FAIL {
		t.Errorf("Expected STATUS_OTP_FAIL, got: %v", status)
	}
}

func TestEmailOTPController_CheckOTP_OTPDoesNotExist_FailStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "test@dso.org.sg"
	otp := "123456"
	exists := false

	mockOTPStorage.EXPECT().GetOTP(userEmail).Return("", time.Time{}, exists, nil).Times(1)

	status := controller.CheckOTP(userEmail, otp)
	if status != utils.STATUS_OTP_FAIL {
		t.Errorf("Expected STATUS_OTP_FAIL, got: %v", status)
	}
}

func TestEmailOTPController_CheckOTP_InvalidOTP_FailStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "test@dso.org.sg"
	otp := "123456"
	storedOTP := "654321"
	createTime := time.Now()
	exists := true

	mockOTPStorage.EXPECT().GetOTP(userEmail).Return(storedOTP, createTime, exists, nil).Times(1)

	status := controller.CheckOTP(userEmail, otp)
	if status != utils.STATUS_OTP_FAIL {
		t.Errorf("Expected STATUS_OTP_FAIL, got: %v", status)
	}
}

func TestEmailOTPController_CheckOTP_ExpiredOTP_FailStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "test@dso.org.sg"
	otp := "123456"
	storedOTP := "123456"
	createTime := time.Now().Add(-time.Minute) // Set create time to be 1 minute ago
	exists := true

	mockOTPStorage.EXPECT().GetOTP(userEmail).Return(storedOTP, createTime, exists, nil).Times(1)
	mockOTPStorage.EXPECT().ClearOTP(userEmail).Return(nil).Times(1)

	status := controller.CheckOTP(userEmail, otp)
	if status != utils.STATUS_OTP_TIMEOUT {
		t.Errorf("Expected STATUS_OTP_TIMEOUT, got: %v", status)
	}
}

func TestEmailOTPController_CheckOTP_StorageError_FailStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmailService := mocks.NewMockEmailServiceInterface(ctrl)
	mockOTPGenerator := mocks.NewMockOTPGeneratorInterface(ctrl)
	mockOTPStorage := mocks.NewMockOTPStorageInterface(ctrl)

	controller := NewEmailOTPController(mockEmailService, mockOTPGenerator, mockOTPStorage)

	userEmail := "test@dso.org.sg"
	otp := "123456"
	storedOTP := "123456"
	createTime := time.Now()
	exists := true
	clearError := errors.New("clear error")

	mockOTPStorage.EXPECT().GetOTP(userEmail).Return(storedOTP, createTime, exists, nil).Times(1)
	mockOTPStorage.EXPECT().ClearOTP(userEmail).Return(clearError).Times(1)

	status := controller.CheckOTP(userEmail, otp)
	if status != utils.STATUS_OTP_FAIL {
		t.Errorf("Expected STATUS_OTP_FAIL, got: %v", status)
	}
}
