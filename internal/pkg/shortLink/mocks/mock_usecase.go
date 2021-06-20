// Code generated by MockGen. DO NOT EDIT.
// Source: linkShortener/internal/pkg/shortLink (interfaces: ShortLinkUsecase)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockShortLinkUsecase is a mock of ShortLinkUsecase interface.
type MockShortLinkUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockShortLinkUsecaseMockRecorder
}

// MockShortLinkUsecaseMockRecorder is the mock recorder for MockShortLinkUsecase.
type MockShortLinkUsecaseMockRecorder struct {
	mock *MockShortLinkUsecase
}

// NewMockShortLinkUsecase creates a new mock instance.
func NewMockShortLinkUsecase(ctrl *gomock.Controller) *MockShortLinkUsecase {
	mock := &MockShortLinkUsecase{ctrl: ctrl}
	mock.recorder = &MockShortLinkUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortLinkUsecase) EXPECT() *MockShortLinkUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockShortLinkUsecase) Create(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockShortLinkUsecaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShortLinkUsecase)(nil).Create), arg0)
}

// Get mocks base method.
func (m *MockShortLinkUsecase) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockShortLinkUsecaseMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShortLinkUsecase)(nil).Get), arg0)
}
