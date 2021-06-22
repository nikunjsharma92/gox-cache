// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mockGoxCache is a generated GoMock package.
package mockGoxCache

import (
	context "context"
	reflect "reflect"

	gox "github.com/devlibx/gox-base"
	goxCache "github.com/devlibx/gox-cache"
	gomock "github.com/golang/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCache) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCacheMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCache)(nil).Close))
}

// Get mocks base method.
func (m *MockCache) Get(ctx context.Context, key string) (interface{}, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), ctx, key)
}

// GetAsMap mocks base method.
func (m *MockCache) GetAsMap(ctx context.Context, key string) (gox.StringObjectMap, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAsMap", ctx, key)
	ret0, _ := ret[0].(gox.StringObjectMap)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAsMap indicates an expected call of GetAsMap.
func (mr *MockCacheMockRecorder) GetAsMap(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAsMap", reflect.TypeOf((*MockCache)(nil).GetAsMap), ctx, key)
}

// IsRunning mocks base method.
func (m *MockCache) IsRunning(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsRunning indicates an expected call of IsRunning.
func (mr *MockCacheMockRecorder) IsRunning(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockCache)(nil).IsRunning), ctx)
}

// Publish mocks base method.
func (m *MockCache) Publish(ctx context.Context, data gox.StringObjectMap) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, data)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockCacheMockRecorder) Publish(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockCache)(nil).Publish), ctx, data)
}

// Put mocks base method.
func (m *MockCache) Put(ctx context.Context, key string, data interface{}, ttlInSec int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", ctx, key, data, ttlInSec)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockCacheMockRecorder) Put(ctx, key, data, ttlInSec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockCache)(nil).Put), ctx, key, data, ttlInSec)
}

// Subscribe mocks base method.
func (m *MockCache) Subscribe(ctx context.Context, callback goxCache.SubscribeCallbackFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockCacheMockRecorder) Subscribe(ctx, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockCache)(nil).Subscribe), ctx, callback)
}

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRegistry) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRegistryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRegistry)(nil).Close))
}

// GetCache mocks base method.
func (m *MockRegistry) GetCache(name string) (goxCache.Cache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache", name)
	ret0, _ := ret[0].(goxCache.Cache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCache indicates an expected call of GetCache.
func (mr *MockRegistryMockRecorder) GetCache(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockRegistry)(nil).GetCache), name)
}

// RegisterCache mocks base method.
func (m *MockRegistry) RegisterCache(config *goxCache.Config) (goxCache.Cache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCache", config)
	ret0, _ := ret[0].(goxCache.Cache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterCache indicates an expected call of RegisterCache.
func (mr *MockRegistryMockRecorder) RegisterCache(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCache", reflect.TypeOf((*MockRegistry)(nil).RegisterCache), config)
}
