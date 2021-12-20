package http

import (
	entity "Calendar/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEventService is a mock of EventService interface.
type MockEventService struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceMockRecorder
}

// MockEventServiceMockRecorder is the mock recorder for MockEventService.
type MockEventServiceMockRecorder struct {
	mock *MockEventService
}

// NewMockEventService creates a new mock instance.
func NewMockEventService(ctrl *gomock.Controller) *MockEventService {
	mock := &MockEventService{ctrl: ctrl}
	mock.recorder = &MockEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventService) EXPECT() *MockEventServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockEventService) Add(event entity.Event) (entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", event)
	ret0, _ := ret[0].(entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockEventServiceMockRecorder) Add(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockEventService)(nil).Add), event)
}

// Delete mocks base method.
func (m *MockEventService) Delete(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", id)
}

// Delete indicates an expected call of Delete.
func (mr *MockEventServiceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEventService)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockEventService) GetAll() ([]entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockEventServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockEventService)(nil).GetAll))
}

// GetOne mocks base method.
func (m *MockEventService) GetOne(id string) (entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", id)
	ret0, _ := ret[0].(entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne.
func (mr *MockEventServiceMockRecorder) GetOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockEventService)(nil).GetOne), id)
}

// Update mocks base method.
func (m *MockEventService) Update(event entity.Event) (entity.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", event)
	ret0, _ := ret[0].(entity.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEventServiceMockRecorder) Update(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEventService)(nil).Update), event)
}
