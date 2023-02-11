package db

import (
	"context"
	"time"

	"github.com/withmandala/go-log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_URL = "mongodb://admin:p4ssw0rd@localhost:27017"

func CreateClient(logger *log.Logger) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		logger.Fatal(err)
	}
	return client
}

func Connect(logger *log.Logger, client mongo.Client) context.Context {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Database connected on " + MONGO_URL)
	return ctx
}

func Disconnect(logger *log.Logger, client mongo.Client, ctx context.Context) {
	defer client.Disconnect(ctx)
}
