// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/linkall-labs/vanus/client/pkg/api"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Disconnect mocks base method.
func (m *MockClient) Disconnect(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnect", ctx)
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockClientMockRecorder) Disconnect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockClient)(nil).Disconnect), ctx)
}

// Eventbus mocks base method.
func (m *MockClient) Eventbus(ctx context.Context, ebName string) api.Eventbus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventbusService", ctx, ebName)
	ret0, _ := ret[0].(api.Eventbus)
	return ret0
}

// Eventbus indicates an expected call of Eventbus.
func (mr *MockClientMockRecorder) Eventbus(ctx, ebName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventbusService", reflect.TypeOf((*MockClient)(nil).Eventbus), ctx, ebName)
}
