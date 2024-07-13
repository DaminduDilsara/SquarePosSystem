package square_client

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
)

type OrderClient interface {
	CreateOrder(request request_schemas.CreateOrderSquareRequest, authHeader string) (*response_schemas.CreateOrderSquareResponse, error)
	SearchOrders(request request_schemas.SearchOrdersSquareRequest, authHeader string) (*response_schemas.SearchOrdersSquareResponse, error)
	FindOrders(request request_schemas.FindOrdersSquareRequest, authHeader string) (*response_schemas.FindOrdersSquareResponse, error)
}
