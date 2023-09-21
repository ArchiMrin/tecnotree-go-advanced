package config

import (
	"context"
	"log"
	"time"

	"github.com/ArchiMrin/tecnotree-go-advanced/ekart/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDataBase() (*mongo.Client, error) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	mongoConnection := options.Client().ApplyURI(constants.DBConnection)

	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return mongoClient, nil
}
