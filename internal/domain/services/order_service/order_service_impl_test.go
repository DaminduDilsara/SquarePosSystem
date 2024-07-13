package order_service

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

func TestCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_square_client.NewMockOrderClient(ctrl)

	service := NewOrderService(mockClient)

	testCases := []struct {
		name        string
		request     request_schemas.CreateOrderIncomingRequest
		authHeader  string
		mockSetup   func()
		expected    *response_schemas.CreateOrderResponse
		expectedErr error
	}{
		{
			name: "successful order creation",
			request: request_schemas.CreateOrderIncomingRequest{
				LocationId:  "location1",
				CustomerId:  "customer1",
				ReferenceId: "ref1",
				TableId:     "table1",
				LineItems: []request_schemas.LineItems{
					{
						ItemId:   "item1",
						ItemName: "item_name_1",
						Quantity: "1",
					},
				},
			},
			authHeader: "Bearer token",
			mockSetup: func() {
				mockClient.EXPECT().
					CreateOrder(gomock.Any(), "Bearer token").
					Return(&response_schemas.CreateOrderSquareResponse{
						Order: response_schemas.SquareOrder{
							Id:        "order1",
							CreatedAt: time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
							State:     "COMPLETED",
							Source:    response_schemas.Source{Name: "table1"},
							LineItems: []response_schemas.LineItem{
								{
									Name:           "item_name_1",
									Quantity:       "1",
									BasePriceMoney: response_schemas.Money{Amount: 100},
									TotalMoney:     response_schemas.Money{Amount: 100},
								},
							},
							TotalMoney: response_schemas.Money{Amount: 100},
						},
					}, nil)
			},
			expected: &response_schemas.CreateOrderResponse{
				OrderResponse: response_schemas.OrderResponse{
					Id:       "order1",
					OpenedAt: time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
					IsClosed: true,
					Table:    "table1",
					Items: []response_schemas.Item{
						{
							Name:      "item_name_1",
							UnitPrice: 100,
							Quantity:  1,
							Amount:    100,
						},
					},
					Totals: response_schemas.Totals{
						Total: 100,
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "error during order creation",
			request: request_schemas.CreateOrderIncomingRequest{
				LocationId:  "location1",
				CustomerId:  "customer1",
				ReferenceId: "ref1",
				TableId:     "table1",
				LineItems: []request_schemas.LineItems{
					{
						ItemId:   "item1",
						ItemName: "item_name_1",
						Quantity: "1",
					},
				},
			},
			authHeader: "Bearer token",
			mockSetup: func() {
				mockClient.EXPECT().
					CreateOrder(gomock.Any(), "Bearer token").
					Return(nil, errors.New("creation error"))
			},
			expected:    nil,
			expectedErr: errors.New("creation error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			response, err := service.CreateOrder(tc.request, tc.authHeader)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected.Id, response.Id)
				assert.Equal(t, tc.expected.IsClosed, response.IsClosed)
				assert.Equal(t, tc.expected.Table, response.Table)
				assert.Equal(t, tc.expected.Totals.Total, response.Totals.Total)
				assert.Equal(t, len(tc.expected.Items), len(response.Items))
				for i := range tc.expected.Items {
					assert.Equal(t, tc.expected.Items[i].Name, response.Items[i].Name)
					assert.Equal(t, tc.expected.Items[i].UnitPrice, response.Items[i].UnitPrice)
					assert.Equal(t, tc.expected.Items[i].Quantity, response.Items[i].Quantity)
					assert.Equal(t, tc.expected.Items[i].Amount, response.Items[i].Amount)
				}
			}
		})
	}
}

func TestSearchOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_square_client.NewMockOrderClient(ctrl)

	service := NewOrderService(mockClient)

	testCases := []struct {
		name        string
		request     request_schemas.SearchOrdersIncomingRequest
		authHeader  string
		mockSetup   func()
		expected    *response_schemas.SearchOrdersResponse
		expectedErr error
	}{
		{
			name: "successful order search",
			request: request_schemas.SearchOrdersIncomingRequest{
				LocationId: "location1",
				TableNo:    "table1",
			},
			authHeader: "Bearer token",
			mockSetup: func() {
				mockClient.EXPECT().
					SearchOrders(gomock.Any(), "Bearer token").
					Return(&response_schemas.SearchOrdersSquareResponse{
						Orders: []response_schemas.SquareOrder{
							{
								Id:        "order1",
								CreatedAt: time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
								State:     "COMPLETED",
								Source:    response_schemas.Source{Name: "table1"},
								LineItems: []response_schemas.LineItem{
									{
										Name:           "item_name_1",
										Quantity:       "1",
										BasePriceMoney: response_schemas.Money{Amount: 100},
										TotalMoney:     response_schemas.Money{Amount: 100},
									},
								},
								TotalMoney: response_schemas.Money{Amount: 100},
							},
						},
					}, nil)
			},
			expected: &response_schemas.SearchOrdersResponse{
				Orders: []response_schemas.OrderResponse{
					{
						Id:       "order1",
						OpenedAt: time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
						IsClosed: true,
						Table:    "table1",
						Items: []response_schemas.Item{
							{
								Name:      "item_name_1",
								UnitPrice: 100,
								Quantity:  1,
								Amount:    100,
							},
						},
						Totals: response_schemas.Totals{
							Total: 100,
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "error during order search",
			request: request_schemas.SearchOrdersIncomingRequest{
				LocationId: "location1",
				TableNo:    "table1",
			},
			authHeader: "Bearer token",
			mockSetup: func() {
				mockClient.EXPECT().
					SearchOrders(gomock.Any(), "Bearer token").
					Return(nil, errors.New("search error"))
			},
			expected:    nil,
			expectedErr: errors.New("search error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			response, err := service.SearchOrders(tc.request, tc.authHeader)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tc.expected.Orders), len(response.Orders))
				for i := range tc.expected.Orders {
					assert.Equal(t, tc.expected.Orders[i].Id, response.Orders[i].Id)
					assert.Equal(t, tc.expected.Orders[i].IsClosed, response.Orders[i].IsClosed)
					assert.Equal(t, tc.expected.Orders[i].Table, response.Orders[i].Table)
					assert.Equal(t, tc.expected.Orders[i].Totals.Total, response.Orders[i].Totals.Total)
					assert.Equal(t, len(tc.expected.Orders[i].Items), len(response.Orders[i].Items))
					for j := range tc.expected.Orders[i].Items {
						assert.Equal(t, tc.expected.Orders[i].Items[j].Name, response.Orders[i].Items[j].Name)
						assert.Equal(t, tc.expected.Orders[i].Items[j].UnitPrice, response.Orders[i].Items[j].UnitPrice)
						assert.Equal(t, tc.expected.Orders[i].Items[j].Quantity, response.Orders[i].Items[j].Quantity)
						assert.Equal(t, tc.expected.Orders[i].Items[j].Amount, response.Orders[i].Items[j].Amount)
					}
				}
			}
		})
	}
}

func TestFindOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_square_client.NewMockOrderClient(ctrl)

	service := NewOrderService(mockClient)

	testCases := []struct {
		name        string
		request     request_schemas.FindOrdersIncomingRequest
		authHeader  string
		mockSetup   func()
		expected    *response_schemas.FindOrdersResponse
		expectedErr error
	}{
		{
			name: "successful order retrieval",
			request: request_schemas.FindOrdersIncomingRequest{
				OrderBatchRetrieveRequest: request_schemas.OrderBatchRetrieveRequest{
					OrderIds:   []string{"order1"},
					LocationId: "location1",
				},
			},
			authHeader: "Bearer token",
			mockSetup: func() {
				mockClient.EXPECT().
					FindOrders(gomock.Any(), "Bearer token").
					Return(&response_schemas.FindOrdersSquareResponse{
						Orders: []response_schemas.SquareOrder{
							{
								Id:        "order1",
								CreatedAt: time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
								State:     "COMPLETED",
								Source:    response_schemas.Source{Name: "table1"},
								LineItems: []response_schemas.LineItem{
									{
										Name:           "item_name_1",
										Quantity:       "1",
										BasePriceMoney: response_schemas.Money{Amount: 100},
										TotalMoney:     response_schemas.Money{Amount: 100},
									},
								},
								TotalMoney: response_schemas.Money{Amount: 100},
							},
						},
					}, nil)
			},
			expected: &response_schemas.FindOrdersResponse{
				Orders: []response_schemas.OrderResponse{
					{
						Id:       "order1",
						OpenedAt: time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
						IsClosed: true,
						Table:    "table1",
						Items: []response_schemas.Item{
							{
								Name:      "item_name_1",
								UnitPrice: 100,
								Quantity:  1,
								Amount:    100,
							},
						},
						Totals: response_schemas.Totals{
							Total: 100,
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "error during order retrieval",
			request: request_schemas.FindOrdersIncomingRequest{
				OrderBatchRetrieveRequest: request_schemas.OrderBatchRetrieveRequest{
					OrderIds:   []string{"order1"},
					LocationId: "location1",
				},
			},
			authHeader: "Bearer token",
			mockSetup: func() {
				mockClient.EXPECT().
					FindOrders(gomock.Any(), "Bearer token").
					Return(nil, errors.New("retrieval error"))
			},
			expected:    nil,
			expectedErr: errors.New("retrieval error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			response, err := service.FindOrders(tc.request, tc.authHeader)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tc.expected.Orders), len(response.Orders))
				for i := range tc.expected.Orders {
					assert.Equal(t, tc.expected.Orders[i].Id, response.Orders[i].Id)
					assert.Equal(t, tc.expected.Orders[i].IsClosed, response.Orders[i].IsClosed)
					assert.Equal(t, tc.expected.Orders[i].Table, response.Orders[i].Table)
					assert.Equal(t, tc.expected.Orders[i].Totals.Total, response.Orders[i].Totals.Total)
					assert.Equal(t, len(tc.expected.Orders[i].Items), len(response.Orders[i].Items))
					for j := range tc.expected.Orders[i].Items {
						assert.Equal(t, tc.expected.Orders[i].Items[j].Name, response.Orders[i].Items[j].Name)
						assert.Equal(t, tc.expected.Orders[i].Items[j].UnitPrice, response.Orders[i].Items[j].UnitPrice)
						assert.Equal(t, tc.expected.Orders[i].Items[j].Quantity, response.Orders[i].Items[j].Quantity)
						assert.Equal(t, tc.expected.Orders[i].Items[j].Amount, response.Orders[i].Items[j].Amount)
					}
				}
			}
		})
	}
}
