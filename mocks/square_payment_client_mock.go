// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/clients/square_client/square_payment_client_impl.go

// Package mock_square_client is a generated GoMock package.
package mocks

import (
	request_schemas "SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	response_schemas "SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPaymentClient is a mock of PaymentClient interface.
type MockPaymentClient struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentClientMockRecorder
}

// MockPaymentClientMockRecorder is the mock recorder for MockPaymentClient.
type MockPaymentClientMockRecorder struct {
	mock *MockPaymentClient
}

// NewMockPaymentClient creates a new mock instance.
func NewMockPaymentClient(ctrl *gomock.Controller) *MockPaymentClient {
	mock := &MockPaymentClient{ctrl: ctrl}
	mock.recorder = &MockPaymentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentClient) EXPECT() *MockPaymentClientMockRecorder {
	return m.recorder
}

// CreatePayment mocks base method.
func (m *MockPaymentClient) CreatePayment(request request_schemas.CreatePaymentSquareRequest, authHeader string) (*response_schemas.CreatePaymentSquareResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", request, authHeader)
	ret0, _ := ret[0].(*response_schemas.CreatePaymentSquareResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockPaymentClientMockRecorder) CreatePayment(request, authHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockPaymentClient)(nil).CreatePayment), request, authHeader)
}
