// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage_node.go

// Package storage is a generated GoMock package.
package storage

import (
	gomock "github.com/golang/mock/gomock"
	types "github.com/kakao/varlog/pkg/varlog/types"
	storage_node "github.com/kakao/varlog/proto/storage_node"
	varlog "github.com/kakao/varlog/proto/varlog"
	reflect "reflect"
)

// MockManagement is a mock of Management interface
type MockManagement struct {
	ctrl     *gomock.Controller
	recorder *MockManagementMockRecorder
}

// MockManagementMockRecorder is the mock recorder for MockManagement
type MockManagementMockRecorder struct {
	mock *MockManagement
}

// NewMockManagement creates a new mock instance
func NewMockManagement(ctrl *gomock.Controller) *MockManagement {
	mock := &MockManagement{ctrl: ctrl}
	mock.recorder = &MockManagementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagement) EXPECT() *MockManagementMockRecorder {
	return m.recorder
}

// GetMetadata mocks base method
func (m *MockManagement) GetMetadata(clusterID types.ClusterID, metadataType storage_node.MetadataType) (*varlog.StorageNodeMetadataDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", clusterID, metadataType)
	ret0, _ := ret[0].(*varlog.StorageNodeMetadataDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata
func (mr *MockManagementMockRecorder) GetMetadata(clusterID, metadataType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockManagement)(nil).GetMetadata), clusterID, metadataType)
}

// AddLogStream mocks base method
func (m *MockManagement) AddLogStream(clusterID types.ClusterID, storageNodeID types.StorageNodeID, logStreamID types.LogStreamID, path string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLogStream", clusterID, storageNodeID, logStreamID, path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLogStream indicates an expected call of AddLogStream
func (mr *MockManagementMockRecorder) AddLogStream(clusterID, storageNodeID, logStreamID, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLogStream", reflect.TypeOf((*MockManagement)(nil).AddLogStream), clusterID, storageNodeID, logStreamID, path)
}

// RemoveLogStream mocks base method
func (m *MockManagement) RemoveLogStream(clusterID types.ClusterID, storageNodeID types.StorageNodeID, logStreamID types.LogStreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLogStream", clusterID, storageNodeID, logStreamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLogStream indicates an expected call of RemoveLogStream
func (mr *MockManagementMockRecorder) RemoveLogStream(clusterID, storageNodeID, logStreamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLogStream", reflect.TypeOf((*MockManagement)(nil).RemoveLogStream), clusterID, storageNodeID, logStreamID)
}

// Seal mocks base method
func (m *MockManagement) Seal(clusterID types.ClusterID, storageNodeID types.StorageNodeID, logStreamID types.LogStreamID, lastCommittedGLSN types.GLSN) (varlog.LogStreamStatus, types.GLSN, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", clusterID, storageNodeID, logStreamID, lastCommittedGLSN)
	ret0, _ := ret[0].(varlog.LogStreamStatus)
	ret1, _ := ret[1].(types.GLSN)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Seal indicates an expected call of Seal
func (mr *MockManagementMockRecorder) Seal(clusterID, storageNodeID, logStreamID, lastCommittedGLSN interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockManagement)(nil).Seal), clusterID, storageNodeID, logStreamID, lastCommittedGLSN)
}

// Unseal mocks base method
func (m *MockManagement) Unseal(clusterID types.ClusterID, storageNodeID types.StorageNodeID, logStreamID types.LogStreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unseal", clusterID, storageNodeID, logStreamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unseal indicates an expected call of Unseal
func (mr *MockManagementMockRecorder) Unseal(clusterID, storageNodeID, logStreamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unseal", reflect.TypeOf((*MockManagement)(nil).Unseal), clusterID, storageNodeID, logStreamID)
}

// MockLogStreamExecutorGetter is a mock of LogStreamExecutorGetter interface
type MockLogStreamExecutorGetter struct {
	ctrl     *gomock.Controller
	recorder *MockLogStreamExecutorGetterMockRecorder
}

// MockLogStreamExecutorGetterMockRecorder is the mock recorder for MockLogStreamExecutorGetter
type MockLogStreamExecutorGetterMockRecorder struct {
	mock *MockLogStreamExecutorGetter
}

// NewMockLogStreamExecutorGetter creates a new mock instance
func NewMockLogStreamExecutorGetter(ctrl *gomock.Controller) *MockLogStreamExecutorGetter {
	mock := &MockLogStreamExecutorGetter{ctrl: ctrl}
	mock.recorder = &MockLogStreamExecutorGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogStreamExecutorGetter) EXPECT() *MockLogStreamExecutorGetterMockRecorder {
	return m.recorder
}

// GetLogStreamExecutor mocks base method
func (m *MockLogStreamExecutorGetter) GetLogStreamExecutor(logStreamID types.LogStreamID) (LogStreamExecutor, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogStreamExecutor", logStreamID)
	ret0, _ := ret[0].(LogStreamExecutor)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetLogStreamExecutor indicates an expected call of GetLogStreamExecutor
func (mr *MockLogStreamExecutorGetterMockRecorder) GetLogStreamExecutor(logStreamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogStreamExecutor", reflect.TypeOf((*MockLogStreamExecutorGetter)(nil).GetLogStreamExecutor), logStreamID)
}

// GetLogStreamExecutors mocks base method
func (m *MockLogStreamExecutorGetter) GetLogStreamExecutors() []LogStreamExecutor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogStreamExecutors")
	ret0, _ := ret[0].([]LogStreamExecutor)
	return ret0
}

// GetLogStreamExecutors indicates an expected call of GetLogStreamExecutors
func (mr *MockLogStreamExecutorGetterMockRecorder) GetLogStreamExecutors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogStreamExecutors", reflect.TypeOf((*MockLogStreamExecutorGetter)(nil).GetLogStreamExecutors))
}
