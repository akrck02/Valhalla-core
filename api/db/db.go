package db

import (
	"context"
	"time"

	"github.com/withmandala/go-log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase(logger *log.Logger) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:p4ssw0rd@localhost:27017"))
	if err != nil {
		logger.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		logger.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(databases)
	logger.Info("Hola")
}
