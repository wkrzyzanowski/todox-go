package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"

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

		// Read the content
		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
		}
		// Restore the io.ReadCloser to its original state
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		bodyString := string(bodyBytes)

		log.Printf(`
		...:: Logger Middleware ::...
		[Method]: %v
		[URL]: %v
		[Body]: %v
		.............................`,
			ctx.Request.Method, ctx.Request.URL, strings.Join(strings.Fields(bodyString), ""))
		ctx.Next()
	}
}
