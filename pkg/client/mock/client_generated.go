// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/IBM-Cloud/bluemix-go/models"
	p_cloud_p_vm_instances "github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	models0 "github.com/IBM-Cloud/power-go-client/power/models"
	gomock "github.com/golang/mock/gomock"
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

// CreateInstance mocks base method.
func (m *MockClient) CreateInstance(arg0 *p_cloud_p_vm_instances.PcloudPvminstancesPostParams) (*models0.PVMInstanceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance", arg0)
	ret0, _ := ret[0].(*models0.PVMInstanceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockClientMockRecorder) CreateInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockClient)(nil).CreateInstance), arg0)
}

// DeleteInstance mocks base method.
func (m *MockClient) DeleteInstance(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstance", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInstance indicates an expected call of DeleteInstance.
func (mr *MockClientMockRecorder) DeleteInstance(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstance", reflect.TypeOf((*MockClient)(nil).DeleteInstance), id)
}

// GetCloudServiceInstances mocks base method.
func (m *MockClient) GetCloudServiceInstances() ([]models.ServiceInstanceV2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudServiceInstances")
	ret0, _ := ret[0].([]models.ServiceInstanceV2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudServiceInstances indicates an expected call of GetCloudServiceInstances.
func (mr *MockClientMockRecorder) GetCloudServiceInstances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudServiceInstances", reflect.TypeOf((*MockClient)(nil).GetCloudServiceInstances))
}

// GetImages mocks base method.
func (m *MockClient) GetImages() (*models0.Images, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImages")
	ret0, _ := ret[0].(*models0.Images)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImages indicates an expected call of GetImages.
func (mr *MockClientMockRecorder) GetImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImages", reflect.TypeOf((*MockClient)(nil).GetImages))
}

// GetInstance mocks base method.
func (m *MockClient) GetInstance(id string) (*models0.PVMInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", id)
	ret0, _ := ret[0].(*models0.PVMInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockClientMockRecorder) GetInstance(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockClient)(nil).GetInstance), id)
}

// GetInstanceByName mocks base method.
func (m *MockClient) GetInstanceByName(name string) (*models0.PVMInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceByName", name)
	ret0, _ := ret[0].(*models0.PVMInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceByName indicates an expected call of GetInstanceByName.
func (mr *MockClientMockRecorder) GetInstanceByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceByName", reflect.TypeOf((*MockClient)(nil).GetInstanceByName), name)
}

// GetInstances mocks base method.
func (m *MockClient) GetInstances() (*models0.PVMInstances, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstances")
	ret0, _ := ret[0].(*models0.PVMInstances)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstances indicates an expected call of GetInstances.
func (mr *MockClientMockRecorder) GetInstances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstances", reflect.TypeOf((*MockClient)(nil).GetInstances))
}

// GetNetworks mocks base method.
func (m *MockClient) GetNetworks() (*models0.Networks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworks")
	ret0, _ := ret[0].(*models0.Networks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworks indicates an expected call of GetNetworks.
func (mr *MockClientMockRecorder) GetNetworks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworks", reflect.TypeOf((*MockClient)(nil).GetNetworks))
}
