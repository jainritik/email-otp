// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jainritik/email-otp/utils (interfaces: OTPGeneratorInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOTPGeneratorInterface is a mock of OTPGeneratorInterface interface.
type MockOTPGeneratorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOTPGeneratorInterfaceMockRecorder
}

// MockOTPGeneratorInterfaceMockRecorder is the mock recorder for MockOTPGeneratorInterface.
type MockOTPGeneratorInterfaceMockRecorder struct {
	mock *MockOTPGeneratorInterface
}

// NewMockOTPGeneratorInterface creates a new mock instance.
func NewMockOTPGeneratorInterface(ctrl *gomock.Controller) *MockOTPGeneratorInterface {
	mock := &MockOTPGeneratorInterface{ctrl: ctrl}
	mock.recorder = &MockOTPGeneratorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOTPGeneratorInterface) EXPECT() *MockOTPGeneratorInterfaceMockRecorder {
	return m.recorder
}

// GenerateOTP mocks base method.
func (m *MockOTPGeneratorInterface) GenerateOTP() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateOTP")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateOTP indicates an expected call of GenerateOTP.
func (mr *MockOTPGeneratorInterfaceMockRecorder) GenerateOTP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateOTP", reflect.TypeOf((*MockOTPGeneratorInterface)(nil).GenerateOTP))
}
