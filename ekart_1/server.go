package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/config"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/controllers"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/routes"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	ctx         context.Context
	server      *gin.Engine
)

func main() {
	//ctx = context.Background()
	server = gin.Default()
	//DbOperations()
	//defer mongoClient.Disconnect(ctx)
	InitializeDatabase()
	InitializeProducts()
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer mongoClient.Disconnect(ctx1)
	// if err != nil {
	// 	fmt.Println("Error connecting database")
	// }
	routes.Approutes(server)
	//server.GET("/", home)
	server.Run(":4000")
}

// func home(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"data": "Welcome to gin",
// 	})
// }

//Initialize Database

func InitializeDatabase() {
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		log.Fatalf("Unable to connect to DB", err)
	} else {
		fmt.Println("Database Connected")
	}

}

func InitializeProducts() {
	//ctx = context.Background()
	//defer mongoClient.Disconnect(ctx)

	productCollection := config.GetCollection(mongoClient, "ekart", "products")
	productSvc := services.InitProductService(productCollection)
	productCtrl := controllers.InitProductController(productSvc)
	routes.ProductRoutes(server, *productCtrl)
	// product := entities.Product{ID: primitive.NewObjectID(), Name: "Redmi", Price: 15000, Description: "MI Phone"}
	// result, err := productCollection.InsertOne(ctx, product)
	// if err != nil {
	// 	fmt.Println("Error in inserting")
	// } else {
	// 	fmt.Println("Successfully inserted", result)
	// }

}
