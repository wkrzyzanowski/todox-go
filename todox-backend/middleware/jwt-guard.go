package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/controller/authenticationcontroller"
	"github.com/wkrzyzanowski/todox-go/server"
	"github.com/wkrzyzanowski/todox-go/service/authentication"
)

func NewJwtGuardMiddleware() server.ApiMiddleware {
	return server.ApiMiddleware{
		Name:     "Authorization Middleware",
		Function: filterRequest(),
	}
}

var notSecuredEndpoints = []string{
	fmt.Sprintf("%v/login", authenticationcontroller.AUTHENTICATION_BASE_URL),
	fmt.Sprintf("%v/refresh", authenticationcontroller.AUTHENTICATION_BASE_URL),
}

func filterRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		log.Println("...:: Authorization Middleware ::...")

		requestUrl := ctx.Request.URL.String()

		log.Println(requestUrl)

		for _, b := range notSecuredEndpoints {
			if requestUrl == b {
				log.Printf("Not secured endpoint is reached: %v", requestUrl)
				ctx.Next()
				return
			}
		}

		token := getToken(ctx)

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, server.MessageResponse{
				Message: "Authorization header is missing.",
			})

			ctx.Abort()
			return
		}

		isValidToken(token, ctx)

		ctx.Next()

	}
}

func isValidToken(token string, ctx *gin.Context) {
	_, err := authentication.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, server.MessageResponse{
			Message: "Not valid token!",
		})
		ctx.Abort()
		return
	}

	// if ctx.Request.URL.String() == notSecuredEndpoints[1] && parsedToken.Claims.(jwt.MapClaims)["type"] != authentication.ACCESS_TOKEN_TYPE {
	// 	ctx.JSON(http.StatusForbidden, server.MessageResponse{
	// 		Message: "Not appropiate token!",
	// 	})
	// 	ctx.Abort()
	// 	return
	// }
}

func getToken(ctx *gin.Context) string {
	headerList := ctx.Request.Header["Authorization"]
	if len(headerList) == 0 {
		return ""
	}
	return headerList[0]
}
