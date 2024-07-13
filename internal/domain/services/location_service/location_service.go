package location_service

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
)

type LocationService interface {
	CreateLocation(request request_schemas.CreateLocationIncomingRequest, authHeader string) (*response_schemas.CreateLocationResponse, error)
}
