// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/proto/metadata_repository (interfaces: ManagementClient,ManagementServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	metadata_repository "github.com/kakao/varlog/proto/metadata_repository"
	grpc "google.golang.org/grpc"
)

// MockManagementClient is a mock of ManagementClient interface.
type MockManagementClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagementClientMockRecorder
}

// MockManagementClientMockRecorder is the mock recorder for MockManagementClient.
type MockManagementClientMockRecorder struct {
	mock *MockManagementClient
}

// NewMockManagementClient creates a new mock instance.
func NewMockManagementClient(ctrl *gomock.Controller) *MockManagementClient {
	mock := &MockManagementClient{ctrl: ctrl}
	mock.recorder = &MockManagementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagementClient) EXPECT() *MockManagementClientMockRecorder {
	return m.recorder
}

// AddPeer mocks base method.
func (m *MockManagementClient) AddPeer(arg0 context.Context, arg1 *metadata_repository.AddPeerRequest, arg2 ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPeer", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPeer indicates an expected call of AddPeer.
func (mr *MockManagementClientMockRecorder) AddPeer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockManagementClient)(nil).AddPeer), varargs...)
}

// GetClusterInfo mocks base method.
func (m *MockManagementClient) GetClusterInfo(arg0 context.Context, arg1 *metadata_repository.GetClusterInfoRequest, arg2 ...grpc.CallOption) (*metadata_repository.GetClusterInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClusterInfo", varargs...)
	ret0, _ := ret[0].(*metadata_repository.GetClusterInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInfo indicates an expected call of GetClusterInfo.
func (mr *MockManagementClientMockRecorder) GetClusterInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInfo", reflect.TypeOf((*MockManagementClient)(nil).GetClusterInfo), varargs...)
}

// RemovePeer mocks base method.
func (m *MockManagementClient) RemovePeer(arg0 context.Context, arg1 *metadata_repository.RemovePeerRequest, arg2 ...grpc.CallOption) (*types.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemovePeer", varargs...)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemovePeer indicates an expected call of RemovePeer.
func (mr *MockManagementClientMockRecorder) RemovePeer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePeer", reflect.TypeOf((*MockManagementClient)(nil).RemovePeer), varargs...)
}

// MockManagementServer is a mock of ManagementServer interface.
type MockManagementServer struct {
	ctrl     *gomock.Controller
	recorder *MockManagementServerMockRecorder
}

// MockManagementServerMockRecorder is the mock recorder for MockManagementServer.
type MockManagementServerMockRecorder struct {
	mock *MockManagementServer
}

// NewMockManagementServer creates a new mock instance.
func NewMockManagementServer(ctrl *gomock.Controller) *MockManagementServer {
	mock := &MockManagementServer{ctrl: ctrl}
	mock.recorder = &MockManagementServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagementServer) EXPECT() *MockManagementServerMockRecorder {
	return m.recorder
}

// AddPeer mocks base method.
func (m *MockManagementServer) AddPeer(arg0 context.Context, arg1 *metadata_repository.AddPeerRequest) (*types.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPeer", arg0, arg1)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPeer indicates an expected call of AddPeer.
func (mr *MockManagementServerMockRecorder) AddPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockManagementServer)(nil).AddPeer), arg0, arg1)
}

// GetClusterInfo mocks base method.
func (m *MockManagementServer) GetClusterInfo(arg0 context.Context, arg1 *metadata_repository.GetClusterInfoRequest) (*metadata_repository.GetClusterInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterInfo", arg0, arg1)
	ret0, _ := ret[0].(*metadata_repository.GetClusterInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterInfo indicates an expected call of GetClusterInfo.
func (mr *MockManagementServerMockRecorder) GetClusterInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterInfo", reflect.TypeOf((*MockManagementServer)(nil).GetClusterInfo), arg0, arg1)
}

// RemovePeer mocks base method.
func (m *MockManagementServer) RemovePeer(arg0 context.Context, arg1 *metadata_repository.RemovePeerRequest) (*types.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePeer", arg0, arg1)
	ret0, _ := ret[0].(*types.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemovePeer indicates an expected call of RemovePeer.
func (mr *MockManagementServerMockRecorder) RemovePeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePeer", reflect.TypeOf((*MockManagementServer)(nil).RemovePeer), arg0, arg1)
}
