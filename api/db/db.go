package db

import (
	"context"
	"time"

	"github.com/akrck02/valhalla-core/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_URL = "mongodb://admin:p4ssw0rd@172.20.0.10:27017"

func CreateClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URL))
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
