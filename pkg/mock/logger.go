// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/logger/logger.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logger "github.com/iyhunko/hash-generation-app/pkg/logger"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockLogger) Error(messages ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range messages {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(messages ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), messages...)
}

// FatalError mocks base method.
func (m *MockLogger) FatalError(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FatalError", err)
}

// FatalError indicates an expected call of FatalError.
func (mr *MockLoggerMockRecorder) FatalError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FatalError", reflect.TypeOf((*MockLogger)(nil).FatalError), err)
}

// Info mocks base method.
func (m *MockLogger) Info(messages ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range messages {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(messages ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), messages...)
}

// Warn mocks base method.
func (m *MockLogger) Warn(messages ...string) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range messages {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerMockRecorder) Warn(messages ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger)(nil).Warn), messages...)
}

// WithStackTrace mocks base method.
func (m *MockLogger) WithStackTrace(directory string) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStackTrace", directory)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// WithStackTrace indicates an expected call of WithStackTrace.
func (mr *MockLoggerMockRecorder) WithStackTrace(directory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStackTrace", reflect.TypeOf((*MockLogger)(nil).WithStackTrace), directory)
}
