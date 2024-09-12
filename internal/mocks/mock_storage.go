// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/kamencov/go-musthave-diploma-tpl/internal/models"
)

// MockStorageServ is a mock of StorageServ interface.
type MockStorageServ struct {
	ctrl     *gomock.Controller
	recorder *MockStorageServMockRecorder
}

// MockStorageServMockRecorder is the mock recorder for MockStorageServ.
type MockStorageServMockRecorder struct {
	mock *MockStorageServ
}

// NewMockStorageServ creates a new mock instance.
func NewMockStorageServ(ctrl *gomock.Controller) *MockStorageServ {
	mock := &MockStorageServ{ctrl: ctrl}
	mock.recorder = &MockStorageServMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageServ) EXPECT() *MockStorageServMockRecorder {
	return m.recorder
}

// CheckTableUserLogin mocks base method.
func (m *MockStorageServ) CheckTableUserLogin(ctx context.Context, login string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTableUserLogin", ctx, login)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckTableUserLogin indicates an expected call of CheckTableUserLogin.
func (mr *MockStorageServMockRecorder) CheckTableUserLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTableUserLogin", reflect.TypeOf((*MockStorageServ)(nil).CheckTableUserLogin), ctx, login)
}

// CheckTableUserPassword mocks base method.
func (m *MockStorageServ) CheckTableUserPassword(ctx context.Context, password string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTableUserPassword", ctx, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckTableUserPassword indicates an expected call of CheckTableUserPassword.
func (mr *MockStorageServMockRecorder) CheckTableUserPassword(ctx, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTableUserPassword", reflect.TypeOf((*MockStorageServ)(nil).CheckTableUserPassword), ctx, password)
}

// Get mocks base method.
func (m *MockStorageServ) Get(query string, args ...interface{}) (*sql.Row, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStorageServMockRecorder) Get(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorageServ)(nil).Get), varargs...)
}

// GetAllUserOrders mocks base method.
func (m *MockStorageServ) GetAllUserOrders(login string) ([]*models.OrdersUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserOrders", login)
	ret0, _ := ret[0].([]*models.OrdersUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserOrders indicates an expected call of GetAllUserOrders.
func (mr *MockStorageServMockRecorder) GetAllUserOrders(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserOrders", reflect.TypeOf((*MockStorageServ)(nil).GetAllUserOrders), login)
}

// GetBalanceUser mocks base method.
func (m *MockStorageServ) GetBalanceUser(login string) (*models.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceUser", login)
	ret0, _ := ret[0].(*models.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalanceUser indicates an expected call of GetBalanceUser.
func (mr *MockStorageServMockRecorder) GetBalanceUser(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceUser", reflect.TypeOf((*MockStorageServ)(nil).GetBalanceUser), login)
}

// GetUserByAccessToken mocks base method.
func (m *MockStorageServ) GetUserByAccessToken(order, login string, now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAccessToken", order, login, now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserByAccessToken indicates an expected call of GetUserByAccessToken.
func (mr *MockStorageServMockRecorder) GetUserByAccessToken(order, login, now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAccessToken", reflect.TypeOf((*MockStorageServ)(nil).GetUserByAccessToken), order, login, now)
}

// Save mocks base method.
func (m *MockStorageServ) Save(query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Save", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockStorageServMockRecorder) Save(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockStorageServ)(nil).Save), varargs...)
}

// SaveTableUser mocks base method.
func (m *MockStorageServ) SaveTableUser(login, passwordHash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTableUser", login, passwordHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTableUser indicates an expected call of SaveTableUser.
func (mr *MockStorageServMockRecorder) SaveTableUser(login, passwordHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTableUser", reflect.TypeOf((*MockStorageServ)(nil).SaveTableUser), login, passwordHash)
}

// SaveTableUserAndUpdateToken mocks base method.
func (m *MockStorageServ) SaveTableUserAndUpdateToken(login, accessToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTableUserAndUpdateToken", login, accessToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTableUserAndUpdateToken indicates an expected call of SaveTableUserAndUpdateToken.
func (mr *MockStorageServMockRecorder) SaveTableUserAndUpdateToken(login, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTableUserAndUpdateToken", reflect.TypeOf((*MockStorageServ)(nil).SaveTableUserAndUpdateToken), login, accessToken)
}
