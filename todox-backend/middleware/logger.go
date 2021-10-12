package middleware

import (
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

		log.Printf(`
		...:: Logger Middleware ::...
		[Method]: %v
		[URL]: %v
		.............................`,
			ctx.Request.Method, ctx.Request.URL)
		ctx.Next()
	}
}
