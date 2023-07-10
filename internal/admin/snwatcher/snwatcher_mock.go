// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/internal/admin/snwatcher (interfaces: EventHandler)

// Package snwatcher is a generated GoMock package.
package snwatcher

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/kakao/varlog/pkg/types"
	snpb "github.com/kakao/varlog/proto/snpb"
)

// MockEventHandler is a mock of EventHandler interface.
type MockEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerMockRecorder
}

// MockEventHandlerMockRecorder is the mock recorder for MockEventHandler.
type MockEventHandlerMockRecorder struct {
	mock *MockEventHandler
}

// NewMockEventHandler creates a new mock instance.
func NewMockEventHandler(ctrl *gomock.Controller) *MockEventHandler {
	mock := &MockEventHandler{ctrl: ctrl}
	mock.recorder = &MockEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandler) EXPECT() *MockEventHandlerMockRecorder {
	return m.recorder
}

// HandleHeartbeatTimeout mocks base method.
func (m *MockEventHandler) HandleHeartbeatTimeout(arg0 context.Context, arg1 types.StorageNodeID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleHeartbeatTimeout", arg0, arg1)
}

// HandleHeartbeatTimeout indicates an expected call of HandleHeartbeatTimeout.
func (mr *MockEventHandlerMockRecorder) HandleHeartbeatTimeout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleHeartbeatTimeout", reflect.TypeOf((*MockEventHandler)(nil).HandleHeartbeatTimeout), arg0, arg1)
}

// HandleReport mocks base method.
func (m *MockEventHandler) HandleReport(arg0 context.Context, arg1 *snpb.StorageNodeMetadataDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleReport", arg0, arg1)
}

// HandleReport indicates an expected call of HandleReport.
func (mr *MockEventHandlerMockRecorder) HandleReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleReport", reflect.TypeOf((*MockEventHandler)(nil).HandleReport), arg0, arg1)
}
