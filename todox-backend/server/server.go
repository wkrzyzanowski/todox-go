package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/wkrzyzanowski/todox-go/model"
)

const VUE_APP_SOURCES = "../webapp/dist/public"
const SERVER_APP_PORT = ":8080"

type Server struct {
	instance *gin.Engine
}

var serverInstance = Server{
	instance: gin.Default(),
}

func GetServerInstance() *Server {
	return &serverInstance
}

func (s *Server) ServeVueWebApp() *Server {
	serverInstance.instance.Use(static.Serve("/", static.LocalFile(VUE_APP_SOURCES, false)))
	return s
}

func (s *Server) RegisterControllers(apiController []model.ApiController) *Server {
	for _, controller := range apiController {
		for _, endpoint := range controller.GetEndpoints() {

			switch x := endpoint.HttpMethod; x {
			case http.MethodGet:
				serverInstance.instance.GET(endpoint.RelativePath, endpoint.HandlerFunc...)
			case http.MethodPost:
				serverInstance.instance.POST(endpoint.RelativePath, endpoint.HandlerFunc...)
			case http.MethodPut:
				serverInstance.instance.PUT(endpoint.RelativePath, endpoint.HandlerFunc...)
			case http.MethodPatch:
				serverInstance.instance.PATCH(endpoint.RelativePath, endpoint.HandlerFunc...)
			case http.MethodDelete:
				serverInstance.instance.DELETE(endpoint.RelativePath, endpoint.HandlerFunc...)
			default:
				msg := fmt.Sprintf("Http Method misconfigured or not supported: %v", controller)
				log.Fatalln(msg)
			}

		}
	}
	return s
}

func (s *Server) Run() {
	serverInstance.instance.Run(SERVER_APP_PORT)
}
