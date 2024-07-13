package v1

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/services/location_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerV1 struct {
	locationService location_service.LocationService
}

func NewControllerV1(locationSrc location_service.LocationService) *ControllerV1 {
	return &ControllerV1{
		locationService: locationSrc,
	}
}

func (con ControllerV1) CreateLocationController(c *gin.Context) {
	var incomingReq request_schemas.CreateLocationIncomingRequest
	if err := c.ShouldBindJSON(&incomingReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "auth token is required"})
		return
	}

	outgoingResp, err := con.locationService.CreateLocation(incomingReq, authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, outgoingResp)
}
