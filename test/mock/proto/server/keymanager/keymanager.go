// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/server/keymanager (interfaces: KeyManager,KeyManagerServer)

// Package mock_keymanager is a generated GoMock package.
package mock_keymanager

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	keymanager "github.com/spiffe/spire/proto/server/keymanager"
	reflect "reflect"
)

// MockKeyManager is a mock of KeyManager interface
type MockKeyManager struct {
	ctrl     *gomock.Controller
	recorder *MockKeyManagerMockRecorder
}

// MockKeyManagerMockRecorder is the mock recorder for MockKeyManager
type MockKeyManagerMockRecorder struct {
	mock *MockKeyManager
}

// NewMockKeyManager creates a new mock instance
func NewMockKeyManager(ctrl *gomock.Controller) *MockKeyManager {
	mock := &MockKeyManager{ctrl: ctrl}
	mock.recorder = &MockKeyManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyManager) EXPECT() *MockKeyManagerMockRecorder {
	return m.recorder
}

// GenerateKey mocks base method
func (m *MockKeyManager) GenerateKey(arg0 context.Context, arg1 *keymanager.GenerateKeyRequest) (*keymanager.GenerateKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateKey", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.GenerateKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateKey indicates an expected call of GenerateKey
func (mr *MockKeyManagerMockRecorder) GenerateKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKey", reflect.TypeOf((*MockKeyManager)(nil).GenerateKey), arg0, arg1)
}

// GetPublicKey mocks base method
func (m *MockKeyManager) GetPublicKey(arg0 context.Context, arg1 *keymanager.GetPublicKeyRequest) (*keymanager.GetPublicKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.GetPublicKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey
func (mr *MockKeyManagerMockRecorder) GetPublicKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockKeyManager)(nil).GetPublicKey), arg0, arg1)
}

// GetPublicKeys mocks base method
func (m *MockKeyManager) GetPublicKeys(arg0 context.Context, arg1 *keymanager.GetPublicKeysRequest) (*keymanager.GetPublicKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKeys", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.GetPublicKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKeys indicates an expected call of GetPublicKeys
func (mr *MockKeyManagerMockRecorder) GetPublicKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKeys", reflect.TypeOf((*MockKeyManager)(nil).GetPublicKeys), arg0, arg1)
}

// SignData mocks base method
func (m *MockKeyManager) SignData(arg0 context.Context, arg1 *keymanager.SignDataRequest) (*keymanager.SignDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignData", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.SignDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignData indicates an expected call of SignData
func (mr *MockKeyManagerMockRecorder) SignData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignData", reflect.TypeOf((*MockKeyManager)(nil).SignData), arg0, arg1)
}

// MockKeyManagerServer is a mock of KeyManagerServer interface
type MockKeyManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockKeyManagerServerMockRecorder
}

// MockKeyManagerServerMockRecorder is the mock recorder for MockKeyManagerServer
type MockKeyManagerServerMockRecorder struct {
	mock *MockKeyManagerServer
}

// NewMockKeyManagerServer creates a new mock instance
func NewMockKeyManagerServer(ctrl *gomock.Controller) *MockKeyManagerServer {
	mock := &MockKeyManagerServer{ctrl: ctrl}
	mock.recorder = &MockKeyManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyManagerServer) EXPECT() *MockKeyManagerServerMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockKeyManagerServer) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockKeyManagerServerMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockKeyManagerServer)(nil).Configure), arg0, arg1)
}

// GenerateKey mocks base method
func (m *MockKeyManagerServer) GenerateKey(arg0 context.Context, arg1 *keymanager.GenerateKeyRequest) (*keymanager.GenerateKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateKey", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.GenerateKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateKey indicates an expected call of GenerateKey
func (mr *MockKeyManagerServerMockRecorder) GenerateKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKey", reflect.TypeOf((*MockKeyManagerServer)(nil).GenerateKey), arg0, arg1)
}

// GetPluginInfo mocks base method
func (m *MockKeyManagerServer) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockKeyManagerServerMockRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockKeyManagerServer)(nil).GetPluginInfo), arg0, arg1)
}

// GetPublicKey mocks base method
func (m *MockKeyManagerServer) GetPublicKey(arg0 context.Context, arg1 *keymanager.GetPublicKeyRequest) (*keymanager.GetPublicKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.GetPublicKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey
func (mr *MockKeyManagerServerMockRecorder) GetPublicKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockKeyManagerServer)(nil).GetPublicKey), arg0, arg1)
}

// GetPublicKeys mocks base method
func (m *MockKeyManagerServer) GetPublicKeys(arg0 context.Context, arg1 *keymanager.GetPublicKeysRequest) (*keymanager.GetPublicKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKeys", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.GetPublicKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKeys indicates an expected call of GetPublicKeys
func (mr *MockKeyManagerServerMockRecorder) GetPublicKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKeys", reflect.TypeOf((*MockKeyManagerServer)(nil).GetPublicKeys), arg0, arg1)
}

// SignData mocks base method
func (m *MockKeyManagerServer) SignData(arg0 context.Context, arg1 *keymanager.SignDataRequest) (*keymanager.SignDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignData", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.SignDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignData indicates an expected call of SignData
func (mr *MockKeyManagerServerMockRecorder) SignData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignData", reflect.TypeOf((*MockKeyManagerServer)(nil).SignData), arg0, arg1)
}
