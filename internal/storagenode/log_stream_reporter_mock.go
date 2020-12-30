// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/internal/storagenode (interfaces: LogStreamReporter)

// Package storagenode is a generated GoMock package.
package storagenode

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/kakao/varlog/pkg/types"
)

// MockLogStreamReporter is a mock of LogStreamReporter interface
type MockLogStreamReporter struct {
	ctrl     *gomock.Controller
	recorder *MockLogStreamReporterMockRecorder
}

// MockLogStreamReporterMockRecorder is the mock recorder for MockLogStreamReporter
type MockLogStreamReporterMockRecorder struct {
	mock *MockLogStreamReporter
}

// NewMockLogStreamReporter creates a new mock instance
func NewMockLogStreamReporter(ctrl *gomock.Controller) *MockLogStreamReporter {
	mock := &MockLogStreamReporter{ctrl: ctrl}
	mock.recorder = &MockLogStreamReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogStreamReporter) EXPECT() *MockLogStreamReporterMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockLogStreamReporter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockLogStreamReporterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLogStreamReporter)(nil).Close))
}

// Commit mocks base method
func (m *MockLogStreamReporter) Commit(arg0 context.Context, arg1 []CommittedLogStreamStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockLogStreamReporterMockRecorder) Commit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockLogStreamReporter)(nil).Commit), arg0, arg1)
}

// GetReport mocks base method
func (m *MockLogStreamReporter) GetReport(arg0 context.Context) (map[types.LogStreamID]UncommittedLogStreamStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReport", arg0)
	ret0, _ := ret[0].(map[types.LogStreamID]UncommittedLogStreamStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReport indicates an expected call of GetReport
func (mr *MockLogStreamReporterMockRecorder) GetReport(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReport", reflect.TypeOf((*MockLogStreamReporter)(nil).GetReport), arg0)
}

// Run mocks base method
func (m *MockLogStreamReporter) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockLogStreamReporterMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockLogStreamReporter)(nil).Run), arg0)
}

// StorageNodeID mocks base method
func (m *MockLogStreamReporter) StorageNodeID() types.StorageNodeID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageNodeID")
	ret0, _ := ret[0].(types.StorageNodeID)
	return ret0
}

// StorageNodeID indicates an expected call of StorageNodeID
func (mr *MockLogStreamReporterMockRecorder) StorageNodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageNodeID", reflect.TypeOf((*MockLogStreamReporter)(nil).StorageNodeID))
}
