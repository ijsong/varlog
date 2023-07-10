// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/internal/admin/snmanager (interfaces: StorageNodeManager)

// Package snmanager is a generated GoMock package.
package snmanager

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/kakao/varlog/pkg/types"
	admpb "github.com/kakao/varlog/proto/admpb"
	snpb "github.com/kakao/varlog/proto/snpb"
	varlogpb "github.com/kakao/varlog/proto/varlogpb"
)

// MockStorageNodeManager is a mock of StorageNodeManager interface.
type MockStorageNodeManager struct {
	ctrl     *gomock.Controller
	recorder *MockStorageNodeManagerMockRecorder
}

// MockStorageNodeManagerMockRecorder is the mock recorder for MockStorageNodeManager.
type MockStorageNodeManagerMockRecorder struct {
	mock *MockStorageNodeManager
}

// NewMockStorageNodeManager creates a new mock instance.
func NewMockStorageNodeManager(ctrl *gomock.Controller) *MockStorageNodeManager {
	mock := &MockStorageNodeManager{ctrl: ctrl}
	mock.recorder = &MockStorageNodeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageNodeManager) EXPECT() *MockStorageNodeManagerMockRecorder {
	return m.recorder
}

// AddLogStream mocks base method.
func (m *MockStorageNodeManager) AddLogStream(arg0 context.Context, arg1 *varlogpb.LogStreamDescriptor) (*varlogpb.LogStreamDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLogStream", arg0, arg1)
	ret0, _ := ret[0].(*varlogpb.LogStreamDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLogStream indicates an expected call of AddLogStream.
func (mr *MockStorageNodeManagerMockRecorder) AddLogStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLogStream", reflect.TypeOf((*MockStorageNodeManager)(nil).AddLogStream), arg0, arg1)
}

// AddLogStreamReplica mocks base method.
func (m *MockStorageNodeManager) AddLogStreamReplica(arg0 context.Context, arg1 types.StorageNodeID, arg2 types.TopicID, arg3 types.LogStreamID, arg4 string) (snpb.LogStreamReplicaMetadataDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLogStreamReplica", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(snpb.LogStreamReplicaMetadataDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLogStreamReplica indicates an expected call of AddLogStreamReplica.
func (mr *MockStorageNodeManagerMockRecorder) AddLogStreamReplica(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLogStreamReplica", reflect.TypeOf((*MockStorageNodeManager)(nil).AddLogStreamReplica), arg0, arg1, arg2, arg3, arg4)
}

// AddStorageNode mocks base method.
func (m *MockStorageNodeManager) AddStorageNode(arg0 context.Context, arg1 types.StorageNodeID, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddStorageNode", arg0, arg1, arg2)
}

// AddStorageNode indicates an expected call of AddStorageNode.
func (mr *MockStorageNodeManagerMockRecorder) AddStorageNode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStorageNode", reflect.TypeOf((*MockStorageNodeManager)(nil).AddStorageNode), arg0, arg1, arg2)
}

// Close mocks base method.
func (m *MockStorageNodeManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageNodeManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorageNodeManager)(nil).Close))
}

// Contains mocks base method.
func (m *MockStorageNodeManager) Contains(arg0 types.StorageNodeID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Contains indicates an expected call of Contains.
func (mr *MockStorageNodeManagerMockRecorder) Contains(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockStorageNodeManager)(nil).Contains), arg0)
}

// ContainsAddress mocks base method.
func (m *MockStorageNodeManager) ContainsAddress(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsAddress", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContainsAddress indicates an expected call of ContainsAddress.
func (mr *MockStorageNodeManagerMockRecorder) ContainsAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsAddress", reflect.TypeOf((*MockStorageNodeManager)(nil).ContainsAddress), arg0)
}

// GetMetadata mocks base method.
func (m *MockStorageNodeManager) GetMetadata(arg0 context.Context, arg1 types.StorageNodeID) (*snpb.StorageNodeMetadataDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", arg0, arg1)
	ret0, _ := ret[0].(*snpb.StorageNodeMetadataDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockStorageNodeManagerMockRecorder) GetMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockStorageNodeManager)(nil).GetMetadata), arg0, arg1)
}

// GetMetadataByAddress mocks base method.
func (m *MockStorageNodeManager) GetMetadataByAddress(arg0 context.Context, arg1 types.StorageNodeID, arg2 string) (*snpb.StorageNodeMetadataDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadataByAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(*snpb.StorageNodeMetadataDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadataByAddress indicates an expected call of GetMetadataByAddress.
func (mr *MockStorageNodeManagerMockRecorder) GetMetadataByAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataByAddress", reflect.TypeOf((*MockStorageNodeManager)(nil).GetMetadataByAddress), arg0, arg1, arg2)
}

// RemoveLogStreamReplica mocks base method.
func (m *MockStorageNodeManager) RemoveLogStreamReplica(arg0 context.Context, arg1 types.StorageNodeID, arg2 types.TopicID, arg3 types.LogStreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLogStreamReplica", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLogStreamReplica indicates an expected call of RemoveLogStreamReplica.
func (mr *MockStorageNodeManagerMockRecorder) RemoveLogStreamReplica(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLogStreamReplica", reflect.TypeOf((*MockStorageNodeManager)(nil).RemoveLogStreamReplica), arg0, arg1, arg2, arg3)
}

// RemoveStorageNode mocks base method.
func (m *MockStorageNodeManager) RemoveStorageNode(arg0 types.StorageNodeID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveStorageNode", arg0)
}

// RemoveStorageNode indicates an expected call of RemoveStorageNode.
func (mr *MockStorageNodeManagerMockRecorder) RemoveStorageNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStorageNode", reflect.TypeOf((*MockStorageNodeManager)(nil).RemoveStorageNode), arg0)
}

// Seal mocks base method.
func (m *MockStorageNodeManager) Seal(arg0 context.Context, arg1 types.TopicID, arg2 types.LogStreamID, arg3 types.GLSN) ([]snpb.LogStreamReplicaMetadataDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]snpb.LogStreamReplicaMetadataDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seal indicates an expected call of Seal.
func (mr *MockStorageNodeManagerMockRecorder) Seal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockStorageNodeManager)(nil).Seal), arg0, arg1, arg2, arg3)
}

// Sync mocks base method.
func (m *MockStorageNodeManager) Sync(arg0 context.Context, arg1 types.TopicID, arg2 types.LogStreamID, arg3, arg4 types.StorageNodeID, arg5 types.GLSN) (*snpb.SyncStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*snpb.SyncStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockStorageNodeManagerMockRecorder) Sync(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockStorageNodeManager)(nil).Sync), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Trim mocks base method.
func (m *MockStorageNodeManager) Trim(arg0 context.Context, arg1 types.TopicID, arg2 types.GLSN) ([]admpb.TrimResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trim", arg0, arg1, arg2)
	ret0, _ := ret[0].([]admpb.TrimResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Trim indicates an expected call of Trim.
func (mr *MockStorageNodeManagerMockRecorder) Trim(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trim", reflect.TypeOf((*MockStorageNodeManager)(nil).Trim), arg0, arg1, arg2)
}

// Unseal mocks base method.
func (m *MockStorageNodeManager) Unseal(arg0 context.Context, arg1 types.TopicID, arg2 types.LogStreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unseal", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unseal indicates an expected call of Unseal.
func (mr *MockStorageNodeManagerMockRecorder) Unseal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unseal", reflect.TypeOf((*MockStorageNodeManager)(nil).Unseal), arg0, arg1, arg2)
}
