package repository

import (
	"Calendar/entity"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"reflect"
)

// MockSqliteRepo is a mock of SqliteRepo interface.
type MockSqliteRepo struct {
	ctrl     *gomock.Controller
	recorder *MockSqliteRepoMockRecorder
}

// MockSqliteRepoMockRecorder is the mock recorder for MockSqliteRepo.
type MockSqliteRepoMockRecorder struct {
	mock *MockSqliteRepo
}

// NewMockSqliteRepo creates a new mock instance.
func NewMockSqliteRepo(ctrl *gomock.Controller) *MockSqliteRepo {
	mock := &MockSqliteRepo{ctrl: ctrl}
	mock.recorder = &MockSqliteRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSqliteRepo) EXPECT() *MockSqliteRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSqliteRepo) Create(arg0 *entity.User) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", arg0)
}

// Create indicates an expected call of Create.
func (mr *MockSqliteRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSqliteRepo)(nil).Create), arg0)
}

// GetEmail mocks base method.
func (m *MockSqliteRepo) GetEmail(email string) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail", email)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockSqliteRepoMockRecorder) GetEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockSqliteRepo)(nil).GetEmail), email)
}

// MockgormConnection is a mock of gormConnection interface.
type MockgormConnection struct {
	ctrl     *gomock.Controller
	recorder *MockgormConnectionMockRecorder
}

// MockgormConnectionMockRecorder is the mock recorder for MockgormConnection.
type MockgormConnectionMockRecorder struct {
	mock *MockgormConnection
}

// NewMockgormConnection creates a new mock instance.
func NewMockgormConnection(ctrl *gomock.Controller) *MockgormConnection {
	mock := &MockgormConnection{ctrl: ctrl}
	mock.recorder = &MockgormConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockgormConnection) EXPECT() *MockgormConnectionMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockgormConnection) Create(value interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockgormConnectionMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockgormConnection)(nil).Create), value)
}

// Where mocks base method.
func (m *MockgormConnection) Where(query interface{}, args ...interface{}) gormScanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(gormScanner)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockgormConnectionMockRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockgormConnection)(nil).Where), varargs...)
}

// MockgormScanner is a mock of gormScanner interface.
type MockgormScanner struct {
	ctrl     *gomock.Controller
	recorder *MockgormScannerMockRecorder
}

// MockgormScannerMockRecorder is the mock recorder for MockgormScanner.
type MockgormScannerMockRecorder struct {
	mock *MockgormScanner
}

// NewMockgormScanner creates a new mock instance.
func NewMockgormScanner(ctrl *gomock.Controller) *MockgormScanner {
	mock := &MockgormScanner{ctrl: ctrl}
	mock.recorder = &MockgormScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockgormScanner) EXPECT() *MockgormScannerMockRecorder {
	return m.recorder
}

// First mocks base method.
func (m *MockgormScanner) First(dest interface{}, conds ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest}
	for _, a := range conds {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockgormScannerMockRecorder) First(dest interface{}, conds ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest}, conds...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockgormScanner)(nil).First), varargs...)
}
