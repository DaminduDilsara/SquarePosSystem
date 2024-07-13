package order_service

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
)

type OrderService interface {
	CreateOrder(request request_schemas.CreateOrderIncomingRequest, authHeader string) (*response_schemas.CreateOrderResponse, error)
	SearchOrders(request request_schemas.SearchOrdersIncomingRequest, authHeader string) (*response_schemas.SearchOrdersResponse, error)
}
