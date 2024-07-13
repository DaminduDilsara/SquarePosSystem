package square_client

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
)

type LocationClient interface {
	CreateLocation(request request_schemas.CreateLocationSquareRequest, authHeader string) (*response_schemas.CreateLocationSquareResponse, error)
}
