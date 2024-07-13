package location_service

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	mock_square_client "SquarePosSystem/mocks"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_square_client.NewMockLocationClient(ctrl)

	locationService := NewLocationService(mockClient)

	testCases := []struct {
		name        string
		request     request_schemas.CreateLocationIncomingRequest
		authHeader  string
		mockSetup   func()
		expected    *response_schemas.CreateLocationResponse
		expectedErr error
	}{
		{
			name: "successful creation",
			request: request_schemas.CreateLocationIncomingRequest{
				BusinessEmail: "test@example.com",
				Description:   "A test location",
				BusinessName:  "Test Business",
			},
			authHeader: "Bearer testtoken",
			mockSetup: func() {
				mockClient.EXPECT().
					CreateLocation(gomock.Any(), "Bearer testtoken").
					Return(&response_schemas.CreateLocationSquareResponse{
						Location: struct {
							Id            string    `json:"id"`
							Name          string    `json:"name"`
							Timezone      string    `json:"timezone"`
							Capabilities  []string  `json:"capabilities"`
							Status        string    `json:"status"`
							CreatedAt     time.Time `json:"created_at"`
							MerchantId    string    `json:"merchant_id"`
							Country       string    `json:"country"`
							LanguageCode  string    `json:"language_code"`
							Currency      string    `json:"currency"`
							BusinessName  string    `json:"business_name"`
							Type          string    `json:"type"`
							BusinessEmail string    `json:"business_email"`
							Description   string    `json:"description"`
							Mcc           string    `json:"mcc"`
						}{
							Id:            "123",
							Name:          "Test Business",
							CreatedAt:     time.Date(2024, time.July, 13, 20, 41, 9, 239897952, time.Local),
							Description:   "A test location",
							BusinessEmail: "test@example.com",
						},
					}, nil)
			},
			expected: &response_schemas.CreateLocationResponse{
				Id:            "123",
				Name:          "Test Business",
				CreatedAt:     time.Date(2024, time.July, 13, 20, 41, 9, 239897952, time.Local),
				Description:   "A test location",
				BusinessEmail: "test@example.com",
			},
			expectedErr: nil,
		},
		{
			name: "error during creation",
			request: request_schemas.CreateLocationIncomingRequest{
				BusinessEmail: "test@example.com",
				Description:   "A test location",
				BusinessName:  "Test Business",
			},
			authHeader: "Bearer testtoken",
			mockSetup: func() {
				mockClient.EXPECT().
					CreateLocation(gomock.Any(), "Bearer testtoken").
					Return(nil, errors.New("creation error"))
			},
			expected:    nil,
			expectedErr: errors.New("creation error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			response, err := locationService.CreateLocation(tc.request, tc.authHeader)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, response)
			}
		})
	}
}
