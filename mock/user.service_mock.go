// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/mac/GolandProjects/mnc-stage2/src/service/user.service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	data "mnc-stage2/src/data"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	decimal "github.com/shopspring/decimal"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockUserService) Login(ctx context.Context, phoneNumber, pin string) (*data.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, phoneNumber, pin)
	ret0, _ := ret[0].(*data.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(ctx, phoneNumber, pin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), ctx, phoneNumber, pin)
}

// Payment mocks base method.
func (m *MockUserService) Payment(ctx context.Context, userID string, req data.PaymentReq) (*data.PaymentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Payment", ctx, userID, req)
	ret0, _ := ret[0].(*data.PaymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Payment indicates an expected call of Payment.
func (mr *MockUserServiceMockRecorder) Payment(ctx, userID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Payment", reflect.TypeOf((*MockUserService)(nil).Payment), ctx, userID, req)
}

// RegisterUser mocks base method.
func (m *MockUserService) RegisterUser(ctx context.Context, firstName, lastName, phoneNumber, address, pin string) (*data.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, firstName, lastName, phoneNumber, address, pin)
	ret0, _ := ret[0].(*data.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserServiceMockRecorder) RegisterUser(ctx, firstName, lastName, phoneNumber, address, pin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserService)(nil).RegisterUser), ctx, firstName, lastName, phoneNumber, address, pin)
}

// TopUp mocks base method.
func (m *MockUserService) TopUp(ctx context.Context, userID string, amount decimal.Decimal) (*data.TopUpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopUp", ctx, userID, amount)
	ret0, _ := ret[0].(*data.TopUpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopUp indicates an expected call of TopUp.
func (mr *MockUserServiceMockRecorder) TopUp(ctx, userID, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopUp", reflect.TypeOf((*MockUserService)(nil).TopUp), ctx, userID, amount)
}

// TransactionReports mocks base method.
func (m *MockUserService) TransactionReports(ctx context.Context, userID string) (*data.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReports", ctx, userID)
	ret0, _ := ret[0].(*data.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReports indicates an expected call of TransactionReports.
func (mr *MockUserServiceMockRecorder) TransactionReports(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReports", reflect.TypeOf((*MockUserService)(nil).TransactionReports), ctx, userID)
}

// Transfer mocks base method.
func (m *MockUserService) Transfer(ctx context.Context, userID string, req data.TransferReq) (*data.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, userID, req)
	ret0, _ := ret[0].(*data.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockUserServiceMockRecorder) Transfer(ctx, userID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockUserService)(nil).Transfer), ctx, userID, req)
}
