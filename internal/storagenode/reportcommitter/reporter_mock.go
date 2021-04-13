// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/internal/storagenode/reportcommitter (interfaces: Reporter)

// Package reportcommitter is a generated GoMock package.
package reportcommitter

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/kakao/varlog/pkg/types"
	snpb "github.com/kakao/varlog/proto/snpb"
)

// MockReporter is a mock of Reporter interface.
type MockReporter struct {
	ctrl     *gomock.Controller
	recorder *MockReporterMockRecorder
}

// MockReporterMockRecorder is the mock recorder for MockReporter.
type MockReporterMockRecorder struct {
	mock *MockReporter
}

// NewMockReporter creates a new mock instance.
func NewMockReporter(ctrl *gomock.Controller) *MockReporter {
	mock := &MockReporter{ctrl: ctrl}
	mock.recorder = &MockReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReporter) EXPECT() *MockReporterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockReporter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReporterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReporter)(nil).Close))
}

// Commit mocks base method.
func (m *MockReporter) Commit(arg0 context.Context, arg1 []*snpb.LogStreamCommitResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockReporterMockRecorder) Commit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockReporter)(nil).Commit), arg0, arg1)
}

// GetReport mocks base method.
func (m *MockReporter) GetReport(arg0 context.Context) ([]*snpb.LogStreamUncommitReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReport", arg0)
	ret0, _ := ret[0].([]*snpb.LogStreamUncommitReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReport indicates an expected call of GetReport.
func (mr *MockReporterMockRecorder) GetReport(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReport", reflect.TypeOf((*MockReporter)(nil).GetReport), arg0)
}

// StorageNodeID mocks base method.
func (m *MockReporter) StorageNodeID() types.StorageNodeID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageNodeID")
	ret0, _ := ret[0].(types.StorageNodeID)
	return ret0
}

// StorageNodeID indicates an expected call of StorageNodeID.
func (mr *MockReporterMockRecorder) StorageNodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageNodeID", reflect.TypeOf((*MockReporter)(nil).StorageNodeID))
}
