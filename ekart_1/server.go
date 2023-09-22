package main

import (
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.Approutes(server)
	server.GET("/", home)
	server.Run(":4000")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Welcome to gin",
	})
}
