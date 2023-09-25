package main

import (
	"context"
	"fmt"

	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/config"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/entities"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	ctx         context.Context
)

func main() {
	//ctx = context.Background()
	server := gin.Default()
	DbOperations()
	//defer mongoClient.Disconnect(ctx)

	routes.Approutes(server)
	server.GET("/", home)
	server.Run(":4000")
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Welcome to gin",
	})
}

func DbOperations() {
	ctx = context.Background()
	defer mongoClient.Disconnect(ctx)
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		fmt.Println("Error connection to DB")
	}
	productCollection := config.GetCollection(mongoClient, "ekart", "products")
	product := entities.Product{ID: primitive.NewObjectID(), Name: "Redmi", Price: 15000, Description: "MI Phone"}
	result, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		fmt.Println("Error in inserting")
	} else {
		fmt.Println("Successfully inserted", result)
	}

}
