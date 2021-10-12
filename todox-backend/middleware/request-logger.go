package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/server"
	"github.com/wkrzyzanowski/todox-go/tools"
)

func NewLoggerMiddleware() server.ApiMiddleware {
	return server.ApiMiddleware{
		Name:     "Logging Middleware",
		Function: logRequest(),
	}
}

func logRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tools.LOGGER.Info(fmt.Sprintf(`
		...:: Logger Middleware ::...
		[Method]: %v
		[URL]: %v
		.............................`,
			ctx.Request.Method, ctx.Request.URL))
		ctx.Next()
	}
}
