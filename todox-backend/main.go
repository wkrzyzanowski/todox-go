package main

import (
	"github.com/wkrzyzanowski/todox-go/controller/hello"
	"github.com/wkrzyzanowski/todox-go/model"
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

func getControllers() []model.ApiController {
	return []model.ApiController{
		hello.NewHelloController(),
	}
}
