package order_service

import (
	"SquarePosSystem/internal/domain/clients/square_client"
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
)

// orderService implements the OrderService interface
type orderService struct {
	client square_client.OrderClient
}

// NewOrderService creates a new OrderService
func NewOrderService(client square_client.OrderClient) OrderService {
	return &orderService{client: client}
}

func (o orderService) CreateOrder(request request_schemas.CreateOrderIncomingRequest, authHeader string) (*response_schemas.CreateOrderResponse, error) {
	// Prepare the internal API request payload
	internalReq := request_schemas.CreateOrderSquareRequest{
		Order: request_schemas.Order{
			LocationId:  request.LocationId,
			CustomerId:  request.CustomerId,
			ReferenceId: request.ReferenceId,
			TicketName:  request.TableId,
			Source: struct {
				Name string `json:"name"`
			}{Name: request.TableId},
			LineItems: make([]struct {
				CatalogObjectId string `json:"catalog_object_id"`
				request_schemas.LineItems
			}, len(request.LineItems)),
		},
	}

	for i, item := range request.LineItems {
		internalReq.Order.LineItems[i] = struct {
			CatalogObjectId string `json:"catalog_object_id"`
			request_schemas.LineItems
		}{
			CatalogObjectId: item.ItemId,
			LineItems:       item,
		}
	}

	// Call the client function
	internalResp, err := o.client.CreateOrder(internalReq, authHeader)
	if err != nil {
		return nil, err
	}

	response := o.ConvertToCreateOrderResponse(internalResp)

	return response, nil
}

func (o orderService) ConvertToCreateOrderResponse(squareResp *response_schemas.CreateOrderSquareResponse) *response_schemas.CreateOrderResponse {
	order := squareResp.Order
	createOrderResp := response_schemas.CreateOrderResponse{
		OrderResponse: response_schemas.OrderResponse{
			Id:       order.Id,
			OpenedAt: order.CreatedAt,
			IsClosed: order.State == "COMPLETED",
			Table:    order.ReferenceId,
			Items: []struct {
				Name      string `json:"name"`
				Comment   string `json:"comment"`
				UnitPrice int    `json:"unit_price"`
				Quantity  int    `json:"quantity"`
				Discounts []struct {
					Name         string `json:"name"`
					IsPercentage bool   `json:"is_percentage"`
					Value        int    `json:"value"`
					Amount       int    `json:"amount"`
				} `json:"discounts"`
				Modifiers []struct {
					Name      string `json:"name"`
					UnitPrice int    `json:"unit_price"`
					Quantity  int    `json:"quantity"`
					Amount    int    `json:"amount"`
				} `json:"modifiers"`
				Amount int `json:"amount"`
			}{},
			Totals: struct {
				Discounts     int `json:"discounts"`
				Due           int `json:"due"`
				Tax           int `json:"tax"`
				ServiceCharge int `json:"service_charge"`
				Paid          int `json:"paid"`
				Tips          int `json:"tips"`
				Total         int `json:"total"`
			}{
				Discounts:     order.TotalDiscountMoney.Amount,
				Due:           order.NetAmountDueMoney.Amount,
				Tax:           order.TotalTaxMoney.Amount,
				ServiceCharge: order.TotalServiceChargeMoney.Amount,
				Paid:          order.TotalMoney.Amount,
				Tips:          order.TotalTipMoney.Amount,
				Total:         order.TotalMoney.Amount,
			},
		},
	}

	return &createOrderResp
}
