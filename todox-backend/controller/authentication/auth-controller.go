package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/server"
	"github.com/wkrzyzanowski/todox-go/service/authentication"
)

const AUTHENTICATION_BASE_URL = server.BASE_API_URL + "/authentication"

type AuthenticationRequest struct {
	Password string `json:"password"`
}

type AuthenticationController struct {
	Endpoints []server.ApiEndpoint
}

func NewAuthenticationController() *AuthenticationController {
	return &AuthenticationController{
		Endpoints: enpoints,
	}
}

func (controller *AuthenticationController) GetEndpoints() []server.ApiEndpoint {
	return controller.Endpoints
}

var enpoints = []server.ApiEndpoint{
	{
		HttpMethod:   http.MethodPost,
		RelativePath: fmt.Sprintf("%v/login", AUTHENTICATION_BASE_URL),
		HandlerFunc: []gin.HandlerFunc{
			Login(),
		},
	},
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(400, server.MessageResponse{
				Message: "Cannot parse request body.",
			})
		}
		var secret AuthenticationRequest
		json.Unmarshal(requestBody, &secret)

		if authentication.GetAuthentication(secret.Password) {
			ctx.JSON(200, server.MessageResponse{
				Message: "====> Authentication successful!",
			})
		} else {
			ctx.JSON(404, server.MessageResponse{
				Message: "====> Authentication failure!",
			})
		}
	}
}
