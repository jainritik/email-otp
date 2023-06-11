// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jainritik/email-otp/controllers (interfaces: EmailOTPControllerInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	utils "github.com/jainritik/email-otp/utils"
)

// MockEmailOTPControllerInterface is a mock of EmailOTPControllerInterface interface.
type MockEmailOTPControllerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEmailOTPControllerInterfaceMockRecorder
}

// MockEmailOTPControllerInterfaceMockRecorder is the mock recorder for MockEmailOTPControllerInterface.
type MockEmailOTPControllerInterfaceMockRecorder struct {
	mock *MockEmailOTPControllerInterface
}

// NewMockEmailOTPControllerInterface creates a new mock instance.
func NewMockEmailOTPControllerInterface(ctrl *gomock.Controller) *MockEmailOTPControllerInterface {
	mock := &MockEmailOTPControllerInterface{ctrl: ctrl}
	mock.recorder = &MockEmailOTPControllerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailOTPControllerInterface) EXPECT() *MockEmailOTPControllerInterfaceMockRecorder {
	return m.recorder
}

// CheckOTP mocks base method.
func (m *MockEmailOTPControllerInterface) CheckOTP(arg0, arg1 string) utils.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOTP", arg0, arg1)
	ret0, _ := ret[0].(utils.StatusCode)
	return ret0
}

// CheckOTP indicates an expected call of CheckOTP.
func (mr *MockEmailOTPControllerInterfaceMockRecorder) CheckOTP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOTP", reflect.TypeOf((*MockEmailOTPControllerInterface)(nil).CheckOTP), arg0, arg1)
}

// GenerateOTP mocks base method.
func (m *MockEmailOTPControllerInterface) GenerateOTP(arg0 string) utils.StatusCode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateOTP", arg0)
	ret0, _ := ret[0].(utils.StatusCode)
	return ret0
}

// GenerateOTP indicates an expected call of GenerateOTP.
func (mr *MockEmailOTPControllerInterfaceMockRecorder) GenerateOTP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateOTP", reflect.TypeOf((*MockEmailOTPControllerInterface)(nil).GenerateOTP), arg0)
}