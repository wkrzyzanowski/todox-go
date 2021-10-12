package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

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

func IsValidJsonRequestBody(ctx *gin.Context) bool {
	bodyString := GetJsonBody(ctx)
	var js json.RawMessage
	return json.Unmarshal([]byte(bodyString), &js) == nil
}

func GetJsonBody(ctx *gin.Context) string {
	var bodyBytes []byte
	if ctx.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}
