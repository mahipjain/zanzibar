// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/zanzibar/examples/example-gateway/build/clients/google-now (interfaces: Client)

// Package clientmock is a generated GoMock package.
package clientmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	googlenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/googlenow/googlenow"
	zanzibar "github.com/uber/zanzibar/runtime"
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

// AddCredentials mocks base method.
func (m *MockClient) AddCredentials(arg0 context.Context, arg1 map[string]string, arg2 *googlenow.GoogleNowService_AddCredentials_Args) (context.Context, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCredentials", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddCredentials indicates an expected call of AddCredentials.
func (mr *MockClientMockRecorder) AddCredentials(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCredentials", reflect.TypeOf((*MockClient)(nil).AddCredentials), arg0, arg1, arg2)
}

// CheckCredentials mocks base method.
func (m *MockClient) CheckCredentials(arg0 context.Context, arg1 map[string]string) (context.Context, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCredentials", arg0, arg1)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(map[string]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckCredentials indicates an expected call of CheckCredentials.
func (mr *MockClientMockRecorder) CheckCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCredentials", reflect.TypeOf((*MockClient)(nil).CheckCredentials), arg0, arg1)
}

// HTTPClient mocks base method.
func (m *MockClient) HTTPClient() *zanzibar.HTTPClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*zanzibar.HTTPClient)
	return ret0
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockClientMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockClient)(nil).HTTPClient))
}
