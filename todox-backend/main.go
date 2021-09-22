package main

import (
	"github.com/wkrzyzanowski/todox-go/controller/authentication"
	"github.com/wkrzyzanowski/todox-go/controller/hello"
	"github.com/wkrzyzanowski/todox-go/middleware"
	"github.com/wkrzyzanowski/todox-go/server"
)

func main() {
	startServer()
}

func startServer() {
	server.
		GetServerInstance().
		RegisterGlobalHandlers(getGlobalMiddleware()).
		ServeVueWebApp().
		RegisterControllers(getControllers()).
		Run()
}

func getControllers() []server.ApiController {
	return []server.ApiController{
		hello.NewHelloController(),
		authentication.NewAuthenticationController(),
	}
}

func getGlobalMiddleware() []server.ApiMiddleware {
	return []server.ApiMiddleware{
		middleware.NewLoggerMiddleware(),
	}
}
