package engines

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) GetEngine() *gin.Engine {
	engine := gin.New()

	v1Group := engine.Group("/v1.0")
	{
		engine.GET("/ping", func(context *gin.Context) {
			context.String(http.StatusOK, "pong")
		})

		v1Group.GET("/testPath", func(context *gin.Context) {
			context.String(http.StatusOK, "test done")
		})

		// add endpoints
	}

	return engine
}
