// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/coordination/v1 (interfaces: LeasesGetter,LeaseInterface)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/coordination/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/coordination/v1"
	v12 "k8s.io/client-go/kubernetes/typed/coordination/v1"
)

// MockLeasesGetter is a mock of LeasesGetter interface.
type MockLeasesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockLeasesGetterMockRecorder
}

// MockLeasesGetterMockRecorder is the mock recorder for MockLeasesGetter.
type MockLeasesGetterMockRecorder struct {
	mock *MockLeasesGetter
}

// NewMockLeasesGetter creates a new mock instance.
func NewMockLeasesGetter(ctrl *gomock.Controller) *MockLeasesGetter {
	mock := &MockLeasesGetter{ctrl: ctrl}
	mock.recorder = &MockLeasesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeasesGetter) EXPECT() *MockLeasesGetterMockRecorder {
	return m.recorder
}

// Leases mocks base method.
func (m *MockLeasesGetter) Leases(arg0 string) v12.LeaseInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Leases", arg0)
	ret0, _ := ret[0].(v12.LeaseInterface)
	return ret0
}

// Leases indicates an expected call of Leases.
func (mr *MockLeasesGetterMockRecorder) Leases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leases", reflect.TypeOf((*MockLeasesGetter)(nil).Leases), arg0)
}

// MockLeaseInterface is a mock of LeaseInterface interface.
type MockLeaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLeaseInterfaceMockRecorder
}

// MockLeaseInterfaceMockRecorder is the mock recorder for MockLeaseInterface.
type MockLeaseInterfaceMockRecorder struct {
	mock *MockLeaseInterface
}

// NewMockLeaseInterface creates a new mock instance.
func NewMockLeaseInterface(ctrl *gomock.Controller) *MockLeaseInterface {
	mock := &MockLeaseInterface{ctrl: ctrl}
	mock.recorder = &MockLeaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaseInterface) EXPECT() *MockLeaseInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockLeaseInterface) Apply(arg0 context.Context, arg1 *v11.LeaseApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockLeaseInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockLeaseInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockLeaseInterface) Create(arg0 context.Context, arg1 *v1.Lease, arg2 v10.CreateOptions) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLeaseInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLeaseInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockLeaseInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLeaseInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLeaseInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockLeaseInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockLeaseInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockLeaseInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockLeaseInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLeaseInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLeaseInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockLeaseInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.LeaseList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.LeaseList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLeaseInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLeaseInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockLeaseInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockLeaseInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockLeaseInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockLeaseInterface) Update(arg0 context.Context, arg1 *v1.Lease, arg2 v10.UpdateOptions) (*v1.Lease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Lease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockLeaseInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLeaseInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockLeaseInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockLeaseInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockLeaseInterface)(nil).Watch), arg0, arg1)
}
