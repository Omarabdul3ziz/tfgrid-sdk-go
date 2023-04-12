// Code generated by MockGen. DO NOT EDIT.
// Source: ./deployer/deployer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gridtypes "github.com/threefoldtech/zos/pkg/gridtypes"
)

// MockDeployer is a mock of Deployer interface.
type MockDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockDeployerMockRecorder
}

// MockDeployerMockRecorder is the mock recorder for MockDeployer.
type MockDeployerMockRecorder struct {
	mock *MockDeployer
}

// NewMockDeployer creates a new mock instance.
func NewMockDeployer(ctrl *gomock.Controller) *MockDeployer {
	mock := &MockDeployer{ctrl: ctrl}
	mock.recorder = &MockDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployer) EXPECT() *MockDeployerMockRecorder {
	return m.recorder
}

// Cancel mocks base method.
func (m *MockDeployer) Cancel(ctx context.Context, contractID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", ctx, contractID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockDeployerMockRecorder) Cancel(ctx, contractID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockDeployer)(nil).Cancel), ctx, contractID)
}

// Deploy mocks base method.
func (m *MockDeployer) Deploy(ctx context.Context, oldDeploymentIDs map[uint32]uint64, newDeployments map[uint32]gridtypes.Deployment, newDeploymentSolutionProvider map[uint32]*uint64) (map[uint32]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", ctx, oldDeploymentIDs, newDeployments, newDeploymentSolutionProvider)
	ret0, _ := ret[0].(map[uint32]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deploy indicates an expected call of Deploy.
func (mr *MockDeployerMockRecorder) Deploy(ctx, oldDeploymentIDs, newDeployments, newDeploymentSolutionProvider interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDeployer)(nil).Deploy), ctx, oldDeploymentIDs, newDeployments, newDeploymentSolutionProvider)
}

// GetDeployments mocks base method.
func (m *MockDeployer) GetDeployments(ctx context.Context, dls map[uint32]uint64) (map[uint32]gridtypes.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployments", ctx, dls)
	ret0, _ := ret[0].(map[uint32]gridtypes.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployments indicates an expected call of GetDeployments.
func (mr *MockDeployerMockRecorder) GetDeployments(ctx, dls interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployments", reflect.TypeOf((*MockDeployer)(nil).GetDeployments), ctx, dls)
}