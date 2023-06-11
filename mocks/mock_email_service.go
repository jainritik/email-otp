// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jainritik/email-otp/services (interfaces: EmailServiceInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmailServiceInterface is a mock of EmailServiceInterface interface.
type MockEmailServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEmailServiceInterfaceMockRecorder
}

// MockEmailServiceInterfaceMockRecorder is the mock recorder for MockEmailServiceInterface.
type MockEmailServiceInterfaceMockRecorder struct {
	mock *MockEmailServiceInterface
}

// NewMockEmailServiceInterface creates a new mock instance.
func NewMockEmailServiceInterface(ctrl *gomock.Controller) *MockEmailServiceInterface {
	mock := &MockEmailServiceInterface{ctrl: ctrl}
	mock.recorder = &MockEmailServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailServiceInterface) EXPECT() *MockEmailServiceInterfaceMockRecorder {
	return m.recorder
}

// SendEmail mocks base method.
func (m *MockEmailServiceInterface) SendEmail(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockEmailServiceInterfaceMockRecorder) SendEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockEmailServiceInterface)(nil).SendEmail), arg0, arg1)
}
