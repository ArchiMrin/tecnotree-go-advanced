package main

import (
	"context"
	"fmt"

	"github.com/ArchiMrin/tecnotree-go-advanced/ekart/config"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	ctx         context.Context
)

func main() {
	ctx = context.Background()
	server := gin.Default()
	defer mongoClient.Disconnect(ctx)
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		fmt.Println("Error in connecting to DB")
	}

	routes.AppRoutes(server)
	server.GET("/", home)
	server.Run(":4001")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{ //H in JSON
		"data": "Welcome to gin --- GO",
	})
}
