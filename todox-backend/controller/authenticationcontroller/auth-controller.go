package authenticationcontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/server"
	"github.com/wkrzyzanowski/todox-go/service/authentication"
)

const AUTHENTICATION_BASE_URL = server.BASE_API_URL + "/authentication"

type AuthenticationRequest struct {
	Username string `json:"username"`
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
	{
		HttpMethod:   http.MethodPost,
		RelativePath: fmt.Sprintf("%v/refresh", AUTHENTICATION_BASE_URL),
		HandlerFunc: []gin.HandlerFunc{
			Refresh(),
		},
	},
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestBody, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, server.MessageResponse{
				Message: "Cannot parse request body.",
			})
		}
		var request AuthenticationRequest
		json.Unmarshal(requestBody, &request)

		token, err := authentication.GetAuthentication(request.Username, request.Password)
		getAuthenticationResponse(ctx, token, err)

	}
}

func Refresh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken := ctx.Request.Header["Authorization"]
		token, err := authentication.RefreshAuthentication(refreshToken[0])
		getAuthenticationResponse(ctx, token, err)
	}
}

func getAuthenticationResponse(ctx *gin.Context, token authentication.AuthorizationToken, err error) {
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, server.MessageResponse{
			Message: "====> Authentication failure!",
		})
	} else {
		ctx.Header("access-token", fmt.Sprintf("Bearer %v", token.AccessToken))
		ctx.Header("refresh-token", fmt.Sprintf("Bearer %v", token.RefreshToken))
		ctx.JSON(http.StatusOK, server.MessageResponse{
			Message: "====> Authentication successful!",
		})

	}
}
