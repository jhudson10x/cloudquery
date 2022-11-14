// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2 (interfaces: AutoscalingV2beta2Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	rest "k8s.io/client-go/rest"
)

// MockAutoscalingV2beta2Interface is a mock of AutoscalingV2beta2Interface interface.
type MockAutoscalingV2beta2Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAutoscalingV2beta2InterfaceMockRecorder
}

// MockAutoscalingV2beta2InterfaceMockRecorder is the mock recorder for MockAutoscalingV2beta2Interface.
type MockAutoscalingV2beta2InterfaceMockRecorder struct {
	mock *MockAutoscalingV2beta2Interface
}

// NewMockAutoscalingV2beta2Interface creates a new mock instance.
func NewMockAutoscalingV2beta2Interface(ctrl *gomock.Controller) *MockAutoscalingV2beta2Interface {
	mock := &MockAutoscalingV2beta2Interface{ctrl: ctrl}
	mock.recorder = &MockAutoscalingV2beta2InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutoscalingV2beta2Interface) EXPECT() *MockAutoscalingV2beta2InterfaceMockRecorder {
	return m.recorder
}

// HorizontalPodAutoscalers mocks base method.
func (m *MockAutoscalingV2beta2Interface) HorizontalPodAutoscalers(arg0 string) v2beta2.HorizontalPodAutoscalerInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HorizontalPodAutoscalers", arg0)
	ret0, _ := ret[0].(v2beta2.HorizontalPodAutoscalerInterface)
	return ret0
}

// HorizontalPodAutoscalers indicates an expected call of HorizontalPodAutoscalers.
func (mr *MockAutoscalingV2beta2InterfaceMockRecorder) HorizontalPodAutoscalers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HorizontalPodAutoscalers", reflect.TypeOf((*MockAutoscalingV2beta2Interface)(nil).HorizontalPodAutoscalers), arg0)
}

// RESTClient mocks base method.
func (m *MockAutoscalingV2beta2Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockAutoscalingV2beta2InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAutoscalingV2beta2Interface)(nil).RESTClient))
}