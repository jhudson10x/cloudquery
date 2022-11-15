// Code generated by MockGen. DO NOT EDIT.
// Source: kafka.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	kafka "github.com/aws/aws-sdk-go-v2/service/kafka"
	gomock "github.com/golang/mock/gomock"
)

// MockKafkaClient is a mock of KafkaClient interface.
type MockKafkaClient struct {
	ctrl     *gomock.Controller
	recorder *MockKafkaClientMockRecorder
}

// MockKafkaClientMockRecorder is the mock recorder for MockKafkaClient.
type MockKafkaClientMockRecorder struct {
	mock *MockKafkaClient
}

// NewMockKafkaClient creates a new mock instance.
func NewMockKafkaClient(ctrl *gomock.Controller) *MockKafkaClient {
	mock := &MockKafkaClient{ctrl: ctrl}
	mock.recorder = &MockKafkaClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKafkaClient) EXPECT() *MockKafkaClientMockRecorder {
	return m.recorder
}

// DescribeCluster mocks base method.
func (m *MockKafkaClient) DescribeCluster(arg0 context.Context, arg1 *kafka.DescribeClusterInput, arg2 ...func(*kafka.Options)) (*kafka.DescribeClusterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCluster", varargs...)
	ret0, _ := ret[0].(*kafka.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCluster indicates an expected call of DescribeCluster.
func (mr *MockKafkaClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCluster", reflect.TypeOf((*MockKafkaClient)(nil).DescribeCluster), varargs...)
}

// DescribeClusterOperation mocks base method.
func (m *MockKafkaClient) DescribeClusterOperation(arg0 context.Context, arg1 *kafka.DescribeClusterOperationInput, arg2 ...func(*kafka.Options)) (*kafka.DescribeClusterOperationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusterOperation", varargs...)
	ret0, _ := ret[0].(*kafka.DescribeClusterOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeClusterOperation indicates an expected call of DescribeClusterOperation.
func (mr *MockKafkaClientMockRecorder) DescribeClusterOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusterOperation", reflect.TypeOf((*MockKafkaClient)(nil).DescribeClusterOperation), varargs...)
}

// DescribeClusterV2 mocks base method.
func (m *MockKafkaClient) DescribeClusterV2(arg0 context.Context, arg1 *kafka.DescribeClusterV2Input, arg2 ...func(*kafka.Options)) (*kafka.DescribeClusterV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusterV2", varargs...)
	ret0, _ := ret[0].(*kafka.DescribeClusterV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeClusterV2 indicates an expected call of DescribeClusterV2.
func (mr *MockKafkaClientMockRecorder) DescribeClusterV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusterV2", reflect.TypeOf((*MockKafkaClient)(nil).DescribeClusterV2), varargs...)
}

// DescribeConfiguration mocks base method.
func (m *MockKafkaClient) DescribeConfiguration(arg0 context.Context, arg1 *kafka.DescribeConfigurationInput, arg2 ...func(*kafka.Options)) (*kafka.DescribeConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfiguration", varargs...)
	ret0, _ := ret[0].(*kafka.DescribeConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfiguration indicates an expected call of DescribeConfiguration.
func (mr *MockKafkaClientMockRecorder) DescribeConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfiguration", reflect.TypeOf((*MockKafkaClient)(nil).DescribeConfiguration), varargs...)
}

// DescribeConfigurationRevision mocks base method.
func (m *MockKafkaClient) DescribeConfigurationRevision(arg0 context.Context, arg1 *kafka.DescribeConfigurationRevisionInput, arg2 ...func(*kafka.Options)) (*kafka.DescribeConfigurationRevisionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfigurationRevision", varargs...)
	ret0, _ := ret[0].(*kafka.DescribeConfigurationRevisionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeConfigurationRevision indicates an expected call of DescribeConfigurationRevision.
func (mr *MockKafkaClientMockRecorder) DescribeConfigurationRevision(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfigurationRevision", reflect.TypeOf((*MockKafkaClient)(nil).DescribeConfigurationRevision), varargs...)
}

// GetBootstrapBrokers mocks base method.
func (m *MockKafkaClient) GetBootstrapBrokers(arg0 context.Context, arg1 *kafka.GetBootstrapBrokersInput, arg2 ...func(*kafka.Options)) (*kafka.GetBootstrapBrokersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBootstrapBrokers", varargs...)
	ret0, _ := ret[0].(*kafka.GetBootstrapBrokersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBootstrapBrokers indicates an expected call of GetBootstrapBrokers.
func (mr *MockKafkaClientMockRecorder) GetBootstrapBrokers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBootstrapBrokers", reflect.TypeOf((*MockKafkaClient)(nil).GetBootstrapBrokers), varargs...)
}

// GetCompatibleKafkaVersions mocks base method.
func (m *MockKafkaClient) GetCompatibleKafkaVersions(arg0 context.Context, arg1 *kafka.GetCompatibleKafkaVersionsInput, arg2 ...func(*kafka.Options)) (*kafka.GetCompatibleKafkaVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCompatibleKafkaVersions", varargs...)
	ret0, _ := ret[0].(*kafka.GetCompatibleKafkaVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompatibleKafkaVersions indicates an expected call of GetCompatibleKafkaVersions.
func (mr *MockKafkaClientMockRecorder) GetCompatibleKafkaVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompatibleKafkaVersions", reflect.TypeOf((*MockKafkaClient)(nil).GetCompatibleKafkaVersions), varargs...)
}

// ListClusterOperations mocks base method.
func (m *MockKafkaClient) ListClusterOperations(arg0 context.Context, arg1 *kafka.ListClusterOperationsInput, arg2 ...func(*kafka.Options)) (*kafka.ListClusterOperationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusterOperations", varargs...)
	ret0, _ := ret[0].(*kafka.ListClusterOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusterOperations indicates an expected call of ListClusterOperations.
func (mr *MockKafkaClientMockRecorder) ListClusterOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusterOperations", reflect.TypeOf((*MockKafkaClient)(nil).ListClusterOperations), varargs...)
}

// ListClusters mocks base method.
func (m *MockKafkaClient) ListClusters(arg0 context.Context, arg1 *kafka.ListClustersInput, arg2 ...func(*kafka.Options)) (*kafka.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*kafka.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockKafkaClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockKafkaClient)(nil).ListClusters), varargs...)
}

// ListClustersV2 mocks base method.
func (m *MockKafkaClient) ListClustersV2(arg0 context.Context, arg1 *kafka.ListClustersV2Input, arg2 ...func(*kafka.Options)) (*kafka.ListClustersV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClustersV2", varargs...)
	ret0, _ := ret[0].(*kafka.ListClustersV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClustersV2 indicates an expected call of ListClustersV2.
func (mr *MockKafkaClientMockRecorder) ListClustersV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClustersV2", reflect.TypeOf((*MockKafkaClient)(nil).ListClustersV2), varargs...)
}

// ListConfigurationRevisions mocks base method.
func (m *MockKafkaClient) ListConfigurationRevisions(arg0 context.Context, arg1 *kafka.ListConfigurationRevisionsInput, arg2 ...func(*kafka.Options)) (*kafka.ListConfigurationRevisionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConfigurationRevisions", varargs...)
	ret0, _ := ret[0].(*kafka.ListConfigurationRevisionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurationRevisions indicates an expected call of ListConfigurationRevisions.
func (mr *MockKafkaClientMockRecorder) ListConfigurationRevisions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurationRevisions", reflect.TypeOf((*MockKafkaClient)(nil).ListConfigurationRevisions), varargs...)
}

// ListConfigurations mocks base method.
func (m *MockKafkaClient) ListConfigurations(arg0 context.Context, arg1 *kafka.ListConfigurationsInput, arg2 ...func(*kafka.Options)) (*kafka.ListConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConfigurations", varargs...)
	ret0, _ := ret[0].(*kafka.ListConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurations indicates an expected call of ListConfigurations.
func (mr *MockKafkaClientMockRecorder) ListConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurations", reflect.TypeOf((*MockKafkaClient)(nil).ListConfigurations), varargs...)
}

// ListKafkaVersions mocks base method.
func (m *MockKafkaClient) ListKafkaVersions(arg0 context.Context, arg1 *kafka.ListKafkaVersionsInput, arg2 ...func(*kafka.Options)) (*kafka.ListKafkaVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListKafkaVersions", varargs...)
	ret0, _ := ret[0].(*kafka.ListKafkaVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKafkaVersions indicates an expected call of ListKafkaVersions.
func (mr *MockKafkaClientMockRecorder) ListKafkaVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKafkaVersions", reflect.TypeOf((*MockKafkaClient)(nil).ListKafkaVersions), varargs...)
}

// ListNodes mocks base method.
func (m *MockKafkaClient) ListNodes(arg0 context.Context, arg1 *kafka.ListNodesInput, arg2 ...func(*kafka.Options)) (*kafka.ListNodesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNodes", varargs...)
	ret0, _ := ret[0].(*kafka.ListNodesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes.
func (mr *MockKafkaClientMockRecorder) ListNodes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockKafkaClient)(nil).ListNodes), varargs...)
}

// ListScramSecrets mocks base method.
func (m *MockKafkaClient) ListScramSecrets(arg0 context.Context, arg1 *kafka.ListScramSecretsInput, arg2 ...func(*kafka.Options)) (*kafka.ListScramSecretsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListScramSecrets", varargs...)
	ret0, _ := ret[0].(*kafka.ListScramSecretsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListScramSecrets indicates an expected call of ListScramSecrets.
func (mr *MockKafkaClientMockRecorder) ListScramSecrets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListScramSecrets", reflect.TypeOf((*MockKafkaClient)(nil).ListScramSecrets), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockKafkaClient) ListTagsForResource(arg0 context.Context, arg1 *kafka.ListTagsForResourceInput, arg2 ...func(*kafka.Options)) (*kafka.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*kafka.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockKafkaClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockKafkaClient)(nil).ListTagsForResource), varargs...)
}