// Code generated by MockGen. DO NOT EDIT.
// Source: linkShortener/internal/pkg/shortLink (interfaces: ShortLinkRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockShortLinkRepository is a mock of ShortLinkRepository interface.
type MockShortLinkRepository struct {
	ctrl     *gomock.Controller
	recorder *MockShortLinkRepositoryMockRecorder
}

// MockShortLinkRepositoryMockRecorder is the mock recorder for MockShortLinkRepository.
type MockShortLinkRepositoryMockRecorder struct {
	mock *MockShortLinkRepository
}

// NewMockShortLinkRepository creates a new mock instance.
func NewMockShortLinkRepository(ctrl *gomock.Controller) *MockShortLinkRepository {
	mock := &MockShortLinkRepository{ctrl: ctrl}
	mock.recorder = &MockShortLinkRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortLinkRepository) EXPECT() *MockShortLinkRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockShortLinkRepository) Create(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockShortLinkRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShortLinkRepository)(nil).Create), arg0, arg1)
}

// Exists mocks base method.
func (m *MockShortLinkRepository) Exists(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockShortLinkRepositoryMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockShortLinkRepository)(nil).Exists), arg0)
}

// Get mocks base method.
func (m *MockShortLinkRepository) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockShortLinkRepositoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShortLinkRepository)(nil).Get), arg0)
}