// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/pkg/varlog (interfaces: LogStreamAppender)

// Package varlog is a generated GoMock package.
package varlog

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogStreamAppender is a mock of LogStreamAppender interface.
type MockLogStreamAppender struct {
	ctrl     *gomock.Controller
	recorder *MockLogStreamAppenderMockRecorder
}

// MockLogStreamAppenderMockRecorder is the mock recorder for MockLogStreamAppender.
type MockLogStreamAppenderMockRecorder struct {
	mock *MockLogStreamAppender
}

// NewMockLogStreamAppender creates a new mock instance.
func NewMockLogStreamAppender(ctrl *gomock.Controller) *MockLogStreamAppender {
	mock := &MockLogStreamAppender{ctrl: ctrl}
	mock.recorder = &MockLogStreamAppenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogStreamAppender) EXPECT() *MockLogStreamAppenderMockRecorder {
	return m.recorder
}

// AppendBatch mocks base method.
func (m *MockLogStreamAppender) AppendBatch(arg0 [][]byte, arg1 BatchCallback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendBatch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendBatch indicates an expected call of AppendBatch.
func (mr *MockLogStreamAppenderMockRecorder) AppendBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendBatch", reflect.TypeOf((*MockLogStreamAppender)(nil).AppendBatch), arg0, arg1)
}

// Close mocks base method.
func (m *MockLogStreamAppender) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockLogStreamAppenderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLogStreamAppender)(nil).Close))
}
