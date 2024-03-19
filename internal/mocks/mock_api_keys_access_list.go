// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: OrganizationAPIKeyAccessListCreator,OrganizationAPIKeyAccessListDeleter,OrganizationAPIKeyAccessListLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

// MockOrganizationAPIKeyAccessListCreator is a mock of OrganizationAPIKeyAccessListCreator interface.
type MockOrganizationAPIKeyAccessListCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyAccessListCreatorMockRecorder
}

// MockOrganizationAPIKeyAccessListCreatorMockRecorder is the mock recorder for MockOrganizationAPIKeyAccessListCreator.
type MockOrganizationAPIKeyAccessListCreatorMockRecorder struct {
	mock *MockOrganizationAPIKeyAccessListCreator
}

// NewMockOrganizationAPIKeyAccessListCreator creates a new mock instance.
func NewMockOrganizationAPIKeyAccessListCreator(ctrl *gomock.Controller) *MockOrganizationAPIKeyAccessListCreator {
	mock := &MockOrganizationAPIKeyAccessListCreator{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIKeyAccessListCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPIKeyAccessListCreator) EXPECT() *MockOrganizationAPIKeyAccessListCreatorMockRecorder {
	return m.recorder
}

// CreateOrganizationAPIKeyAccessList mocks base method.
func (m *MockOrganizationAPIKeyAccessListCreator) CreateOrganizationAPIKeyAccessList(arg0 *admin.CreateApiKeyAccessListApiParams) (*admin.PaginatedApiUserAccessListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganizationAPIKeyAccessList", arg0)
	ret0, _ := ret[0].(*admin.PaginatedApiUserAccessListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganizationAPIKeyAccessList indicates an expected call of CreateOrganizationAPIKeyAccessList.
func (mr *MockOrganizationAPIKeyAccessListCreatorMockRecorder) CreateOrganizationAPIKeyAccessList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganizationAPIKeyAccessList", reflect.TypeOf((*MockOrganizationAPIKeyAccessListCreator)(nil).CreateOrganizationAPIKeyAccessList), arg0)
}

// MockOrganizationAPIKeyAccessListDeleter is a mock of OrganizationAPIKeyAccessListDeleter interface.
type MockOrganizationAPIKeyAccessListDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyAccessListDeleterMockRecorder
}

// MockOrganizationAPIKeyAccessListDeleterMockRecorder is the mock recorder for MockOrganizationAPIKeyAccessListDeleter.
type MockOrganizationAPIKeyAccessListDeleterMockRecorder struct {
	mock *MockOrganizationAPIKeyAccessListDeleter
}

// NewMockOrganizationAPIKeyAccessListDeleter creates a new mock instance.
func NewMockOrganizationAPIKeyAccessListDeleter(ctrl *gomock.Controller) *MockOrganizationAPIKeyAccessListDeleter {
	mock := &MockOrganizationAPIKeyAccessListDeleter{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIKeyAccessListDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPIKeyAccessListDeleter) EXPECT() *MockOrganizationAPIKeyAccessListDeleterMockRecorder {
	return m.recorder
}

// DeleteOrganizationAPIKeyAccessList mocks base method.
func (m *MockOrganizationAPIKeyAccessListDeleter) DeleteOrganizationAPIKeyAccessList(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrganizationAPIKeyAccessList", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganizationAPIKeyAccessList indicates an expected call of DeleteOrganizationAPIKeyAccessList.
func (mr *MockOrganizationAPIKeyAccessListDeleterMockRecorder) DeleteOrganizationAPIKeyAccessList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganizationAPIKeyAccessList", reflect.TypeOf((*MockOrganizationAPIKeyAccessListDeleter)(nil).DeleteOrganizationAPIKeyAccessList), arg0, arg1, arg2)
}

// MockOrganizationAPIKeyAccessListLister is a mock of OrganizationAPIKeyAccessListLister interface.
type MockOrganizationAPIKeyAccessListLister struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationAPIKeyAccessListListerMockRecorder
}

// MockOrganizationAPIKeyAccessListListerMockRecorder is the mock recorder for MockOrganizationAPIKeyAccessListLister.
type MockOrganizationAPIKeyAccessListListerMockRecorder struct {
	mock *MockOrganizationAPIKeyAccessListLister
}

// NewMockOrganizationAPIKeyAccessListLister creates a new mock instance.
func NewMockOrganizationAPIKeyAccessListLister(ctrl *gomock.Controller) *MockOrganizationAPIKeyAccessListLister {
	mock := &MockOrganizationAPIKeyAccessListLister{ctrl: ctrl}
	mock.recorder = &MockOrganizationAPIKeyAccessListListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationAPIKeyAccessListLister) EXPECT() *MockOrganizationAPIKeyAccessListListerMockRecorder {
	return m.recorder
}

// OrganizationAPIKeyAccessLists mocks base method.
func (m *MockOrganizationAPIKeyAccessListLister) OrganizationAPIKeyAccessLists(arg0 *admin.ListApiKeyAccessListsEntriesApiParams) (*admin.PaginatedApiUserAccessListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrganizationAPIKeyAccessLists", arg0)
	ret0, _ := ret[0].(*admin.PaginatedApiUserAccessListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrganizationAPIKeyAccessLists indicates an expected call of OrganizationAPIKeyAccessLists.
func (mr *MockOrganizationAPIKeyAccessListListerMockRecorder) OrganizationAPIKeyAccessLists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrganizationAPIKeyAccessLists", reflect.TypeOf((*MockOrganizationAPIKeyAccessListLister)(nil).OrganizationAPIKeyAccessLists), arg0)
}