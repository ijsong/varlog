// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kakao/varlog/internal/admin (interfaces: ReplicaSelector)

// Package admin is a generated GoMock package.
package admin

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	varlogpb "github.com/kakao/varlog/proto/varlogpb"
)

// MockReplicaSelector is a mock of ReplicaSelector interface.
type MockReplicaSelector struct {
	ctrl     *gomock.Controller
	recorder *MockReplicaSelectorMockRecorder
}

// MockReplicaSelectorMockRecorder is the mock recorder for MockReplicaSelector.
type MockReplicaSelectorMockRecorder struct {
	mock *MockReplicaSelector
}

// NewMockReplicaSelector creates a new mock instance.
func NewMockReplicaSelector(ctrl *gomock.Controller) *MockReplicaSelector {
	mock := &MockReplicaSelector{ctrl: ctrl}
	mock.recorder = &MockReplicaSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplicaSelector) EXPECT() *MockReplicaSelectorMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockReplicaSelector) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockReplicaSelectorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockReplicaSelector)(nil).Name))
}

// Select mocks base method.
func (m *MockReplicaSelector) Select(arg0 context.Context) ([]*varlogpb.ReplicaDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", arg0)
	ret0, _ := ret[0].([]*varlogpb.ReplicaDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockReplicaSelectorMockRecorder) Select(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockReplicaSelector)(nil).Select), arg0)
}
