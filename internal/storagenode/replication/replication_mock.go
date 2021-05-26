// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/internal/storagenode/replication (interfaces: Replicator,Getter)

// Package replication is a generated GoMock package.
package replication

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/kakao/varlog/pkg/types"
	snpb "github.com/kakao/varlog/proto/snpb"
)

// MockReplicator is a mock of Replicator interface.
type MockReplicator struct {
	ctrl     *gomock.Controller
	recorder *MockReplicatorMockRecorder
}

// MockReplicatorMockRecorder is the mock recorder for MockReplicator.
type MockReplicatorMockRecorder struct {
	mock *MockReplicator
}

// NewMockReplicator creates a new mock instance.
func NewMockReplicator(ctrl *gomock.Controller) *MockReplicator {
	mock := &MockReplicator{ctrl: ctrl}
	mock.recorder = &MockReplicatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicator) EXPECT() *MockReplicatorMockRecorder {
	return m.recorder
}

// Replicate mocks base method.
func (m *MockReplicator) Replicate(arg0 context.Context, arg1 types.LLSN, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replicate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replicate indicates an expected call of Replicate.
func (mr *MockReplicatorMockRecorder) Replicate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replicate", reflect.TypeOf((*MockReplicator)(nil).Replicate), arg0, arg1, arg2)
}

// Sync mocks base method.
func (m *MockReplicator) Sync(arg0 context.Context, arg1 snpb.Replica, arg2 types.GLSN) (*SyncTaskStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0, arg1, arg2)
	ret0, _ := ret[0].(*SyncTaskStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockReplicatorMockRecorder) Sync(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockReplicator)(nil).Sync), arg0, arg1, arg2)
}

// SyncInit mocks base method.
func (m *MockReplicator) SyncInit(arg0 context.Context, arg1, arg2 snpb.SyncPosition) (snpb.SyncPosition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncInit", arg0, arg1, arg2)
	ret0, _ := ret[0].(snpb.SyncPosition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncInit indicates an expected call of SyncInit.
func (mr *MockReplicatorMockRecorder) SyncInit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncInit", reflect.TypeOf((*MockReplicator)(nil).SyncInit), arg0, arg1, arg2)
}

// SyncReplicate mocks base method.
func (m *MockReplicator) SyncReplicate(arg0 context.Context, arg1 snpb.SyncPayload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncReplicate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncReplicate indicates an expected call of SyncReplicate.
func (mr *MockReplicatorMockRecorder) SyncReplicate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncReplicate", reflect.TypeOf((*MockReplicator)(nil).SyncReplicate), arg0, arg1)
}

// MockGetter is a mock of Getter interface.
type MockGetter struct {
	ctrl     *gomock.Controller
	recorder *MockGetterMockRecorder
}

// MockGetterMockRecorder is the mock recorder for MockGetter.
type MockGetterMockRecorder struct {
	mock *MockGetter
}

// NewMockGetter creates a new mock instance.
func NewMockGetter(ctrl *gomock.Controller) *MockGetter {
	mock := &MockGetter{ctrl: ctrl}
	mock.recorder = &MockGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetter) EXPECT() *MockGetterMockRecorder {
	return m.recorder
}

// Replicator mocks base method.
func (m *MockGetter) Replicator(arg0 types.LogStreamID) (Replicator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replicator", arg0)
	ret0, _ := ret[0].(Replicator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Replicator indicates an expected call of Replicator.
func (mr *MockGetterMockRecorder) Replicator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replicator", reflect.TypeOf((*MockGetter)(nil).Replicator), arg0)
}

// Replicators mocks base method.
func (m *MockGetter) Replicators() []Replicator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replicators")
	ret0, _ := ret[0].([]Replicator)
	return ret0
}

// Replicators indicates an expected call of Replicators.
func (mr *MockGetterMockRecorder) Replicators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replicators", reflect.TypeOf((*MockGetter)(nil).Replicators))
}
