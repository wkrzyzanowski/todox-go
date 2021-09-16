package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/model"
)

type HelloController struct {
	Endpoints []model.ApiEndpoint
}

func NewHelloController() *HelloController {
	return &HelloController{
		Endpoints: enpoints,
	}
}

func (controller *HelloController) GetEndpoints() []model.ApiEndpoint {
	return controller.Endpoints
}

var enpoints = []model.ApiEndpoint{
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
