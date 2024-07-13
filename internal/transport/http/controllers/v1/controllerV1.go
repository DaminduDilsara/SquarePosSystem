package v1

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/services/location_service"
	"SquarePosSystem/internal/domain/services/order_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerV1 struct {
	locationService location_service.LocationService
	orderService    order_service.OrderService
}

func NewControllerV1(
	locationSrc location_service.LocationService,
	orderSrc order_service.OrderService,
) *ControllerV1 {
	return &ControllerV1{
		locationService: locationSrc,
		orderService:    orderSrc,
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}

	outgoingResp, err := con.locationService.CreateLocation(incomingReq, authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, outgoingResp)
}

func (con ControllerV1) CreateOrderController(c *gin.Context) {
	var incomingReq request_schemas.CreateOrderIncomingRequest
	if err := c.ShouldBindJSON(&incomingReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}

	outgoingResp, err := con.orderService.CreateOrder(incomingReq, authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, outgoingResp)
}

func (con ControllerV1) SearchOrdersController(c *gin.Context) {
	var searchReq request_schemas.SearchOrdersIncomingRequest
	if err := c.ShouldBindJSON(&searchReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}

	searchResp, err := con.orderService.SearchOrders(searchReq, authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, searchResp)
}

func (con ControllerV1) FindOrdersController(c *gin.Context) {
	var searchReq request_schemas.FindOrdersIncomingRequest
	if err := c.ShouldBindJSON(&searchReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}

	searchResp, err := con.orderService.FindOrders(searchReq, authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, searchResp)
}
