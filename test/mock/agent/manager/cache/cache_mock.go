// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/pkg/agent/manager/cache (interfaces: Cache)

// Package mock_cache is a generated GoMock package.
package mock_cache

import (
	gomock "github.com/golang/mock/gomock"
	client "github.com/spiffe/spire/pkg/agent/client"
	cache "github.com/spiffe/spire/pkg/agent/manager/cache"
	bundleutil "github.com/spiffe/spire/pkg/common/bundleutil"
	common "github.com/spiffe/spire/proto/common"
	reflect "reflect"
)

// MockCache is a mock of Cache interface
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Bundle mocks base method
func (m *MockCache) Bundle() *bundleutil.Bundle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bundle")
	ret0, _ := ret[0].(*bundleutil.Bundle)
	return ret0
}

// Bundle indicates an expected call of Bundle
func (mr *MockCacheMockRecorder) Bundle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bundle", reflect.TypeOf((*MockCache)(nil).Bundle))
}

// DeleteEntry mocks base method
func (m *MockCache) DeleteEntry(arg0 *common.RegistrationEntry) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntry", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteEntry indicates an expected call of DeleteEntry
func (mr *MockCacheMockRecorder) DeleteEntry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockCache)(nil).DeleteEntry), arg0)
}

// Entries mocks base method
func (m *MockCache) Entries() []*cache.Entry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entries")
	ret0, _ := ret[0].([]*cache.Entry)
	return ret0
}

// Entries indicates an expected call of Entries
func (mr *MockCacheMockRecorder) Entries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entries", reflect.TypeOf((*MockCache)(nil).Entries))
}

// FetchEntry mocks base method
func (m *MockCache) FetchEntry(arg0 string) *cache.Entry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchEntry", arg0)
	ret0, _ := ret[0].(*cache.Entry)
	return ret0
}

// FetchEntry indicates an expected call of FetchEntry
func (mr *MockCacheMockRecorder) FetchEntry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEntry", reflect.TypeOf((*MockCache)(nil).FetchEntry), arg0)
}

// FetchWorkloadUpdate mocks base method
func (m *MockCache) FetchWorkloadUpdate(arg0 cache.Selectors) *cache.WorkloadUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchWorkloadUpdate", arg0)
	ret0, _ := ret[0].(*cache.WorkloadUpdate)
	return ret0
}

// FetchWorkloadUpdate indicates an expected call of FetchWorkloadUpdate
func (mr *MockCacheMockRecorder) FetchWorkloadUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchWorkloadUpdate", reflect.TypeOf((*MockCache)(nil).FetchWorkloadUpdate), arg0)
}

// GetJWTSVID mocks base method
func (m *MockCache) GetJWTSVID(arg0 string, arg1 []string) (*client.JWTSVID, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJWTSVID", arg0, arg1)
	ret0, _ := ret[0].(*client.JWTSVID)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetJWTSVID indicates an expected call of GetJWTSVID
func (mr *MockCacheMockRecorder) GetJWTSVID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJWTSVID", reflect.TypeOf((*MockCache)(nil).GetJWTSVID), arg0, arg1)
}

// SetBundles mocks base method
func (m *MockCache) SetBundles(arg0 map[string]*bundleutil.Bundle) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBundles", arg0)
}

// SetBundles indicates an expected call of SetBundles
func (mr *MockCacheMockRecorder) SetBundles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBundles", reflect.TypeOf((*MockCache)(nil).SetBundles), arg0)
}

// SetEntry mocks base method
func (m *MockCache) SetEntry(arg0 *cache.Entry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetEntry", arg0)
}

// SetEntry indicates an expected call of SetEntry
func (mr *MockCacheMockRecorder) SetEntry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEntry", reflect.TypeOf((*MockCache)(nil).SetEntry), arg0)
}

// SetJWTSVID mocks base method
func (m *MockCache) SetJWTSVID(arg0 string, arg1 []string, arg2 *client.JWTSVID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetJWTSVID", arg0, arg1, arg2)
}

// SetJWTSVID indicates an expected call of SetJWTSVID
func (mr *MockCacheMockRecorder) SetJWTSVID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJWTSVID", reflect.TypeOf((*MockCache)(nil).SetJWTSVID), arg0, arg1, arg2)
}

// Subscribe mocks base method
func (m *MockCache) Subscribe(arg0 cache.Selectors) cache.Subscriber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(cache.Subscriber)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockCacheMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockCache)(nil).Subscribe), arg0)
}

// SubscribeToBundleChanges mocks base method
func (m *MockCache) SubscribeToBundleChanges() *cache.BundleStream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToBundleChanges")
	ret0, _ := ret[0].(*cache.BundleStream)
	return ret0
}

// SubscribeToBundleChanges indicates an expected call of SubscribeToBundleChanges
func (mr *MockCacheMockRecorder) SubscribeToBundleChanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToBundleChanges", reflect.TypeOf((*MockCache)(nil).SubscribeToBundleChanges))
}
