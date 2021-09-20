package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/server"
)

func NewLoggerMiddleware() server.ApiMiddleware {
	return server.ApiMiddleware{
		Name:     "Logging Middleware",
		Function: logRequest(),
	}
}

func logRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		body := ""

		jsonData, err := ctx.GetRawData()

		if err == nil {
			body = fmt.Sprintf("<%v>", string(jsonData))
		} else {
			body = "<Body parse error>"
		}

		log.Printf(`
		...:: Logger Middleware ::...
		[Method]: %v
		[URL]: %v
		[Body]: %v
		.............................`,
			ctx.Request.Method, ctx.Request.URL, body)
		ctx.Next()
	}
}
