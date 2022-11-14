// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/apps/v1beta1 (interfaces: AppsV1beta1Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	rest "k8s.io/client-go/rest"
)

// MockAppsV1beta1Interface is a mock of AppsV1beta1Interface interface.
type MockAppsV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAppsV1beta1InterfaceMockRecorder
}

// MockAppsV1beta1InterfaceMockRecorder is the mock recorder for MockAppsV1beta1Interface.
type MockAppsV1beta1InterfaceMockRecorder struct {
	mock *MockAppsV1beta1Interface
}

// NewMockAppsV1beta1Interface creates a new mock instance.
func NewMockAppsV1beta1Interface(ctrl *gomock.Controller) *MockAppsV1beta1Interface {
	mock := &MockAppsV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAppsV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppsV1beta1Interface) EXPECT() *MockAppsV1beta1InterfaceMockRecorder {
	return m.recorder
}

// ControllerRevisions mocks base method.
func (m *MockAppsV1beta1Interface) ControllerRevisions(arg0 string) v1beta1.ControllerRevisionInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerRevisions", arg0)
	ret0, _ := ret[0].(v1beta1.ControllerRevisionInterface)
	return ret0
}

// ControllerRevisions indicates an expected call of ControllerRevisions.
func (mr *MockAppsV1beta1InterfaceMockRecorder) ControllerRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerRevisions", reflect.TypeOf((*MockAppsV1beta1Interface)(nil).ControllerRevisions), arg0)
}

// Deployments mocks base method.
func (m *MockAppsV1beta1Interface) Deployments(arg0 string) v1beta1.DeploymentInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployments", arg0)
	ret0, _ := ret[0].(v1beta1.DeploymentInterface)
	return ret0
}

// Deployments indicates an expected call of Deployments.
func (mr *MockAppsV1beta1InterfaceMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockAppsV1beta1Interface)(nil).Deployments), arg0)
}

// RESTClient mocks base method.
func (m *MockAppsV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockAppsV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAppsV1beta1Interface)(nil).RESTClient))
}

// StatefulSets mocks base method.
func (m *MockAppsV1beta1Interface) StatefulSets(arg0 string) v1beta1.StatefulSetInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatefulSets", arg0)
	ret0, _ := ret[0].(v1beta1.StatefulSetInterface)
	return ret0
}

// StatefulSets indicates an expected call of StatefulSets.
func (mr *MockAppsV1beta1InterfaceMockRecorder) StatefulSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatefulSets", reflect.TypeOf((*MockAppsV1beta1Interface)(nil).StatefulSets), arg0)
}