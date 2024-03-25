// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/atlas/deployments/options (interfaces: DeploymentTelemetry)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeploymentTelemetry is a mock of DeploymentTelemetry interface.
type MockDeploymentTelemetry struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentTelemetryMockRecorder
}

// MockDeploymentTelemetryMockRecorder is the mock recorder for MockDeploymentTelemetry.
type MockDeploymentTelemetryMockRecorder struct {
	mock *MockDeploymentTelemetry
}

// NewMockDeploymentTelemetry creates a new mock instance.
func NewMockDeploymentTelemetry(ctrl *gomock.Controller) *MockDeploymentTelemetry {
	mock := &MockDeploymentTelemetry{ctrl: ctrl}
	mock.recorder = &MockDeploymentTelemetryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentTelemetry) EXPECT() *MockDeploymentTelemetryMockRecorder {
	return m.recorder
}

// AppendDeploymentType mocks base method.
func (m *MockDeploymentTelemetry) AppendDeploymentType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendDeploymentType")
}

// AppendDeploymentType indicates an expected call of AppendDeploymentType.
func (mr *MockDeploymentTelemetryMockRecorder) AppendDeploymentType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendDeploymentType", reflect.TypeOf((*MockDeploymentTelemetry)(nil).AppendDeploymentType))
}