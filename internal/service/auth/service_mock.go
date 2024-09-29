// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/kamencov/go-musthave-diploma-tpl/internal/service/auth/entity"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// AuthUser mocks base method.
func (m *MockAuthService) AuthUser(login, password string) (entity.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUser", login, password)
	ret0, _ := ret[0].(entity.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthUser indicates an expected call of AuthUser.
func (mr *MockAuthServiceMockRecorder) AuthUser(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockAuthService)(nil).AuthUser), login, password)
}

// GeneratedTokens mocks base method.
func (m *MockAuthService) GeneratedTokens(login string) (entity.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratedTokens", login)
	ret0, _ := ret[0].(entity.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratedTokens indicates an expected call of GeneratedTokens.
func (mr *MockAuthServiceMockRecorder) GeneratedTokens(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratedTokens", reflect.TypeOf((*MockAuthService)(nil).GeneratedTokens), login)
}

// HashPassword mocks base method.
func (m *MockAuthService) HashPassword(password string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	return ret0
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockAuthServiceMockRecorder) HashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockAuthService)(nil).HashPassword), password)
}

// RefreshToken mocks base method.
func (m *MockAuthService) RefreshToken(token string) (entity.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", token)
	ret0, _ := ret[0].(entity.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockAuthServiceMockRecorder) RefreshToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockAuthService)(nil).RefreshToken), token)
}

// RegisterUser mocks base method.
func (m *MockAuthService) RegisterUser(login, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", login, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockAuthServiceMockRecorder) RegisterUser(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockAuthService)(nil).RegisterUser), login, password)
}

// VerifyUser mocks base method.
func (m *MockAuthService) VerifyUser(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyUser indicates an expected call of VerifyUser.
func (mr *MockAuthServiceMockRecorder) VerifyUser(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockAuthService)(nil).VerifyUser), token)
}

// MockStorageAuth is a mock of StorageAuth interface.
type MockStorageAuth struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAuthMockRecorder
}

// MockStorageAuthMockRecorder is the mock recorder for MockStorageAuth.
type MockStorageAuthMockRecorder struct {
	mock *MockStorageAuth
}

// NewMockStorageAuth creates a new mock instance.
func NewMockStorageAuth(ctrl *gomock.Controller) *MockStorageAuth {
	mock := &MockStorageAuth{ctrl: ctrl}
	mock.recorder = &MockStorageAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageAuth) EXPECT() *MockStorageAuthMockRecorder {
	return m.recorder
}

// CheckTableUserLogin mocks base method.
func (m *MockStorageAuth) CheckTableUserLogin(login string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTableUserLogin", login)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckTableUserLogin indicates an expected call of CheckTableUserLogin.
func (mr *MockStorageAuthMockRecorder) CheckTableUserLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTableUserLogin", reflect.TypeOf((*MockStorageAuth)(nil).CheckTableUserLogin), login)
}

// CheckTableUserPassword mocks base method.
func (m *MockStorageAuth) CheckTableUserPassword(password string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTableUserPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckTableUserPassword indicates an expected call of CheckTableUserPassword.
func (mr *MockStorageAuthMockRecorder) CheckTableUserPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTableUserPassword", reflect.TypeOf((*MockStorageAuth)(nil).CheckTableUserPassword), password)
}

// SaveTableUser mocks base method.
func (m *MockStorageAuth) SaveTableUser(login, passwordHash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTableUser", login, passwordHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTableUser indicates an expected call of SaveTableUser.
func (mr *MockStorageAuthMockRecorder) SaveTableUser(login, passwordHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTableUser", reflect.TypeOf((*MockStorageAuth)(nil).SaveTableUser), login, passwordHash)
}

// SaveTableUserAndUpdateToken mocks base method.
func (m *MockStorageAuth) SaveTableUserAndUpdateToken(login, accessToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTableUserAndUpdateToken", login, accessToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTableUserAndUpdateToken indicates an expected call of SaveTableUserAndUpdateToken.
func (mr *MockStorageAuthMockRecorder) SaveTableUserAndUpdateToken(login, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTableUserAndUpdateToken", reflect.TypeOf((*MockStorageAuth)(nil).SaveTableUserAndUpdateToken), login, accessToken)
}

// SearchLoginByToken mocks base method.
func (m *MockStorageAuth) SearchLoginByToken(accessToken string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchLoginByToken", accessToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLoginByToken indicates an expected call of SearchLoginByToken.
func (mr *MockStorageAuthMockRecorder) SearchLoginByToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLoginByToken", reflect.TypeOf((*MockStorageAuth)(nil).SearchLoginByToken), accessToken)
}
