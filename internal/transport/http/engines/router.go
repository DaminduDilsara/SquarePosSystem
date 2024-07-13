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

		// add endpoints
	}

	return engine
}
