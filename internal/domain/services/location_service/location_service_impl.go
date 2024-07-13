package location_service

import (
	"SquarePosSystem/internal/domain/clients/square_client"
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	"log"
)

const locationServiceLogPrefix = "location_service_impl"

type locationService struct {
	client square_client.LocationClient
}

func NewLocationService(client square_client.LocationClient) LocationService {
	return &locationService{client: client}
}

func (l locationService) CreateLocation(request request_schemas.CreateLocationIncomingRequest, authHeader string) (*response_schemas.CreateLocationResponse, error) {
	internalReq := request_schemas.CreateLocationSquareRequest{
		Location: request_schemas.Location{
			BusinessEmail: request.BusinessEmail,
			Description:   request.Description,
			Name:          request.BusinessName,
		},
	}

	internalResp, err := l.client.CreateLocation(internalReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", locationServiceLogPrefix, err)
		return nil, err
	}

	outGoingResp := response_schemas.CreateLocationResponse{
		Id:            internalResp.Location.Id,
		Name:          internalResp.Location.Name,
		CreatedAt:     internalResp.Location.CreatedAt,
		Description:   internalResp.Location.Description,
		BusinessEmail: internalResp.Location.BusinessEmail,
	}

	return &outGoingResp, nil

}
