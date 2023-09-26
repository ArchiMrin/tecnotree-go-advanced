package services

import (
	"context"

	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/entities"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection //a data called Product carrying the context of collection instance
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct { //instantiating the product service
	return &ProductService{Product: collection}
}

func (p *ProductService) Insert(product *entities.Product) (string, error) {
	product.ID = primitive.NewObjectID()
	_, err := p.Product.InsertOne(context.Background(), product) //p.Product gives the Collection & then we call the operation InsertOne

	if err != nil {
		return "", err
	} else {
		return "Record Inserted", nil
	}

}
