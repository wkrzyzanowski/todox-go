package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/server"
)

type HelloController struct {
	Endpoints []server.ApiEndpoint
}

func NewHelloController() *HelloController {
	return &HelloController{
		Endpoints: enpoints,
	}
}

func (controller *HelloController) GetEndpoints() []server.ApiEndpoint {
	return controller.Endpoints
}

var enpoints = []server.ApiEndpoint{
	{
		HttpMethod:   http.MethodGet,
		RelativePath: "/api/hello",
		HandlerFunc: []gin.HandlerFunc{
			func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Hello world!",
				})
			},
		},
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: "/api/hello",
		HandlerFunc: []gin.HandlerFunc{
			func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Hello world!",
				})
			},
		},
	},
}
