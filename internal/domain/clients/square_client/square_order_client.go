package square_client

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
)

type OrderClient interface {
	CreateOrder(request request_schemas.CreateOrderSquareRequest, authHeader string) (*response_schemas.CreateOrderSquareResponse, error)
}
