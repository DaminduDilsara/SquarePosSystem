package payment_service

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	mock_square_client "SquarePosSystem/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreatePayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_square_client.NewMockPaymentClient(ctrl)

	service := NewPaymentService(mockClient)

	testCases := []struct {
		name        string
		request     request_schemas.CreatePaymentRequest
		authHeader  string
		mockSetup   func()
		expected    *response_schemas.CreatePaymentResponse
		expectedErr error
	}{
		{
			name: "successful payment creation",
			request: request_schemas.CreatePaymentRequest{
				BillAmount: 1000,
				TipAmount:  200,
				OrderID:    "order1",
			},
			authHeader: "Bearer token",
			mockSetup: func() {
				mockClient.EXPECT().
					CreatePayment(gomock.Any(), "Bearer token").
					Return(&response_schemas.CreatePaymentSquareResponse{
						Payment: struct {
							Id            string                 `json:"id"`
							CreatedAt     time.Time              `json:"created_at"`
							UpdatedAt     time.Time              `json:"updated_at"`
							Amount        response_schemas.Money `json:"amount_money"`
							Tip           response_schemas.Money `json:"tip_money"`
							Status        string                 `json:"status"`
							SourceType    string                 `json:"source_type"`
							LocationId    string                 `json:"location_id"`
							OrderId       string                 `json:"order_id"`
							Total         response_schemas.Money `json:"total_money"`
							ReceiptNumber string                 `json:"receipt_number"`
							ReceiptUrl    string                 `json:"receipt_url"`
						}{
							Id:        "payment1",
							CreatedAt: time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
							Amount: response_schemas.Money{
								Amount:   1000,
								Currency: "LKR",
							},
							Tip: response_schemas.Money{
								Amount:   200,
								Currency: "LKR",
							},
							Total: response_schemas.Money{
								Amount:   1200,
								Currency: "LKR",
							},
							ReceiptNumber: "12345",
							ReceiptUrl:    "http://receipt.url",
						},
					}, nil)
			},
			expected: &response_schemas.CreatePaymentResponse{
				PaymentId:   "payment1",
				PaymentTime: time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
				Amount: response_schemas.Money{
					Amount:   1000,
					Currency: "LKR",
				},
				Tip: response_schemas.Money{
					Amount:   200,
					Currency: "LKR",
				},
				Total: response_schemas.Money{
					Amount:   1200,
					Currency: "LKR",
				},
				ReceiptNumber: "12345",
				ReceiptUrl:    "http://receipt.url",
			},
			expectedErr: nil,
		},
		{
			name: "error during payment creation",
			request: request_schemas.CreatePaymentRequest{
				BillAmount: 1000,
				TipAmount:  200,
				OrderID:    "order1",
			},
			authHeader: "Bearer token",
			mockSetup: func() {
				mockClient.EXPECT().
					CreatePayment(gomock.Any(), "Bearer token").
					Return(nil, errors.New("creation error"))
			},
			expected:    nil,
			expectedErr: errors.New("creation error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			response, err := service.CreatePayment(tc.request, tc.authHeader)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected.PaymentId, response.PaymentId)
				assert.Equal(t, tc.expected.PaymentTime, response.PaymentTime)
				assert.Equal(t, tc.expected.Amount, response.Amount)
				assert.Equal(t, tc.expected.Tip, response.Tip)
				assert.Equal(t, tc.expected.Total, response.Total)
				assert.Equal(t, tc.expected.ReceiptNumber, response.ReceiptNumber)
				assert.Equal(t, tc.expected.ReceiptUrl, response.ReceiptUrl)
			}
		})
	}
}
