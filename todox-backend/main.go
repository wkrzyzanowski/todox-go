package main

import (
	"github.com/wkrzyzanowski/todox-go/controller/hello"
	"github.com/wkrzyzanowski/todox-go/server"
)

func main() {
	startServer()
}

func startServer() {
	server.
		GetServerInstance().
		ServeVueWebApp().
		RegisterControllers(getControllers()).
		Run()
}

func getControllers() []server.ApiController {
	return []server.ApiController{
		hello.NewHelloController(),
	}
}
