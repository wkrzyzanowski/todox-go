package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(static.Serve("/", static.LocalFile("../webapp/dist/public", false)))

	server.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, "Hello World!")
	})

	server.Run()
}
