// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package gomocktest is a generated GoMock package.
package gomocktest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFoo is a mock of Foo interface.
type MockFoo struct {
	ctrl     *gomock.Controller
	recorder *MockFooMockRecorder
}

// MockFooMockRecorder is the mock recorder for MockFoo.
type MockFooMockRecorder struct {
	mock *MockFoo
}

// NewMockFoo creates a new mock instance.
func NewMockFoo(ctrl *gomock.Controller) *MockFoo {
	mock := &MockFoo{ctrl: ctrl}
	mock.recorder = &MockFooMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFoo) EXPECT() *MockFooMockRecorder {
	return m.recorder
}

// A mocks base method.
func (m *MockFoo) A(arg0 int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "A", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// A indicates an expected call of A.
func (mr *MockFooMockRecorder) A(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "A", reflect.TypeOf((*MockFoo)(nil).A), arg0)
}
