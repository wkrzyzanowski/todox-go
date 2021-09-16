package server

import (
	"github.com/gin-gonic/gin"
)

type ApiEndpoint struct {
	RelativePath string
	HttpMethod   string
	HandlerFunc  []gin.HandlerFunc
}

type ApiController interface {
	GetEndpoints() []ApiEndpoint
}
