// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/pkg/mrc (interfaces: MetadataRepositoryClient)

// Package mrc is a generated GoMock package.
package mrc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/kakao/varlog/pkg/types"
	varlogpb "github.com/kakao/varlog/proto/varlogpb"
)

// MockMetadataRepositoryClient is a mock of MetadataRepositoryClient interface.
type MockMetadataRepositoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataRepositoryClientMockRecorder
}

// MockMetadataRepositoryClientMockRecorder is the mock recorder for MockMetadataRepositoryClient.
type MockMetadataRepositoryClientMockRecorder struct {
	mock *MockMetadataRepositoryClient
}

// NewMockMetadataRepositoryClient creates a new mock instance.
func NewMockMetadataRepositoryClient(ctrl *gomock.Controller) *MockMetadataRepositoryClient {
	mock := &MockMetadataRepositoryClient{ctrl: ctrl}
	mock.recorder = &MockMetadataRepositoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataRepositoryClient) EXPECT() *MockMetadataRepositoryClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMetadataRepositoryClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockMetadataRepositoryClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).Close))
}

// GetMetadata mocks base method.
func (m *MockMetadataRepositoryClient) GetMetadata(arg0 context.Context) (*varlogpb.MetadataDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", arg0)
	ret0, _ := ret[0].(*varlogpb.MetadataDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockMetadataRepositoryClientMockRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).GetMetadata), arg0)
}

// RegisterLogStream mocks base method.
func (m *MockMetadataRepositoryClient) RegisterLogStream(arg0 context.Context, arg1 *varlogpb.LogStreamDescriptor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterLogStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterLogStream indicates an expected call of RegisterLogStream.
func (mr *MockMetadataRepositoryClientMockRecorder) RegisterLogStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLogStream", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).RegisterLogStream), arg0, arg1)
}

// RegisterStorageNode mocks base method.
func (m *MockMetadataRepositoryClient) RegisterStorageNode(arg0 context.Context, arg1 *varlogpb.StorageNodeDescriptor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterStorageNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterStorageNode indicates an expected call of RegisterStorageNode.
func (mr *MockMetadataRepositoryClientMockRecorder) RegisterStorageNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterStorageNode", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).RegisterStorageNode), arg0, arg1)
}

// RegisterTopic mocks base method.
func (m *MockMetadataRepositoryClient) RegisterTopic(arg0 context.Context, arg1 types.TopicID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTopic", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterTopic indicates an expected call of RegisterTopic.
func (mr *MockMetadataRepositoryClientMockRecorder) RegisterTopic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTopic", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).RegisterTopic), arg0, arg1)
}

// Seal mocks base method.
func (m *MockMetadataRepositoryClient) Seal(arg0 context.Context, arg1 types.LogStreamID) (types.GLSN, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", arg0, arg1)
	ret0, _ := ret[0].(types.GLSN)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seal indicates an expected call of Seal.
func (mr *MockMetadataRepositoryClientMockRecorder) Seal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).Seal), arg0, arg1)
}

// UnregisterLogStream mocks base method.
func (m *MockMetadataRepositoryClient) UnregisterLogStream(arg0 context.Context, arg1 types.LogStreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterLogStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterLogStream indicates an expected call of UnregisterLogStream.
func (mr *MockMetadataRepositoryClientMockRecorder) UnregisterLogStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterLogStream", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).UnregisterLogStream), arg0, arg1)
}

// UnregisterStorageNode mocks base method.
func (m *MockMetadataRepositoryClient) UnregisterStorageNode(arg0 context.Context, arg1 types.StorageNodeID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterStorageNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterStorageNode indicates an expected call of UnregisterStorageNode.
func (mr *MockMetadataRepositoryClientMockRecorder) UnregisterStorageNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterStorageNode", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).UnregisterStorageNode), arg0, arg1)
}

// UnregisterTopic mocks base method.
func (m *MockMetadataRepositoryClient) UnregisterTopic(arg0 context.Context, arg1 types.TopicID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnregisterTopic", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnregisterTopic indicates an expected call of UnregisterTopic.
func (mr *MockMetadataRepositoryClientMockRecorder) UnregisterTopic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterTopic", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).UnregisterTopic), arg0, arg1)
}

// Unseal mocks base method.
func (m *MockMetadataRepositoryClient) Unseal(arg0 context.Context, arg1 types.LogStreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unseal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unseal indicates an expected call of Unseal.
func (mr *MockMetadataRepositoryClientMockRecorder) Unseal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unseal", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).Unseal), arg0, arg1)
}

// UpdateLogStream mocks base method.
func (m *MockMetadataRepositoryClient) UpdateLogStream(arg0 context.Context, arg1 *varlogpb.LogStreamDescriptor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLogStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLogStream indicates an expected call of UpdateLogStream.
func (mr *MockMetadataRepositoryClientMockRecorder) UpdateLogStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLogStream", reflect.TypeOf((*MockMetadataRepositoryClient)(nil).UpdateLogStream), arg0, arg1)
}
