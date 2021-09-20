package hello

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/server"
)

const HELLO_BASE_URL = server.BASE_API_URL + "/hello"

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
		RelativePath: HELLO_BASE_URL,
		HandlerFunc: []gin.HandlerFunc{
			GetHello(),
		},
	},
	{
		HttpMethod:   http.MethodPost,
		RelativePath: HELLO_BASE_URL,
		HandlerFunc: []gin.HandlerFunc{
			PostHello(),
		},
	},
}

func GetHello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world!",
		})
	}
}

func PostHello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jsonData, err := ctx.GetRawData()
		if err == nil {
			message := fmt.Sprintf("Hello world! Your message is: %v", string(jsonData))
			ctx.JSON(200, gin.H{
				"message": message,
			})
		} else {
			ctx.JSON(404, gin.H{
				"message": "Bad Request",
			})
		}
	}
}
