package engines

import (
	v1 "SquarePosSystem/internal/transport/http/controllers/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Engine struct {
	controllerV1 *v1.ControllerV1
}

func NewEngine(
	controllerV1 *v1.ControllerV1,
) *Engine {
	return &Engine{
		controllerV1: controllerV1,
	}
}

func (e *Engine) GetEngine() *gin.Engine {
	engine := gin.New()

	v1Group := engine.Group("/v1.0")
	{
		engine.GET("/ping", func(context *gin.Context) {
			context.String(http.StatusOK, "pong")
		})

		v1Group.POST("/location/create", e.controllerV1.CreateLocationController)

		v1Group.POST("/order/create", e.controllerV1.CreateOrderController)

		v1Group.POST("/orders/search", e.controllerV1.SearchOrdersController)

		v1Group.POST("/orders/find", e.controllerV1.FindOrdersController)

		v1Group.POST("/payment/create", e.controllerV1.CreatePaymentController)

		// add endpoints
	}

	return engine
}
