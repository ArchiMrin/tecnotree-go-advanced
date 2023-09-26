package controllers

import (
	"fmt"
	"net/http"

	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/entities"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/interfaces"
	"github.com/gin-gonic/gin"
	//"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/services"
)

type ProductController struct {
	ProductService interfaces.IProduct
	//ProductController services.ProductService //this is our hard coupling dependency which we won't do
}

func InitProductController(productSvc interfaces.IProduct) *ProductController {
	return &ProductController{ProductService: productSvc}
}

func (p ProductController) InsertProduct(c *gin.Context) {
	fmt.Println("Invoked Controller")
	//extract the product from the Request Object
	var product entities.Product
	err := c.BindJSON(&product) // we call the rest api to extract the product details
	if err != nil {
		return
	}
	//call the service.Insert
	result, err := p.ProductService.Insert(&product)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
	//call p.Product service.InsertOne()
}
