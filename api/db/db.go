package db

import (
	"context"
	"time"

	"github.com/akrck02/valhalla-core/configuration"
	"github.com/akrck02/valhalla-core/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_URL = "mongodb://"
const MONGO_USER = "admin"
const MONGO_PASSWORD = "p4ssw0rd"
const MONGO_PORT = "27017"

func CreateClient() *mongo.Client {

	var host = configuration.Params.Mongo

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URL + MONGO_USER + ":" + MONGO_PASSWORD + "@" + host + ":" + MONGO_PORT))
	if err != nil {
		log.Fatal(err.Error())
	}
	return client
}

func Connect(client mongo.Client) context.Context {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("Database connected on " + MONGO_URL)
	return ctx
}

func Disconnect(client mongo.Client, ctx context.Context) {
	defer client.Disconnect(ctx)
}
