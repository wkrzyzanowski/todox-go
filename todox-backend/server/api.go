package server

import (
	"github.com/gin-gonic/gin"
)

const BASE_API_URL = "/api"

type ApiEndpoint struct {
	RelativePath string
	HttpMethod   string
	HandlerFunc  []gin.HandlerFunc
}

type ApiController interface {
	GetEndpoints() []ApiEndpoint
}

type ApiMiddleware struct {
	Name     string
	Function gin.HandlerFunc
}

type MessageResponse struct {
	Message string `json:"message"`
}
